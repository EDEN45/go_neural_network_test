package dataset

import (
	"fmt"
	rgbPixel "github.com/EDEN45/go_neural_network_test/internal/rgb-pixel"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var ErrImgDigitsEmptyPath = fmt.Errorf("empty path")
var ErrImgDigitsCountDatasetZero = fmt.Errorf("count dataset zero")
var ErrImgDigitsDatasetNotFound = fmt.Errorf("dataset doen`t find")

// DigitBuff based on dataset in https://github.com/pjreddie
// git: https://github.com/pjreddie/mnist-csv-png

type DigitBuff struct {
	FileName string
	Digit    int8
	Pixels   []float64 // just array of all dots
}

func LoadDigits(filesPath string) ([]DigitBuff, int, error) {
	fileDigits, err := os.ReadDir(filesPath)
	if err != nil {
		return nil, 0, fmt.Errorf("ImgDigits.Load.os.ReadDir, err: %w", err)
	}

	loadedDigits := make([]DigitBuff, 0, len(fileDigits))

	log.Println("START load digits, count: ", len(fileDigits))

	digBuffer := make(chan DigitBuff)

	wg := sync.WaitGroup{}
	wg.Add(len(fileDigits))
	for _, de := range fileDigits {
		go func(de os.DirEntry) {
			defer wg.Done()
			fileName := de.Name()
			digit, err := parseDigit(fileName)
			if err != nil {
				log.Println("Parse number digit, err: ", err.Error())
				return
			}

			rawPixels, err := rgbPixel.ReadPixels(filesPath + string(os.PathSeparator) + fileName)
			if err != nil {
				log.Println("error read image: err", err)
				return
			}

			pixels := make([]float64, 0, len(rawPixels)*len(rawPixels[0]))
			for _, yy := range rawPixels {
				for _, xx := range yy {
					pixels = append(pixels, float64(xx.B)/255) // Get Blue and set only exist color
				}
			}

			digBuffer <- DigitBuff{
				FileName: de.Name(),
				Digit:    digit,
				Pixels:   pixels,
			}
		}(de)
		time.Sleep(500 * time.Microsecond)
	}

	go func() {
		wg.Wait()
		close(digBuffer)
	}()

	for v := range digBuffer {
		loadedDigits = append(loadedDigits, v)
	}

	log.Println("FINISH load digits")

	return loadedDigits, len(fileDigits), nil
}

func LoadPixelsImage(filePath string) ([]float64, error) {
	if filePath == "" {
		return nil, ErrImgDigitsEmptyPath
	}

	rawPixels, err := rgbPixel.ReadPixels(filePath)
	if err != nil {
		log.Println("error read image: err", err)
		return nil, err
	}

	pixels := make([]float64, 0, len(rawPixels)*len(rawPixels[0]))
	for _, yy := range rawPixels {
		for _, xx := range yy {
			pixels = append(pixels, float64(xx.B)/255) // Get Blue and set only exist color
		}
	}

	return pixels, nil
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
