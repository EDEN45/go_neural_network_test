package neural_network

type Layer struct {
	size     int
	nextSize int
	neurons  []float64
	biases   []float64
	weights  map[int][]float64
}

func NewLayer(size, nextSize int) *Layer {
	return &Layer{
		size:     size,
		nextSize: nextSize,
		neurons:  make([]float64, 0, size),
		biases:   make([]float64, 0, size),
		weights:  make(map[int][]float64),
	}
}
