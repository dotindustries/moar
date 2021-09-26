package s3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/nadilas/moar/internal"
	"github.com/nadilas/moar/internal/storage"
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

func (s *Storage) PutModule(ctx context.Context, module *internal.Module) error {
	bts, err := json.Marshal(module)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(bts)
	objectName := moduleManifestObjectName(module)
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
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) PutVersion(ctx context.Context, module string, version string, data []byte) error {
	r := bytes.NewReader(data)
	objectName := moduleVersionObjectName(module, version)
	_, err := s.minioClient.PutObject(
		ctx,
		modulesBucket,
		objectName,
		r,
		r.Size(),
		minio.PutObjectOptions{
			ContentType: "text/javascript",
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) RemoveVersion(ctx context.Context, module string, version string) error {
	objectName := moduleVersionObjectName(module, version)
	err := s.minioClient.RemoveObject(
		ctx,
		modulesBucket,
		objectName,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UriForModule(ctx context.Context, module string, version string) (string, error) {
	return fmt.Sprintf("%s/%s", modulesBucket, moduleVersionObjectName(module, version)), nil
}

func (s *Storage) GetModule(ctx context.Context, name string) (*internal.Module, error) {
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
	// make sure to init versions
	module.Init()

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

func moduleManifestObjectName(module *internal.Module) string {
	return module.Name + manifestSuffix
}

func moduleManifestObjectNameFromString(module string) string {
	return module + manifestSuffix
}

func moduleVersionObjectName(module string, version string) string {
	return fmt.Sprintf("%s/%s@%s.umd.js", module, module, version)
}
