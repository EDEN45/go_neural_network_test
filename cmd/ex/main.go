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

	pixels := make([]float64, len(p)*len(p[0]))
	i := 0
	for _, yy := range p {
		for _, xx := range yy {
			px := float64(xx.B) / 255.0
			pixels[i] = px // Get Blue and set only exist color
			if px != 0 {
				fmt.Println(i, xx, px)
				return
			}
			i++
		}
	}

	fmt.Println("-------------------")

	for i, v := range pixels {
		if v == 0 {
			continue
		}

		fmt.Println(i, v)
	}
}
