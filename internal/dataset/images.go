package dataset

import (
	"fmt"
	rgbPixel "github.com/EDEN45/go_neural_network_test/internal/rgb-pixel"
	"log"
	"os"
	"strconv"
)

var ErrImgDigitsEmptyPath = fmt.Errorf("empty path")
var ErrImgDigitsCountDatasetZero = fmt.Errorf("count dataset zero")
var ErrImgDigitsDatasetNotFound = fmt.Errorf("dataset doen`t find")

// ImgDigits based on dataset in https://github.com/pjreddie
// git: https://github.com/pjreddie/mnist-csv-png
type ImgDigits struct {
	filesPath    string
	countDataset int32

	// map[JustNumber]map[NumberDigit]ImageFile
	images map[int8][]*DigitBuff
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
		return nil, ErrImgDigitsEmptyPath
	}

	if conf.CountDataset == 0 {
		return nil, ErrImgDigitsCountDatasetZero
	}

	return &ImgDigits{
		filesPath:    conf.Path,
		countDataset: conf.CountDataset,
		images:       make(map[int8][]*DigitBuff),
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

	log.Println("START load digits")
	for _, de := range dirEntries {
		fileName := de.Name()
		digit, err := l.parseDigit(fileName)
		if err != nil {
			log.Println("Parse number digit, err: ", err.Error())
			continue
		}

		pixelsRaw, err := rgbPixel.ReadPixels(l.filesPath + string(os.PathSeparator) + fileName)
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

		l.images[digit] = append(l.images[digit], &DigitBuff{
			Digit:  digit,
			Pixels: pixels,
		})
	}
	log.Println("FINISH load digits")

	return nil
}

func (l *ImgDigits) GetDigitDataset(digit int8) ([]*DigitBuff, error) {
	if digit < 0 || digit > 9 {
		return nil, ErrImgDigitsDatasetNotFound
	}

	return l.images[digit], nil
}

func (l *ImgDigits) parseDigit(fileName string) (int8, error) {
	rfn := []rune(fileName)
	// 059700-num4.png remove .png
	rfn = rfn[:len(rfn)-4]
	// set first path - count digit
	//first := rfn[:len(rfn)-5]
	// set digitInName path - number digit
	digitInName := rfn[10:]
	digitRaw, err := strconv.ParseInt(string(digitInName), 10, 8)
	if err != nil {
		return 0, fmt.Errorf("parseDigit, err: %w", err)
	}

	return int8(digitRaw), nil
}
