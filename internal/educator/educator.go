package educator

import (
	"context"
	"fmt"
	"github.com/EDEN45/go_neural_network_test/internal/dataset"
	neural_network "github.com/EDEN45/go_neural_network_test/internal/neural-network"
	"math/rand/v2"
)

func Learn(
	ctx context.Context,
	nn *neural_network.SimpleNeuralNetwork,
	digits []dataset.DigitBuff,
	countFiles int,
	trainingSteps int,
	batchSize int,
) {
	if trainingSteps == 0 {
		trainingSteps = 1000
	}

	if batchSize == 0 {
		batchSize = 100
	}

	for i := 1; i < trainingSteps; i++ {
		var (
			right    int     = 0
			errorSum float64 = 0
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
}
