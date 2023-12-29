package main

import (
	"fmt"
	"github.com/EDEN45/go_neural_network_test/internal/dataset"
	_ "image/png"
	"log"
)

func main() {
	digits, countFiles, err := dataset.LoadDigits("/Users/eden/SDAPFS512/Projects/train")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(len(digits), countFiles)
}
