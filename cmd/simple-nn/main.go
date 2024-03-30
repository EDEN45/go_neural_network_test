package main

import (
	"fmt"
	"github.com/EDEN45/go_neural_network_test/internal/dataset"
	neural_network "github.com/EDEN45/go_neural_network_test/internal/neural-network"
	_ "image/png"
	"log"
	"math/rand"
	"time"
)

func main() {
	digits, countFiles, err := dataset.LoadDigits("/Users/eden/SDAPFS512/Projects/train")
	if err != nil {
		log.Println(err)
		return
	}

	rand.Seed(time.Now().UnixNano())

	fmt.Println("---------------------------------")
	fmt.Println(len(digits), countFiles)
	fmt.Println("---------------------------------")

	nn := neural_network.NewSimpleNeuralNetwork(neural_network.SimpleNeuralNetworkConf{
		LearningRate: 0.001,
		FNActivation: neural_network.DefaultActivation,
		FNDerivative: neural_network.DefaultDerivative,
		SizeLayers:   []int{784, 512, 128, 32, 10},
	})

	for i := 1; i < 1000; i++ {
		var (
			right     int     = 0
			errorSum  float64 = 0
			batchSize         = 100
		)

		for bs := 0; bs < batchSize; bs++ {
			imgIndex := int(rand.Float64() * float64(countFiles))
			targets := make([]float64, 10)
			digit := digits[imgIndex].Digit
			targets[digit] = 1

			outputs := nn.FeedForward(digits[imgIndex].Pixels)

			maxDigit := int8(0)
			var maxDigitWeight float64 = -1

			for k := int8(0); k < 10; k++ {
				if outputs[k] > maxDigitWeight {
					maxDigitWeight = outputs[k]
					maxDigit = k
				}
			}

			if digit == maxDigit {
				right++
			}

			for k := int8(0); k < 10; k++ {
				errorSum += (targets[k] - outputs[k]) * (targets[k] - outputs[k])
			}
			nn.BackPropagation(targets)
		}
		fmt.Println("step: ", i, ". correct: ", right, ". error: ", errorSum)
	}

	rez := nn.FeedForward(digits[1].Pixels)

	maxRez := -1.0
	maxDig := 0
	for dig, v := range rez {
		if v > maxRez {
			maxRez = v
			maxDig = dig
		}
	}

	fmt.Println("----------------------------")
	fmt.Println(digits[1].Digit, maxDig)
	fmt.Println(rez)
	fmt.Println(maxRez)
	fmt.Println("----------------------------")

	rez = nn.FeedForward(digits[100].Pixels)
	maxRez = -1.0
	maxDig = 0
	for dig, v := range rez {
		if v > maxRez {
			maxRez = v
			maxDig = dig
		}
	}

	fmt.Println("----------------------------")
	fmt.Println(digits[100].Digit, maxDig)
	fmt.Println(rez)
	fmt.Println(maxRez)
	fmt.Println("----------------------------")

	rez = nn.FeedForward(digits[1000].Pixels)
	maxRez = -1.0
	maxDig = 0
	for dig, v := range rez {
		if v > maxRez {
			maxRez = v
			maxDig = dig
		}
	}

	fmt.Println("----------------------------")
	fmt.Println(digits[1000].Digit, maxDig)
	fmt.Println(rez)
	fmt.Println(maxRez)
	fmt.Println("----------------------------")

	rez = nn.FeedForward(digits[10003].Pixels)
	maxRez = -1.0
	maxDig = 0
	for dig, v := range rez {
		if v > maxRez {
			maxRez = v
			maxDig = dig
		}
	}

	fmt.Println("----------------------------")
	fmt.Println(digits[10003].Digit, maxDig)
	fmt.Println(rez)
	fmt.Println(maxRez)
	fmt.Println("----------------------------")
}
