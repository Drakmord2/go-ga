package controller

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Drakmord2/go-ga/model"
)

func parentSelection(population *[]model.Chromosome) []int {
	fmt.Println("  - Parent Selection")

	p1 := random(1, len(*population)-2)
	p2 := random(1, len(*population)-2)

	for p1 == p2 {
		p2 = random(1, len(*population)-2)
	}

	parents := []int{p1, p2}

	return parents
}

func crossover(population *[]model.Chromosome, parents []int, pc float32) []model.Chromosome {
	fmt.Println("  - Crossover")
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

func mutation(offspring *[]model.Chromosome, pm float32) {
	fmt.Println("  - Mutation")
	rand.Seed(time.Now().UnixNano())

	for i := range *offspring {
		genes := (*offspring)[i].GetGenes()
		for j := range genes {
			if rand.Float32() <= pm {
				fmt.Println("    - Mutation occured!")
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
	fmt.Println("  - Survivor Selection")
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
