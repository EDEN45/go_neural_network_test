package main

import (
	"fmt"
	rgb_pixel "github.com/EDEN45/go_neural_network_test/internal/rgb-pixel"
	_ "image/png"
)

func main() {
	p, err := rgb_pixel.ReadPixels("/Users/eden/SDAPFS512/Projects/train-exp/000000-num5.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p[24][5])
}
