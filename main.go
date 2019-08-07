package main

import (
	"fmt"

	"github.com/Drakmord2/go-ga/controller"
	"github.com/Drakmord2/go-ga/util"
)

func main() {
	config := util.Config{
		Population:    10,
		CrossoverRate: 0.3,
		MutationRate:  0.005,
		MaxIteration:  2,
	}

	fmt.Println("\t- Go-GA -")
	controller.GeneticAlgorithm(config)
}
