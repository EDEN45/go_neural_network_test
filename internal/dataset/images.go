package dataset

import (
	"fmt"
	rgb_pixel "github.com/EDEN45/go_neural_network_test/internal/rgb-pixel"
	"log"
	"os"
	"strconv"
)

var ErrImgLdEmptyPath = fmt.Errorf("empty path")
var ErrImgLCountDatasetZero = fmt.Errorf("count dataset zero")

type ImgDigits struct {
	filesPath    string
	countDataset int32

	// map[JustNumber]map[NumberDigit]ImageFile
	images map[int8][]DigitBuff
}

type DigitBuff struct {
	Digit  int8
	Pixels []float64
}

type Conf struct {
	Path         string
	CountDataset int32
	CountPixels  int32
}

func NewImgDigitise(conf Conf) (*ImgDigits, error) {
	if conf.Path == "" {
		return nil, ErrImgLdEmptyPath
	}

	if conf.CountDataset == 0 {
		return nil, ErrImgLCountDatasetZero
	}

	return &ImgDigits{
		filesPath:    conf.Path,
		countDataset: conf.CountDataset,
		images:       make(map[int8][]DigitBuff),
	}, nil
}

func (l *ImgDigits) Load() error {
	dirEntries, err := os.ReadDir(l.filesPath)
	if err != nil {
		return fmt.Errorf("ImgDigits.Load.os.ReadDir, err: %w", err)
	}
	if len(dirEntries) != int(l.countDataset) {
		return fmt.Errorf(
			"ImgDigits.Load, count dirEntries: %d isn`t equils l.countDataset: %d",
			len(dirEntries),
			l.countDataset,
		)
	}

	counter := 0
	for _, de := range dirEntries {
		if counter%100 == 0 {
			fmt.Println(de.Name())

			fileName := de.Name()
			rfn := []rune(fileName)
			// 059700-num4.png remove png
			rfn = rfn[:len(rfn)-4]
			// set first path - count digit
			first := rfn[:len(rfn)-5]
			// set second path - number digit
			second := rfn[10:]

			_ = first
			_ = second

			digitRaw, err := strconv.ParseInt(string(second), 10, 8)
			if err != nil {
				log.Println("Parse number digit, err: ", err.Error())
				continue
			}

			pixelsRaw, err := rgb_pixel.ReadPixels(l.filesPath + string(os.PathSeparator) + fileName)
			if err != nil {
				log.Println("error read image: err", err)
				continue
			}

			pixels := make([]float64, 0, len(pixelsRaw))
			for _, yy := range pixelsRaw {
				for _, xx := range yy {
					pixels = append(pixels, float64(xx.B/255)) // Get Blue and set only exist color
				}
			}

			digit := int8(digitRaw)
			l.images[digit] = append(l.images[digit], DigitBuff{
				Digit:  digit,
				Pixels: pixels,
			})
		}
		counter++
	}

	return nil
}
