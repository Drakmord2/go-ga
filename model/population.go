package model

// Gene defines a parameter to be found
type Gene struct {
	parameter string
	allele    uint8
}

// Chromosome is a solution to the problem
type Chromosome struct {
	id    int
	genes []Gene
}

// Population is a set of solutions
type Population struct {
	Chromosomes []Chromosome
}
