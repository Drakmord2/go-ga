package main

import (
	"fmt"

	"github.com/Drakmord2/go-ga/controller"
	"github.com/Drakmord2/go-ga/util"
)

func main() {
	config := util.Config{
		MaxIteration:  100,
		FitnessGoal:   0.,
		Population:    10,
		Parameters:    []string{"b7", "b6", "b5", "b4", "b3", "b2", "b1", "b0"},
		CrossoverRate: 0.7,
		MutationRate:  0.01,
	}

	fmt.Println("\n\t- Go-GA -")
	fmt.Printf("\nConfiguration: %v\n", config)

	solution, iterations := controller.GeneticAlgorithm(config)

	fmt.Printf("\n\n\t- Results -\n\nSolution:   %v\nFitness:    %f\nIterations: %d\n\n", solution.GetGenes(), solution.GetFitness(), iterations)
}
