package main

import (
	"github.com/EDEN45/go_neural_network_test/internal/appcli"
	_ "image/png"
)

func main() {
	app := appcli.NewCli()
	app.Run()

	//digits, countFiles, err := dataset.LoadDigits("/Users/eden/SDAPFS512/Projects/train")
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//fmt.Println("---------------------------------")
	//fmt.Println(len(digits), countFiles)
	//fmt.Println("---------------------------------")
	//
	//nn := neural_network.NewSimpleNeuralNetwork(neural_network.SimpleNeuralNetworkConf{
	//	LearningRate: 0.001,
	//	FNActivation: neural_network.DefaultActivation,
	//	FNDerivative: neural_network.DefaultDerivative,
	//	SizeLayers:   []int{784, 512, 128, 32, 10},
	//})
	//
	//rez := nn.FeedForward(digits[1].Pixels)
	//
	//maxRez := -1.0
	//maxDig := 0
	//for dig, v := range rez {
	//	if v > maxRez {
	//		maxRez = v
	//		maxDig = dig
	//	}
	//}
	//
	//fmt.Println("----------------------------")
	//fmt.Println(digits[1].Digit, maxDig)
	//fmt.Println(rez)
	//fmt.Println(maxRez)
	//fmt.Println("----------------------------")
	//
	//rez = nn.FeedForward(digits[100].Pixels)
	//maxRez = -1.0
	//maxDig = 0
	//for dig, v := range rez {
	//	if v > maxRez {
	//		maxRez = v
	//		maxDig = dig
	//	}
	//}
	//
	//fmt.Println("----------------------------")
	//fmt.Println(digits[100].Digit, maxDig)
	//fmt.Println(rez)
	//fmt.Println(maxRez)
	//fmt.Println("----------------------------")
	//
	//rez = nn.FeedForward(digits[1000].Pixels)
	//maxRez = -1.0
	//maxDig = 0
	//for dig, v := range rez {
	//	if v > maxRez {
	//		maxRez = v
	//		maxDig = dig
	//	}
	//}
	//
	//fmt.Println("----------------------------")
	//fmt.Println(digits[1000].Digit, maxDig)
	//fmt.Println(rez)
	//fmt.Println(maxRez)
	//fmt.Println("----------------------------")
	//
	//rez = nn.FeedForward(digits[10003].Pixels)
	//maxRez = -1.0
	//maxDig = 0
	//for dig, v := range rez {
	//	if v > maxRez {
	//		maxRez = v
	//		maxDig = dig
	//	}
	//}
	//
	//fmt.Println("----------------------------")
	//fmt.Println(digits[10003].Digit, maxDig)
	//fmt.Println(rez)
	//fmt.Println(maxRez)
	//fmt.Println("----------------------------")
}
