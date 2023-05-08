package rgb_pixel

import (
	"fmt"
	"image"
	"os"
)

type Pixel struct {
	R uint32 // red
	G uint32 // green
	B uint32 // blue
	A uint32 // alpha channel
}

func PixelFromRGBA(r, g, b, a uint32) *Pixel {
	return &Pixel{R: r / 257, G: g / 257, B: b / 257, A: a / 257}
}

func ReadPixels(fullFileName string) ([][]*Pixel, error) {
	picture, err := os.Open(fullFileName)
	if err != nil {
		return nil, fmt.Errorf("can not open file, err: %w", err)
	}
	defer func() {
		_ = picture.Close()
	}()

	fmt.Println(fullFileName)

	img, _, err := image.Decode(picture)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var pixels [][]*Pixel
	for y := 0; y < height; y++ {
		var row []*Pixel
		for x := 0; x < width; x++ {
			row = append(row, PixelFromRGBA(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}
