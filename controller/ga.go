package controller

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/Drakmord2/go-ga/model"
	"github.com/Drakmord2/go-ga/util"
)

// Verbose output for debugging
var Verbose = false

// GeneticAlgorithm finds the best solution to the problem
func GeneticAlgorithm(config util.Config) (model.Chromosome, int) {
	solution := model.Chromosome{}
	solution.SetFitness(1000.)

	// Initialization
	Population := initialPopulation(config.Population, config.Parameters)
	heuristic(&Population)

	var iterations int
	for i := 0; i < config.MaxIteration; i++ {
		log("\nIteration: %d\n", i+1)

		log("  - Parent Selection")
		Parents := parentSelection(&Population)

		log("  - Crossover")
		Offspring := crossover(&Population, Parents, config.CrossoverRate)

		log("  - Mutation")
		mutation(&Offspring, config.MutationRate)

		log("  - New Population")
		Population = append(Population, Offspring...)

		log("  - Evaluation")
		heuristic(&Population)

		log("  - Survivor Selection")
		survivorSelection(&Population)

		// Fittest solution
		solution = pickBest(&Population, solution)
		bestfit := solution.GetFitness()
		fmt.Printf("Fitness: %f", bestfit)

		if bestfit == config.FitnessGoal {
			iterations = i + 1
			break
		}
	}

	if iterations == 0 {
		iterations = config.MaxIteration
	}

	return solution, iterations
}

// Picks best solution
func pickBest(population *[]model.Chromosome, solution model.Chromosome) model.Chromosome {
	localBest := model.Chromosome{}
	localBest.SetFitness(1000.)
	for i := range *population {
		if (*population)[i].GetFitness() < localBest.GetFitness() {
			localBest = (*population)[i]
		}
	}

	if localBest.GetFitness() < solution.GetFitness() {
		solution = localBest
	}

	return solution
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
			(*population)[i].SetFitness(1000.)
			continue
		}

		fitness := math.Abs(180. - sum)
		(*population)[i].SetFitness(fitness)
	}
}

func initialPopulation(populationSize int, parameters []string) []model.Chromosome {
	rand.Seed(time.Now().UnixNano())
	population := make([]model.Chromosome, populationSize, populationSize+4)

	for j := 0; j < populationSize; j++ {
		genes := make([]model.Gene, len(parameters))
		for i := range parameters {
			var gene model.Gene
			gene.SetParameter(parameters[i])
			gene.SetAllele(rand.Intn(2))

			genes[i] = gene
		}

		var chromosome model.Chromosome
		chromosome.SetGenes(genes)

		population[j] = chromosome
	}

	return population
}

func log(format string, text ...interface{}) {
	if Verbose {
		if text != nil {
			fmt.Printf(format, text...)
			return
		}
		fmt.Println(format)
	}
}
