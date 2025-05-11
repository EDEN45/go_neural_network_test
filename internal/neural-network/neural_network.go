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

	// we create Neurons layers from of count layers
	layers := make([]Layer, lenSizeLayers)

	for i := 0; i < lenSizeLayers; i++ {
		var nextSize = 0

		if i < lenSizeLayers-1 {
			nextSize = conf.SizeLayers[i+1]
		}

		layers[i] = NewLayer(conf.SizeLayers[i])

		for j := 0; j < conf.SizeLayers[i]; j++ {
			//TODO -why?????
			layers[i].Biases[j] = rand.Float64()*2.0 - 1.0
			layers[i].Weights[j] = make([]float64, nextSize)
			for k := 0; k < nextSize; k++ {
				layers[i].Weights[j][k] = rand.Float64()*2.0 - 1.0
			}
		}
	}

	fmt.Println("--------------")
	for _, v := range layers {
		fmt.Println(v.Size)
	}
	fmt.Println("--------------")

	return &SimpleNeuralNetwork{
		learningRate: conf.LearningRate,
		fnActivation: conf.FNActivation,
		fnDerivative: conf.FNDerivative,
		layers:       layers,
	}
}

func (s *SimpleNeuralNetwork) GetModel() (learningRate float64, layers []Layer) {
	learningRate = s.learningRate
	layers = s.layers
	return
}

func (s *SimpleNeuralNetwork) SetModel(learningRate float64, layers []Layer) {
	s.learningRate = learningRate
	s.layers = layers
	s.fnDerivative = DefaultDerivative
	s.fnActivation = DefaultActivation
}

func (s *SimpleNeuralNetwork) FeedForward(inputs []float64) []float64 {
	copy(s.layers[0].Neurons, inputs)

	for i := 1; i < len(s.layers); i++ {
		l := &s.layers[i-1]
		l1 := &s.layers[i]

		for j := 0; j < l1.Size; j++ {
			l1.Neurons[j] = 0

			for k := 0; k < l.Size; k++ {
				//if len(l.Weights) == 0 {
				//	l.Weights = make([][]float64, l.Size)
				//	fmt.Println("FIRST", i, j, k)
				//}
				//
				//if len(l.Weights[k]) == 0 {
				//	l.Weights = make([][]float64, l1.Size)
				//	fmt.Println("SECOND", i, j, k)
				//}
				a2 := l.Weights[k][j]

				a1 := l.Neurons[k]
				l1.Neurons[j] += a1 * a2
			}
			l1.Neurons[j] += l1.Biases[j]
			l1.Neurons[j] = s.fnActivation(l1.Neurons[j])
		}
	}
	return s.layers[len(s.layers)-1].Neurons
}

func (s *SimpleNeuralNetwork) BackPropagation(targets []float64) {
	commonErrs := make([]float64, s.layers[len(s.layers)-1].Size)

	for i := 0; i < s.layers[len(s.layers)-1].Size; i++ {
		commonErrs[i] = targets[i] - s.layers[len(s.layers)-1].Neurons[i]
	}

	for k := len(s.layers) - 2; k >= 0; k-- {
		l := &s.layers[k] //TODO что это?
		l1 := &s.layers[k+1]
		nextErrs := make([]float64, l.Size)
		gradients := make([]float64, l1.Size)

		for i := 0; i < l1.Size; i++ {
			gradients[i] = commonErrs[i] * s.fnDerivative(l1.Neurons[i])
			gradients[i] *= s.learningRate
		}

		deltas := make([][]float64, l1.Size)
		for i := 0; i < l1.Size; i++ {
			deltas[i] = make([]float64, l.Size)
			for j := 0; j < l.Size; j++ {
				deltas[i][j] = gradients[i] * l.Neurons[j]
			}
		}

		for i := 0; i < l.Size; i++ {
			nextErrs[i] = 0
			for j := 0; j < l1.Size; j++ {
				nextErrs[i] += l.Weights[i][j] * commonErrs[j]
			}
		}

		//commonErrs = make([]float64, s.layers[k].Size)
		//for i, v := range nextErrs {
		//	commonErrs[i] = v
		//}

		//commonErrs = nextErrs

		commonErrs = make([]float64, l.Size)
		copy(commonErrs, nextErrs)

		weightsNew := make([][]float64, len(l.Weights))

		for i := 0; i < l1.Size; i++ {

			for j := 0; j < l.Size; j++ {
				if len(weightsNew[j]) == 0 {
					weightsNew[j] = make([]float64, len(l.Weights[0]))
				}

				weightsNew[j][i] = l.Weights[j][i] + deltas[i][j]
			}
		}

		l.Weights = weightsNew
		for i := 0; i < l1.Size; i++ {
			l1.Biases[i] += gradients[i]
		}
	}
}
