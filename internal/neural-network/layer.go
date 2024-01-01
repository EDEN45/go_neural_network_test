package neural_network

type Layer struct {
	size     int
	nextSize int
	neurons  []float64
	biases   []float64
	weights  [][]float64
}

func NewLayer(size, nextSize int) Layer {

	weights := make([][]float64, size)

	for i := range weights {
		weights[i] = make([]float64, nextSize)
	}

	return Layer{
		size:     size,
		nextSize: nextSize,
		neurons:  make([]float64, size),
		biases:   make([]float64, size),
		weights:  weights,
	}
}
