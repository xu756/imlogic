package cos

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

// CompressImg 压缩图片
func CompressImg(imgPath string) string {
	file, err := os.Open(imgPath)
	if err != nil {
		return imgPath
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return imgPath
	}

	// 使用 nfnt/resize 库来调整图像大小
	newImg := resize.Resize(750, 0, img, resize.Lanczos3)

	// 创建一个新的文件来保存压缩后的图片
	output := strings.TrimSuffix(imgPath, ".") + "_compressed." + format
	out, err := os.Create(output)
	if err != nil {
		return imgPath
	}
	defer out.Close()

	// 根据图片格式进行编码并保存
	switch format {
	case "jpeg":
		opt := jpeg.Options{
			Quality: 80,
		}
		err = jpeg.Encode(out, newImg, &opt)
	case "png":
		err = png.Encode(out, newImg)
	default:
		return imgPath
	}

	if err != nil {
		return imgPath
	}
	return output
}
