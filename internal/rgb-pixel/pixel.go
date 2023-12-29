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

func PixelFromRGBA(r, g, b, a uint32) Pixel {
	// color can be from 0 to 255
	// we get color [0, 0xffff] or [0, 65535]. If we have 65535, then divide 257 and get 255
	return Pixel{R: r / 257, G: g / 257, B: b / 257, A: a / 257}
}

// ReadPixels from full file path
// note:
//
//	[y][x]Pixel
func ReadPixels(fullFileName string) ([][]Pixel, error) {
	picture, err := os.Open(fullFileName)
	if err != nil {
		return nil, fmt.Errorf("can not open file, err: %w", err)
	}
	defer func() {
		_ = picture.Close()
	}()

	img, _, err := image.Decode(picture)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row,
				PixelFromRGBA(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}
