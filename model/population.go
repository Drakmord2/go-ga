package model

// Gene defines a parameter to be found
type Gene struct {
	parameter string
	allele    int
}

// Chromosome is a solution to the problem
type Chromosome struct {
	genes   []Gene
	fitness float64
}

// GetGenes returns the genes of the chromosome
func (c Chromosome) GetGenes() []Gene {
	return c.genes
}

// SetGenes inserts the genes of the chromosome
func (c *Chromosome) SetGenes(genes []Gene) {
	c.genes = genes
}

// GetFitness returns the fitness of the chromosome
func (c Chromosome) GetFitness() float64 {
	return c.fitness
}

// SetFitness inserts the fitness of the chromosome
func (c *Chromosome) SetFitness(fitness float64) {
	c.fitness = fitness
}

// SetParameter sets the parameter of the gene
func (g *Gene) SetParameter(parameter string) {
	g.parameter = parameter
}

// SetAllele sets the allele of the gene
func (g *Gene) SetAllele(allele int) {
	g.allele = allele
}

// GetParameter returns the parameter of the gene
func (g Gene) GetParameter() string {
	return g.parameter
}

// GetAllele returns the allele of the gene
func (g Gene) GetAllele() int {
	return g.allele
}
