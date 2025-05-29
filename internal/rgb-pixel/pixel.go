package rgb_pixel

import (
	"fmt"
	"image"
	"image/color"
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
	return Pixel{R: r & 0xff, G: g & 0xff, B: b & 0xff, A: a & 0xff}
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

	return ExtractPixels(img), nil
}

func ExtractPixels(img image.Image) [][]Pixel {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			oldPixel := img.At(x, y)
			newColor := color.GrayModel.Convert(oldPixel)
			row = append(row,
				PixelFromRGBA(newColor.RGBA()),
			)
		}
		pixels = append(pixels, row)
	}

	return pixels
}
