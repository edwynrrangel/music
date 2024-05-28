package bucket

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type minIO struct {
	client *minio.Client
}

type minIObuilder struct {
	endpoint  string
	accessKey string
	secretKey string
	useSSL    bool
}

// NewBuilder creates a new builder with the provided endpoint, access key, and secret key.
func NewBuilder(endpoint, accessKey, secretKey string) *minIObuilder {
	return &minIObuilder{
		endpoint:  endpoint,
		accessKey: accessKey,
		secretKey: secretKey,
	}
}

// WithSSL sets the useSSL flag to true.
func (b *minIObuilder) WithSSL() *minIObuilder {
	b.useSSL = true
	return b
}

func (b *minIObuilder) GetClient() (*minIO, error) {
	options := &minio.Options{
		Creds: credentials.NewStaticV4(b.accessKey, b.secretKey, ""),
	}
	if b.useSSL {
		options.Secure = true
	}
	client, err := minio.New(b.endpoint, options)
	if err != nil {
		return nil, err
	}
	return &minIO{client: client}, nil
}

func (m *minIO) DownloadFile(ctx context.Context, bucketName, fileName string) (io.ReadCloser, error) {
	return m.client.GetObject(ctx, bucketName, fmt.Sprintf("%s.mp3", fileName), minio.GetObjectOptions{})
}
