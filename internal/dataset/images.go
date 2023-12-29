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

// DigitBuff based on dataset in https://github.com/pjreddie
// git: https://github.com/pjreddie/mnist-csv-png

type DigitBuff struct {
	Digit  int8
	Pixels []float64
}

func LoadDigits(filesPath string) ([]DigitBuff, int, error) {
	fileDigits, err := os.ReadDir(filesPath)
	if err != nil {
		return nil, 0, fmt.Errorf("ImgDigits.Load.os.ReadDir, err: %w", err)
	}

	loadedDigits := make([]DigitBuff, 0, len(fileDigits))

	log.Println("START load digits, count: ", len(fileDigits))

	for i, de := range fileDigits {
		fmt.Printf("%.2f %s \n", float64(i)/float64(len(fileDigits))*100, " %")
		fileName := de.Name()
		digit, err := parseDigit(fileName)
		if err != nil {
			log.Println("Parse number digit, err: ", err.Error())
			continue
		}

		rawPixels, err := rgbPixel.ReadPixels(filesPath + string(os.PathSeparator) + fileName)
		if err != nil {
			log.Println("error read image: err", err)
			continue
		}

		pixels := make([]float64, 0, len(rawPixels))
		for _, yy := range rawPixels {
			for _, xx := range yy {
				pixels = append(pixels, float64(xx.B/255)) // Get Blue and set only exist color
			}
		}

		loadedDigits = append(loadedDigits, DigitBuff{
			Digit:  digit,
			Pixels: pixels,
		})

	}
	log.Println("FINISH load digits")

	return loadedDigits, len(fileDigits), nil
}

func parseDigit(fileName string) (int8, error) {
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
