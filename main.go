package main

import (
	"fmt"

	"github.com/Drakmord2/go-ga/controller"
	"github.com/Drakmord2/go-ga/util"
)

func main() {
	config := util.Config{
		MaxIteration:  2,
		Population:    10,
		Parameters:    []string{"b7", "b6", "b5", "b4", "b3", "b2", "b1", "b0"},
		CrossoverRate: 0.3,
		MutationRate:  0.005,
	}

	fmt.Println("\n\t- Go-GA -")
	solution := controller.GeneticAlgorithm(config)

	fmt.Printf("- Solution: %v\n\n", solution)
}
