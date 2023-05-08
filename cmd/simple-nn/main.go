package main

import (
	"github.com/EDEN45/go_neural_network_test/internal/dataset"
	_ "image/png"
	"log"
)

func main() {
	imgLd, err := dataset.NewImgDigitise(dataset.Conf{
		Path:         "/home/eden/sdb500/eden/projects/simple-nn/digist-dataset/mnist_train/train",
		CountDataset: 60000,
	})
	if err != nil {
		log.Println("Error init image loader")
		return
	}

	if err := imgLd.Load(); err != nil {
		log.Println(err)
		return
	}

}
