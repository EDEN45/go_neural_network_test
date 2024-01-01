package neural_network

import (
	"math"
	"math/rand"
)

// TODO need find, why?
var DefaultActivation ModFn = func(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// TODO need find, why?
var DefaultDerivative ModFn = func(y float64) float64 {
	return y * (1 - y)
}

type ModFn func(n float64) float64

type SimpleNeuralNetwork struct {
	learningRate float64
	layers       []Layer // we have more layers, first layer is it input, last layer is output result calculating neural network
	fnActivation ModFn
	fnDerivative ModFn
}

type SimpleNeuralNetworkConf struct {
	LearningRate float64
	FNActivation ModFn
	FNDerivative ModFn
	SizeLayers   []int
}

func NewSimpleNeuralNetwork(conf SimpleNeuralNetworkConf) *SimpleNeuralNetwork {
	lenSizeLayers := len(conf.SizeLayers)

	// we create neurons layers from of count layers
	layers := make([]Layer, len(conf.SizeLayers))

	for i := 0; i < len(conf.SizeLayers); i++ {
		var nextSize = 0

		if i < lenSizeLayers-1 {
			nextSize = conf.SizeLayers[i+1]
		}
		layers[i] = NewLayer(conf.SizeLayers[i], nextSize)

		for j := 0; j < conf.SizeLayers[i]; j++ {
			//TODO -why?????
			layers[i].biases[j] = rand.Float64()*2.0 - 1.0

			for k := 0; k < nextSize; k++ {
				if layers[i].weights[j] == nil {
					layers[i].weights[j] = make([]float64, nextSize)
				}
				layers[i].weights[j][k] = rand.Float64()*2.0 - 1.0
			}
		}
	}

	return &SimpleNeuralNetwork{
		learningRate: conf.LearningRate,
		fnActivation: conf.FNActivation,
		fnDerivative: conf.FNDerivative,
		layers:       layers,
	}
}

func (s *SimpleNeuralNetwork) FeedForward(input []float64) []float64 {
	for i, v := range input {
		s.layers[0].neurons[i] = v
	}

	for i := 1; i < len(s.layers); i++ {
		l := s.layers[i-1]
		l1 := s.layers[i]

		for j := 0; j < l1.size; j++ {
			l1.neurons[j] = 0

			for k := 0; k < l.size; k++ {
				l1.neurons[j] += l.neurons[k] * l.weights[k][j]
			}
			l1.neurons[j] += l1.biases[j]
			l1.neurons[j] = s.fnActivation(l1.neurons[j])
		}
	}
	return s.layers[len(s.layers)-1].neurons
}

func (s *SimpleNeuralNetwork) BackPropagation(targets []float64) {
	fErrors := make([]float64, s.layers[len(s.layers)-1].size)

	for i := 0; i < s.layers[len(s.layers)-1].size; i++ {
		fErrors[i] = targets[i] - s.layers[len(s.layers)-1].neurons[i]
	}

	for k := len(s.layers) - 2; k >= 0; k-- {
		l := s.layers[k] //TODO что это?
		l1 := s.layers[k+1]
		errsNext := make([]float64, l.size)
		gradients := make([]float64, l1.size)

		for i := 0; i < l1.size; i++ {
			gradients[i] = fErrors[i] * s.fnDerivative(s.layers[k+1].neurons[i])
			gradients[i] *= s.learningRate
		}

		deltas := make(map[int][]float64)
		for i := 0; i < l1.size; i++ {
			for j := 0; j < l.size; j++ {
				if deltas[i] == nil {
					deltas[i] = make([]float64, l.size)
				}
				deltas[i][j] = gradients[i] * l.neurons[j]
			}
		}

		for i := 0; i < l.size; i++ {
			errsNext[i] = 0
			for j := 0; j < l1.size; j++ {
				errsNext[i] += l.weights[i][j] * errsNext[j]
			}
		}

		fErrors = make([]float64, l.size)
		for i, v := range errsNext {
			fErrors[i] = v
		}

		var weightsNew [][]float64
		weightsNew = make([][]float64, len(l.weights))

		for i := 0; i < l1.size; i++ {
			for j := 0; j < l.size; j++ {
				if weightsNew[j] == nil {
					weightsNew[j] = make([]float64, l1.size)
				}
				weightsNew[j][i] = l.weights[j][i] + deltas[i][j]
			}
		}

		l.weights = weightsNew
		for i := 0; i < l1.size; i++ {
			l1.biases[i] += gradients[i]
		}
	}
}
