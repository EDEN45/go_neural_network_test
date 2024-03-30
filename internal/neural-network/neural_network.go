package neural_network

import (
	"fmt"
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
	layers := make([]Layer, lenSizeLayers)

	for i := 0; i < lenSizeLayers; i++ {
		var nextSize = 0

		if i < lenSizeLayers-1 {
			nextSize = conf.SizeLayers[i+1]
		}

		layers[i] = NewLayer(conf.SizeLayers[i])

		for j := 0; j < conf.SizeLayers[i]; j++ {
			//TODO -why?????
			layers[i].biases[j] = rand.Float64()*2.0 - 1.0
			layers[i].weights[j] = make([]float64, nextSize)
			for k := 0; k < nextSize; k++ {
				layers[i].weights[j][k] = rand.Float64()*2.0 - 1.0
			}
		}
	}

	fmt.Println("--------------")
	for _, v := range layers {
		fmt.Println(v.size)
	}
	fmt.Println("--------------")

	return &SimpleNeuralNetwork{
		learningRate: conf.LearningRate,
		fnActivation: conf.FNActivation,
		fnDerivative: conf.FNDerivative,
		layers:       layers,
	}
}

func (s *SimpleNeuralNetwork) FeedForward(inputs []float64) []float64 {
	copy(s.layers[0].neurons, inputs)

	for i := 1; i < len(s.layers); i++ {
		l := &s.layers[i-1]
		l1 := &s.layers[i]

		for j := 0; j < l1.size; j++ {
			l1.neurons[j] = 0

			for k := 0; k < l.size; k++ {
				if len(l.weights) == 0 {
					l.weights = make([][]float64, l.size)
					fmt.Println("FIRST", i, j, k)
				}

				if len(l.weights[k]) == 0 {
					l.weights = make([][]float64, l1.size)
					fmt.Println("SECOND", i, j, k)
				}
				a2 := l.weights[k][j]

				a1 := l.neurons[k]
				l1.neurons[j] += a1 * a2
			}
			l1.neurons[j] += l1.biases[j]
			l1.neurons[j] = s.fnActivation(l1.neurons[j])
		}
	}
	return s.layers[len(s.layers)-1].neurons
}

func (s *SimpleNeuralNetwork) BackPropagation(targets []float64) {
	commonErrs := make([]float64, s.layers[len(s.layers)-1].size)

	for i := 0; i < s.layers[len(s.layers)-1].size; i++ {
		commonErrs[i] = targets[i] - s.layers[len(s.layers)-1].neurons[i]
	}

	for k := len(s.layers) - 2; k >= 0; k-- {
		l := &s.layers[k] //TODO что это?
		l1 := &s.layers[k+1]
		nextErrs := make([]float64, l.size)
		gradients := make([]float64, l1.size)

		for i := 0; i < l1.size; i++ {
			gradients[i] = commonErrs[i] * s.fnDerivative(l1.neurons[i])
			gradients[i] *= s.learningRate
		}

		deltas := make([][]float64, l1.size)
		for i := 0; i < l1.size; i++ {
			deltas[i] = make([]float64, l.size)
			for j := 0; j < l.size; j++ {
				deltas[i][j] = gradients[i] * l.neurons[j]
			}
		}

		for i := 0; i < l.size; i++ {
			nextErrs[i] = 0
			for j := 0; j < l1.size; j++ {
				nextErrs[i] += l.weights[i][j] * commonErrs[j]
			}
		}

		//commonErrs = make([]float64, s.layers[k].size)
		//for i, v := range nextErrs {
		//	commonErrs[i] = v
		//}

		//commonErrs = nextErrs

		commonErrs = make([]float64, l.size)
		copy(commonErrs, nextErrs)

		weightsNew := make([][]float64, len(l.weights))

		for i := 0; i < l1.size; i++ {

			for j := 0; j < l.size; j++ {
				if len(weightsNew[j]) == 0 {
					weightsNew[j] = make([]float64, len(l.weights[0]))
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
