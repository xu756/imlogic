package cos

import (
	"context"
	"server/common/config"
	"testing"
)

func TestCosClient_UploadFile(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	client, err := NewClient()
	if err != nil {
		t.Error(err)
	}
	url, err := client.UploadImg(context.Background(), "1.png", "/Users/xu756/Downloads/iShot_2024-01-26_00.53.42.png")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(url)
}
