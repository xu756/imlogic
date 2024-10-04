package cos

import (
	"context"
	"imlogic/common/config"
	"imlogic/internal/xerr"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Cos struct {
	client *minio.Client
	bucket string
	cosUrl string
}

// 初始化对象存储
func NewCos() *Cos {
	mConfig := config.RunData().Minio
	// Initialize minio client object.
	minioClient, err := minio.New(mConfig.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(mConfig.AccessKeyID, mConfig.SecretAccessKey, ""),
		Secure: mConfig.UseSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return &Cos{
		client: minioClient,
		bucket: mConfig.Bucket,
		cosUrl: config.RunData().CosUrl,
	}
}

// 上传图片
func (c *Cos) UploadImage(ctx context.Context, imgName, imgPath string) (url string, err error) {
	// 获取图片格式(匹配) 设置后缀
	const fileDir = "/images/"
	imgPath = CompressImg(imgPath)
	file, err := c.client.FPutObject(ctx, c.bucket, fileDir+imgName, imgPath, minio.PutObjectOptions{})
	if err != nil {
		return "", xerr.UploadImageErr(err, "上传图片失败 图片路径:%s", imgPath)
	}
	url = config.RunData().CosUrl + file.Key
	return
}
