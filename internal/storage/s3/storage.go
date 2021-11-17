package s3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/dotindustries/moar/internal"
	"github.com/dotindustries/moar/internal/storage"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
)

const manifestSuffix = "/manifest.json"

type Storage struct {
	minioClient *minio.Client
	logger      *logrus.Entry
	bucket      string
}

func (s *Storage) PutModule(ctx context.Context, module internal.Module) error {
	bts, err := json.Marshal(module)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(bts)
	// clear file data before serializing
	for _, version := range module.Versions {
		for _, file := range version.Files {
			file.Data = nil
		}
	}
	objectName := moduleManifestObjectName(module.Name)
	_, err = s.minioClient.PutObject(
		ctx,
		s.bucket,
		objectName,
		reader,
		int64(reader.Len()),
		minio.PutObjectOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) RemoveModule(ctx context.Context, name string) error {
	objectName := moduleManifestObjectNameFromString(name)
	err := s.minioClient.RemoveObject(
		ctx,
		s.bucket,
		objectName,
		minio.RemoveObjectOptions{
			ForceDelete: true,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) PutVersion(ctx context.Context, module string, version string, files []internal.File) error {
	var added []string
	var err error
	basePath := moduleVersionBasePath(module, version)
	defer func() {
		// rollback function in case of errors
		if err != nil {
			for _, objectName := range added {
				rollBackErr := s.minioClient.RemoveObject(ctx, s.bucket, objectName, minio.RemoveObjectOptions{ForceDelete: true})
				if rollBackErr != nil {
					s.logger.Errorf("Failed to roll back after failing to put stylesheet: %s", rollBackErr)
					// update err to rollback error
					err = rollBackErr
				}
			}
		}
	}()
	for _, file := range files {
		r := bytes.NewReader(file.Data)
		objectName := basePath + file.Name
		info, err := s.minioClient.PutObject(
			ctx,
			s.bucket,
			objectName,
			r,
			r.Size(),
			minio.PutObjectOptions{
				ContentType: file.MimeType,
			},
		)
		if err != nil {
			return err
		}
		s.logger.Infof("Successfully written %d bytes for %s", info.Size, info.Key)
		added = append(added, objectName)
	}
	return nil
}

func (s *Storage) RemoveVersion(ctx context.Context, module string, version string) error {
	basePath := moduleVersionBasePath(module, version)
	return s.minioClient.RemoveObject(
		ctx,
		s.bucket,
		basePath,
		minio.RemoveObjectOptions{
			ForceDelete: true,
		},
	)
}

func (s *Storage) ModuleResources(ctx context.Context, module string, version string, loadData bool) ([]internal.File, error) {
	basePath := moduleVersionBasePath(module, version)
	var files []internal.File

	for object := range s.minioClient.ListObjects(ctx, s.bucket, minio.ListObjectsOptions{
		Prefix:    basePath,
		Recursive: true,
	}) {
		var bts []byte
		if loadData {
			obj, err := s.minioClient.GetObject(ctx, s.bucket, object.Key, minio.GetObjectOptions{})
			if err != nil {
				return nil, err
			}
			bts, err = ioutil.ReadAll(obj)
			if err != nil {
				errResp := minio.ToErrorResponse(err)
				if errResp.Code == "NoSuchKey" {
					return nil, storage.FileNotFound
				}
				return nil, err
			}
		}
		stat, err := s.minioClient.StatObject(ctx, s.bucket, object.Key, minio.StatObjectOptions{})
		if err != nil {
			return nil, err
		}
		files = append(files, internal.File{
			Name:     path.Base(object.Key),
			MimeType: stat.ContentType,
			Data:     bts,
			Uri:      fmt.Sprintf("%s/%s", s.bucket, object.Key),
		})
	}
	return files, nil
}

func (s *Storage) checkObjectExists(ctx context.Context, objectName string) bool {
	_, err := s.minioClient.StatObject(ctx, s.bucket, objectName, minio.StatObjectOptions{})
	if err != nil {
		return false
	}
	return true
}

func (s *Storage) GetModule(ctx context.Context, name string, loadData bool) (*internal.Module, error) {
	objectName := moduleManifestObjectNameFromString(name)
	return s.loadModule(ctx, objectName, loadData)
}

func (s *Storage) GetModules(ctx context.Context, loadData bool) (modules []*internal.Module, err error) {
	for object := range s.minioClient.ListObjects(ctx, s.bucket, minio.ListObjectsOptions{Recursive: true}) {
		isManifestObj := strings.HasSuffix(object.Key, manifestSuffix)
		if isManifestObj {
			var module *internal.Module
			module, err = s.loadModule(ctx, object.Key, loadData)
			if err != nil {
				return nil, err
			}
			modules = append(modules, module)
		}
	}
	return
}

func (s *Storage) loadModule(ctx context.Context, objectName string, loadData bool) (*internal.Module, error) {
	manifestObj, err := s.minioClient.GetObject(ctx, s.bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	var module = &internal.Module{}
	bts, err := ioutil.ReadAll(manifestObj)
	if err != nil {
		errResp := minio.ToErrorResponse(err)
		if errResp.Code == "NoSuchKey" {
			return nil, storage.ModuleNotFound
		}
		return nil, err
	}
	err = json.Unmarshal(bts, module)
	if err != nil {
		return nil, err
	}
	// FIXME: remove this hardcoded dependency MUST be called first after unmarshaling
	module.Init()
	// load available resources
	for _, version := range module.Versions {
		version.Files, _ = s.ModuleResources(ctx, module.Name, version.Version().String(), loadData)
	}

	return module, nil
}

func (s *Storage) Close() error {
	return nil
}

func (s *Storage) setup() {
	exists, err := s.minioClient.BucketExists(context.Background(), s.bucket)
	if err != nil {
		s.logger.Fatal(err)
	}
	if !exists {
		err = s.minioClient.MakeBucket(context.Background(), s.bucket, minio.MakeBucketOptions{})
		if err != nil {
			s.logger.Fatal(err)
		}
		s.logger.Info("Modules bucket setup complete")
	} else {
		s.logger.Info("Modules bucket verified")
	}

	policy, err := s.minioClient.GetBucketPolicy(context.Background(), s.bucket)
	if err != nil {
		return
	}
	if policy == "" {
		readOnlyPolicy := s.readOnlyPolicy()
		s.logger.Info("Updating bucket policy from '", policy, "' to '", readOnlyPolicy, "'")
		err = s.minioClient.SetBucketPolicy(context.Background(), s.bucket, readOnlyPolicy)
		if err != nil {
			s.logger.Fatal(err)
			return
		}
		s.logger.Info("Bucket policy updated")
	}
}

func (s *Storage) readOnlyPolicy() string {
	if customPolicy := os.Getenv("S3_BUCKET_POLICY"); customPolicy != "" {
		return customPolicy
	}
	// minio and AWS default
	return `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::` + s.bucket + `/*"],"Sid": ""}]}`
}

func moduleManifestObjectName(module string) string {
	return module + manifestSuffix
}

func moduleManifestObjectNameFromString(module string) string {
	return module + manifestSuffix
}

func moduleVersionBasePath(module, version string) string {
	return fmt.Sprintf("%s/%s/", module, version)
}
