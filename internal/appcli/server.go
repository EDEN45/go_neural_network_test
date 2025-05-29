package appcli

import (
	"encoding/json"
	"fmt"
	"github.com/EDEN45/go_neural_network_test/internal/dataset"
	neural_network "github.com/EDEN45/go_neural_network_test/internal/neural-network"
	rgb_pixel "github.com/EDEN45/go_neural_network_test/internal/rgb-pixel"
	"image"
	"net/http"
	"os"
)

type PredictResult struct {
	Number int `json:"number"`
}

type nnHandler struct {
	nn *neural_network.SimpleNeuralNetwork
}

func (nh *nnHandler) predictHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, "cannot parse form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "cannot get image", http.StatusBadRequest)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, "invalid image", http.StatusBadRequest)
		return
	}
	pixelsRaw := rgb_pixel.ExtractPixels(img)
	pixelNeurons := dataset.PixelsToArr(pixelsRaw)

	rez := nh.nn.FeedForward(pixelNeurons)

	maxRez := -1.0
	maxDig := 0
	for dig, v := range rez {
		if v > maxRez {
			maxRez = v
			maxDig = dig
		}
	}

	resp := PredictResult{Number: maxDig}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func server(addr string, modelPath string) error {
	if addr == "" {
		addr = ":8080"
	}

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

	nh := &nnHandler{nn: nn}

	http.HandleFunc("/predict", nh.predictHandler)

	fmt.Println("🔮 Сервер запущен на http://localhost:8080")
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		return err
	}

	return nil
}
