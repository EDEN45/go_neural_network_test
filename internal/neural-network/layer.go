package neural_network

type Layer struct {
	Size int `json:"size"`
	//nextSize int
	Neurons []float64   `json:"neurons"`
	Biases  []float64   `json:"biases"`
	Weights [][]float64 `json:"weights"`
}

func NewLayer(size int) Layer {
	//for i := range Weights {
	//	Weights[i] = make([]float64, Size)
	//}

	return Layer{
		Size: size,
		//nextSize: nextSize,
		Neurons: make([]float64, size),
		Biases:  make([]float64, size),
		Weights: make([][]float64, size),
	}
}
