package controller

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/Drakmord2/go-ga/model"
	"github.com/Drakmord2/go-ga/util"
)

// GA()
//    initialize population
//    find fitness of population

//    while (termination criteria is reached) do
//       parent selection
//       crossover with probability pc
//       mutation with probability pm
//       decode and fitness calculation
//       survivor selection
//       find best
//    return best

// GeneticAlgorithm finds the best solution to the problem
func GeneticAlgorithm(config util.Config) model.Chromosome {
	fmt.Printf("\nConfiguration: %v\n\n", config)

	var solution model.Chromosome
	Population := initialPopulation(config.Population, config.Parameters)
	heuristic(&Population)
	Population = selection(Population)

	for i := 0; i < config.MaxIteration; i++ {
		fmt.Printf("Iteration: %d\n\n", i)

		crossover()
		mutation()
		heuristic(&Population)
		Population = selection(Population)

		solution = pickBest(Population)
	}

	return solution
}

// Picks best solution
func pickBest(population []model.Chromosome) model.Chromosome {
	best := model.Chromosome{}
	best.SetFitness(1000.)
	for i := range population {
		if population[i].GetFitness() < best.GetFitness() {
			best = population[i]
		}
	}

	return best
}

// Finds the value 180
func heuristic(population *[]model.Chromosome) {
	for i := range *population {
		var sum float64
		genes := (*population)[i].GetGenes()
		for j := range genes {
			allele := genes[j].GetAllele()
			sum += math.Pow(2, float64(7-j)) * float64(allele)
		}
		if sum == 0 {
			(*population)[i].SetFitness(0)
			continue
		}

		fitness := math.Abs(180. - sum)
		(*population)[i].SetFitness(fitness)
	}
}

func initialPopulation(populationSize int, parameters []string) []model.Chromosome {
	rand.Seed(time.Now().UnixNano())
	population := make([]model.Chromosome, populationSize)

	for j := 0; j < populationSize; j++ {
		genes := make([]model.Gene, len(parameters))
		for i := range parameters {
			var gene model.Gene
			gene.SetParameter(parameters[i])
			gene.SetAllele(rand.Intn(2))

			genes[i] = gene
		}

		var chromosome model.Chromosome
		chromosome.SetID(j)
		chromosome.SetGenes(genes)

		population[j] = chromosome
	}

	return population
}
