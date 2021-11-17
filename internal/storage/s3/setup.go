package s3

import (
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

const defaultBucket = "modules"
const defaultEndpoint = "localhost:9000"

func New(endpoint string) *Storage {
	logger := logrus.WithField("op", "storage")
	endpoint = validateEndpoint(endpoint)
	creds := collectCredentials()
	useSSL := false
	if os.Getenv("S3_USE_SSL") != "" {
		useSSL = true
	}
	minioClient, err := minio.New(endpoint,
		&minio.Options{
			Creds:  creds,
			Secure: useSSL,
		},
	)
	if err != nil {
		logger.Fatalln(err)
	}

	s := &Storage{
		minioClient: minioClient,
		logger:      logger,
		bucket:      bucket(),
	}
	s.setup()
	return s
}

func collectCredentials() *credentials.Credentials {
	accessKeyID := os.Getenv("S3_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("S3_SECRET_ACCESS_KEY")
	sessionToken := os.Getenv("AWS_SESSION_TOKEN")
	return credentials.NewStaticV4(accessKeyID, secretAccessKey, sessionToken)
}

func bucket() string {
	bucket := os.Getenv("S3_BUCKET")
	if bucket == "" {
		bucket = defaultBucket
	}
	return bucket
}

func validateEndpoint(endpoint string) string {
	if endpoint == "" {
		endpoint = os.Getenv("S3_ENDPOINT_URL")
		if endpoint == "" {
			endpoint = defaultEndpoint
		}
	}
	logrus.Infof("Using S3 at %s", endpoint)
	return endpoint
}
