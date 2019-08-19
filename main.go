package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/Drakmord2/go-ga/controller"
	"github.com/Drakmord2/go-ga/util"
)

func main() {
	virtualThreads := 32
	osthreads := runtime.GOMAXPROCS(virtualThreads)

	config := util.Config{
		MaxIteration: 100000,
		Population:   100,
		Parameters: []string{
			"b15", "b14", "b13", "b12", "b11",
			"b10", "b9", "b8", "b7", "b6",
			"b5", "b4", "b3", "b2", "b1", "b0",
		},
		FitnessGoal:   0.,
		Objective:     52428,
		CrossoverRate: 0.7,
		MutationRate:  0.07,
		Parallel:      true,
	}

	fmt.Println("\n\t- Go-GA -")
	if config.Parallel {
		fmt.Printf("\nConfiguration:   %v\nOS Threads:      %d\nVirtual Threads: %d\n\n", config, osthreads, virtualThreads)
	} else {
		fmt.Printf("\nConfiguration: %v\n\n", config)
	}

	start := time.Now()
	solution, iterations := controller.GeneticAlgorithm(config)

	fmt.Printf("\n\t- Results -\n\n%s\nIterations: %d\nDuration:   %v\n\n", solution, iterations, time.Since(start))
}
