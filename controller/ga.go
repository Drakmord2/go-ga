package controller

import (
	"fmt"

	"github.com/Drakmord2/go-ga/model"
	"github.com/Drakmord2/go-ga/util"
)

// GeneticAlgorithm finds the best solution to the problem
func GeneticAlgorithm(config util.Config) {
	fmt.Printf("\nConfiguration: %v\n\n", config)

	Population := initialPopulation(config.Population)

	for i := 0; i < config.MaxIteration; i++ {
		heuristic(Population)

		crossover()
		mutation()
		Population = selection()
	}

}

func heuristic(population model.Population) int {
	return len(population.Chromosomes)
}

func initialPopulation(populationSize int) model.Population {
	return model.Population{}
}
