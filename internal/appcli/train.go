package appcli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/EDEN45/go_neural_network_test/internal/dataset"
	"github.com/EDEN45/go_neural_network_test/internal/educator"
	neural_network "github.com/EDEN45/go_neural_network_test/internal/neural-network"
	"os"
	"time"
)

// Подразумевается, что у тебя есть реализация этих функций:
func trainAndSaveModel(inputPath, output string) error {
	if inputPath == "" {
		inputPath = "/Users/eden/SDAPFS512/Projects/train"
	}

	if output == "" {
		t := time.Now()
		output = "/Volumes/SDAPFS512/Projects/github.com/EDEN45/go_neural_network_test/model" + t.String() + ".json"
	}

	digits, countFiles, err := dataset.LoadDigits(inputPath)
	if err != nil {
		return err
	}

	fmt.Println("---------------------------------")
	fmt.Println(len(digits), countFiles)
	fmt.Println("---------------------------------")

	nn := neural_network.NewSimpleNeuralNetwork(neural_network.SimpleNeuralNetworkConf{
		LearningRate: 0.001,
		FNActivation: neural_network.DefaultActivation,
		FNDerivative: neural_network.DefaultDerivative,
		SizeLayers:   []int{784, 512, 128, 32, 10},
	})

	educator.Learn(
		context.Background(),
		nn,
		digits,
		countFiles,
		1000,
		100,
	)

	learningRate, layers := nn.GetModel()

	dto := NnModel{
		LearningRate: learningRate,
		Layers:       layers,
	}

	b, err := json.Marshal(dto)
	if err != nil {
		return err
	}

	err = os.WriteFile(output, b, 0644)
	if err != nil {
		return err
	}

	fmt.Println(">> Модель сохранена в:", output)

	return nil
}
