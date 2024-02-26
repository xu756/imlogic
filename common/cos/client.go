package cos

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"server/common/config"
)

type cosClient struct {
	client *minio.Client
}

func NewClient() (Client, error) {
	minioConfig := config.RunData.MinioConfig
	minioClient, err := minio.New(minioConfig.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioConfig.AccessKeyID, minioConfig.SecretAccessKey, ""),
		Secure: minioConfig.UseSSL,
	})
	if err != nil {
		return nil, err
	}
	return &cosClient{
		client: minioClient,
	}, nil

}
