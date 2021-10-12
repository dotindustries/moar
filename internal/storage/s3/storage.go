package s3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/dotindustries/moar/internal"
	"github.com/dotindustries/moar/internal/storage"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

const modulesBucket = "modules"
const manifestSuffix = "/manifest.json"

type Storage struct {
	endpoint    string
	minioClient *minio.Client
	logger      *logrus.Entry
}

func New(endpoint string) *Storage {
	logger := logrus.WithField("op", "storage")
	if endpoint == "" {
		endpoint = "localhost:9000"
	}
	accessKeyID := "minio"
	secretAccessKey := "minio123"
	useSSL := false
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		logger.Fatalln(err)
	}

	s := &Storage{
		endpoint:    endpoint,
		minioClient: minioClient,
		logger:      logger,
	}
	s.setup()
	return s
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
		modulesBucket,
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
		modulesBucket,
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
				rollBackErr := s.minioClient.RemoveObject(ctx, modulesBucket, objectName, minio.RemoveObjectOptions{ForceDelete: true})
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
			modulesBucket,
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
		modulesBucket,
		basePath,
		minio.RemoveObjectOptions{
			ForceDelete: true,
		},
	)
}

func (s *Storage) ModuleResources(ctx context.Context, module string, version string, loadData bool) ([]internal.File, error) {
	basePath := moduleVersionBasePath(module, version)
	var files []internal.File

	for object := range s.minioClient.ListObjects(ctx, modulesBucket, minio.ListObjectsOptions{
		Prefix:    basePath,
		Recursive: true,
	}) {
		var bts []byte
		if loadData {
			obj, err := s.minioClient.GetObject(ctx, modulesBucket, object.Key, minio.GetObjectOptions{})
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
		stat, err := s.minioClient.StatObject(ctx, modulesBucket, object.Key, minio.StatObjectOptions{})
		if err != nil {
			return nil, err
		}
		files = append(files, internal.File{
			Name:     path.Base(object.Key),
			MimeType: stat.ContentType,
			Data:     bts,
			Uri:      fmt.Sprintf("%s/%s", modulesBucket, object.Key),
		})
	}
	return files, nil
}

func (s *Storage) checkObjectExists(ctx context.Context, objectName string) bool {
	_, err := s.minioClient.StatObject(ctx, modulesBucket, objectName, minio.StatObjectOptions{})
	if err != nil {
		return false
	}
	return true
}

func (s *Storage) GetModule(ctx context.Context, name string, loadData bool) (*internal.Module, error) {
	objectName := moduleManifestObjectNameFromString(name)
	manifestObj, err := s.minioClient.GetObject(ctx, modulesBucket, objectName, minio.GetObjectOptions{})
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
	exists, err := s.minioClient.BucketExists(context.Background(), modulesBucket)
	if err != nil {
		s.logger.Fatal(err)
	}
	if !exists {
		err = s.minioClient.MakeBucket(context.Background(), modulesBucket, minio.MakeBucketOptions{})
		if err != nil {
			s.logger.Fatal(err)
		}
		readOnlyPolicy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::` + modulesBucket + `/*"],"Sid": ""}]}`
		err = s.minioClient.SetBucketPolicy(context.Background(), modulesBucket, readOnlyPolicy)
		if err != nil {
			s.logger.Fatal(err)
			return
		}
		s.logger.Info("Modules bucket setup complete")
	} else {
		s.logger.Info("Modules bucket verified")
	}
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
