package appcli

import neural_network "github.com/EDEN45/go_neural_network_test/internal/neural-network"

type NnModel struct {
	LearningRate float64                `json:"learningRate"`
	Layers       []neural_network.Layer `json:"layers"`
}
