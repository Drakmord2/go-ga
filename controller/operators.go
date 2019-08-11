package controller

import (
	"math/rand"
	"sort"
	"time"

	"github.com/Drakmord2/go-ga/model"
)

// Tournament Selection
func parentSelection(population *[]model.Chromosome) []int {
	parents := make([]int, 2)

	for i := 0; i < 2; i++ {
		contestants := make([]byte, random(len(*population)/2, len(*population)-1))
		rand.Read(contestants)

		best := 0
		prev := 1000.
		for j := 0; j < len(contestants); j++ {
			if (*population)[j].GetFitness() < prev {
				best = j
				(*population)[j].SetFitness(1000.)
			}
		}

		parents[i] = best
	}

	return parents
}

// Single-point crossover
func crossover(population *[]model.Chromosome, parents []int, pc float32) []model.Chromosome {
	rand.Seed(time.Now().UnixNano())
	geneSize := len((*population)[parents[0]].GetGenes())

	genesp1 := make([]model.Gene, geneSize)
	genesp2 := make([]model.Gene, geneSize)
	copy(genesp1, (*population)[parents[0]].GetGenes())
	copy(genesp2, (*population)[parents[1]].GetGenes())

	pivot := random(1, len(genesp1)-1)

	var cross1 []model.Gene
	var cross2 []model.Gene
	for i := pivot; i < len(genesp1); i++ {
		cross1 = append(cross1, genesp2[i])
		cross2 = append(cross2, genesp1[i])
	}

	j := 0
	for i := pivot; i < len(genesp1); i++ {
		genesp1[i] = cross1[j]
		genesp2[i] = cross2[j]
		j++
	}

	var (
		child1 model.Chromosome
		child2 model.Chromosome
	)
	child1.SetGenes(genesp1)
	child2.SetGenes(genesp2)

	offspring := []model.Chromosome{child1, child2}

	return offspring
}

// Mutation (bit-flip) of genes
func mutation(offspring *[]model.Chromosome, pm float32) {
	rand.Seed(time.Now().UnixNano())

	for i := range *offspring {
		genes := (*offspring)[i].GetGenes()
		for j := range genes {
			if rand.Float32() <= pm {
				if genes[j].GetAllele() == 0 {
					genes[j].SetAllele(1)
					continue
				}
				genes[j].SetAllele(0)
			}
		}
	}

}

func survivorSelection(population *[]model.Chromosome) {
	var index int
	populationSize := len(*population)
	scores := make([]float64, populationSize)

	for i := range *population {
		scores[i] = (*population)[i].GetFitness()
	}
	sort.Float64s(scores)

	for j := 0; j < 2; j++ {
		for i := range *population {
			fitness := (*population)[i].GetFitness()
			notfit := scores[populationSize-1]

			if j == 1 {
				notfit = scores[populationSize-2]
			}

			if fitness == notfit {
				index = i
				break
			}
		}
		*population = removeIndex(*population, index)
	}
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func removeIndex(s []model.Chromosome, i int) []model.Chromosome {
	s[i] = s[0]
	return s[1:]
}
