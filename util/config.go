package util

// Config stores the parameters of the Genetic Algorithm
type Config struct {
	MaxIteration  int
	FitnessGoal   float64
	Objective     float64
	Population    int
	Parameters    []string
	MutationRate  float32
	CrossoverRate float32
	Parallel      bool
}
