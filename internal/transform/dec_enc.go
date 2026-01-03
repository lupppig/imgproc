package transform

import (
	"errors"
	"image"
	"os"

	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"

	"github.com/chai2010/webp"
)

func Decode(path string) (image.Image, string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()

	return image.Decode(f)
}

func Encode(outPath string, img image.Image, format string, quality int) error {
	out, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer out.Close()

	if quality == 0 {
		quality = 85
	}

	switch format {
	case "jpeg", "jpg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: quality})
	case "png":
		return png.Encode(out, img)
	case "webp":
		return webp.Encode(out, img, &webp.Options{Lossless: false, Quality: float32(quality)})
	case "avif":
		return errors.New("AVIF encoding not implemented; use avifenc or libavif")
	default:
		return errors.New("unsupported format")
	}
}
