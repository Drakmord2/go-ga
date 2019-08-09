package util

// Config stores the parameters of the Genetic Algorithm
type Config struct {
	MaxIteration  int
	Population    int
	Parameters    []string
	MutationRate  float32
	CrossoverRate float32
}