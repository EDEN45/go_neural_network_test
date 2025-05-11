package appcli

import (
	"encoding/json"
	"fmt"
	"github.com/EDEN45/go_neural_network_test/internal/dataset"
	neural_network "github.com/EDEN45/go_neural_network_test/internal/neural-network"
	"os"
)

func predictImage(modelPath, inputImage string) error {
	modelRaw, err := os.ReadFile(modelPath)
	if err != nil {
		return err
	}

	var model NnModel

	err = json.Unmarshal(modelRaw, &model)
	if err != nil {
		return err
	}

	nn := neural_network.NewSimpleNeuralNetwork(neural_network.SimpleNeuralNetworkConf{
		LearningRate: 0.001,
		FNActivation: neural_network.DefaultActivation,
		FNDerivative: neural_network.DefaultDerivative,
		SizeLayers:   []int{784, 512, 128, 32, 10},
	})

	nn.SetModel(model.LearningRate, model.Layers)

	pixels, err := dataset.LoadPixelsImage(inputImage)
	if err != nil {
		return err
	}

	rez := nn.FeedForward(pixels)

	maxRez := -1.0
	maxDig := 0
	for dig, v := range rez {
		if v > maxRez {
			maxRez = v
			maxDig = dig
		}
	}

	fmt.Println(">> На картинке: ", maxDig)

	return nil
}
