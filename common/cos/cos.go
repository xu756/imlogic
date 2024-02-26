package cos

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"server/common/config"
)

var _ Client = (*cosClient)(nil)

type Client interface {
	UploadImg(ctx context.Context, objectName, filePath string) (url string, err error)
}

func (c cosClient) UploadImg(ctx context.Context, objectName, filePath string) (url string, err error) {
	uploadInfo, err := c.client.FPutObject(ctx, config.RunData.MinioConfig.BucketName, "images/"+objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s/%s/%s", config.RunData.MinioConfig.Endpoint, config.RunData.MinioConfig.BucketName, uploadInfo.Key), nil
}
