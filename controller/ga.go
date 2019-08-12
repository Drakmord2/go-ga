package controller

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/Drakmord2/go-ga/model"
	"github.com/Drakmord2/go-ga/util"
)

// Verbose output for debugging
var Verbose = false

// GeneticAlgorithm finds the best solution to the problem
func GeneticAlgorithm(config util.Config) (string, int) {
	var iterations int
	solution := model.Chromosome{}
	solution.SetFitness(9999.)

	// Initialization
	Population := initialPopulation(config.Population, config.Parameters)
	heuristic(&Population, config.Objective)

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
		heuristic(&Population, config.Objective)

		log("  - Survivor Selection")
		survivorSelection(&Population)

		// Fittest solution
		solution = pickBest(&Population, solution)
		bestfit := solution.GetFitness()
		fmt.Printf("Fitness: %f\n", bestfit)

		if bestfit == config.FitnessGoal {
			iterations = i + 1
			break
		}
	}

	if iterations == 0 {
		iterations = config.MaxIteration
	}

	return solutionString(solution), iterations
}

// Picks best solution
func pickBest(population *[]model.Chromosome, solution model.Chromosome) model.Chromosome {
	localBest := model.Chromosome{}
	localBest.SetFitness(9999.)
	for i := range *population {
		if (*population)[i].GetFitness() <= localBest.GetFitness() {
			localBest = (*population)[i]
		}
	}

	if localBest.GetFitness() <= solution.GetFitness() {
		solution = localBest
	}

	return solution
}

// Fittest individuals are close to the objective value
func heuristic(population *[]model.Chromosome, objectiveValue float64) {
	for i := range *population {
		genes := (*population)[i].GetGenes()

		sum := 0.
		for j := range genes {
			allele := genes[j].GetAllele()
			sum += math.Pow(2, float64(15-j)) * float64(allele)
		}

		fitness := math.Abs(objectiveValue - sum)
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

func solutionString(solution model.Chromosome) string {
	genes := solution.GetGenes()
	geneString := ""
	for i := 0; i < len(genes); i++ {
		geneString += strconv.Itoa(genes[i].GetAllele())
	}
	result := "Solution: {\n    Genes:   "
	result += geneString
	result += "\n    Fitness: "
	result += strconv.FormatFloat(solution.GetFitness(), 'f', 6, 64)
	result += "\n}"

	return result
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
