package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"

	"github.com/cbarrick/evo"
	"github.com/cbarrick/evo/perm"
	"github.com/cbarrick/evo/pop/graph"
	"github.com/cbarrick/evo/sel"
)

// Tuneables
const (
	dim = 4 // dimension of the problem
)

// Global objects
var (
	// Counts the number of fitness evaluations
	count struct {
		sync.Mutex
		n int
	}
)

// The magic type specifies our genome. We use a permutation of [0,dim^2) to
// represent the square. Note that Wikipedia specifies magic squares as
// permutations of (0,dim^2]. This difference may have subtle implications on
// the theory.
type magic struct {
	gene    []int     // the square, interpreted as the concatination of the rows
	fitness float64   // the number of rows, cols, & diags without conflicts
	once    sync.Once // used to compute the fitness lazily
}

// String returns a representation of the square. The output is a valid Prolog
// term, and the numeric values are incremented by 1 to match the theory on
// Wikipedia. Thus the output can be verified directly by the Prolog solver.
func (m magic) String() (s string) {
	n := 0
	s += "["
	for i := range m.gene {
		if i != 0 {
			s += ","
		}
		if n == 0 {
			s += "["
		}
		s += strconv.Itoa(m.gene[i] + 1)
		n = (n + 1) % dim
		if n == 0 {
			s += "]"
		}
	}
	s += "]"
	return s
}

// Close does nothing because we use no closable resources.
func (m *magic) Close() {}

// Fitness returns the number of rows, columns, and diagonal that do not contain
// conflicts. The fitness is only computed once across all calls to the method
// by using a sync.Once.
func (m *magic) Fitness() float64 {
	m.once.Do(func() {
		sum := dim * (dim * dim - 1) / 2
		var part int // partial sum, used to compute sums of rows, cols, diags.

		// rows
		for i := 0; i < dim; i++ {
			part = 0
			for j := 0; j < dim; j++ {
				part += m.gene[dim*i+j]
			}
			if part == sum {
				m.fitness++
			}
		}

		// columns
		for i := 0; i < dim; i++ {
			part = 0
			for j := 0; j < dim; j++ {
				part += m.gene[dim*j+i]
			}
			if part == sum {
				m.fitness++
			}
		}

		// diagonals
		part = 0
		for i := 0; i < dim; i++ {
			part += m.gene[dim*i+i]
		}
		if part == sum {
			m.fitness++
		}
		part = 0
		for i := 0; i < dim; i++ {
			part += m.gene[dim*(i+1)-i-1]
		}
		if part == sum {
			m.fitness++
		}

		// increment the counter of fitness evals
		count.Lock()
		count.n++
		count.Unlock()
	})
	return m.fitness
}

// Evolve specifies the inner loop of our evolutionary algorithm.
func (m *magic) Evolve(suitors ...evo.Genome) evo.Genome {
	// Creation:
	// We allocate a new genome for the child on the heap
	child := &magic{gene: make([]int, dim*dim)}

	// Crossover:
	// We're using a diffusion model. Thus we pick one suitor to cross with the
	// current genome in this spot. We pick the mate using a single random
	// binary tournament. We use cycle crossover to inherit positioning.
	mate := sel.BinaryTournament(suitors...).(*magic)
	perm.CycleX(child.gene, m.gene, mate.gene)

	// Mutation:
	// 20% chance of a single random swap, 4% chance of two swaps, 0.8% chance
	// of three swaps, etc.
	for {
		n := rand.Float64()
		if n < 0.2 {
			perm.RandSwap(child.gene)
		} else {
			break
		}
	}

	// Replacement:
	// Only replace if the child is better or equal.
	if child.Fitness() < m.Fitness() {
		return m
	}
	return child
}

func main() {
	fmt.Println("Dimension:", dim)

	// Setup:
	// We create an initial random population to evolve in a diffusion model.
	init := make([]evo.Genome, 1024)
	for i := range init {
		init[i] = &magic{gene: rand.Perm(dim*dim)}
	}
	pop := graph.Hypercube(init)
	pop.Start()

	// Tear-down:
	// Cleanup and print a solution if we have one.
	defer func() {
		pop.Close()
		max := evo.Max(pop)
		fmt.Println()
		if max.Fitness() != dim*2+2 {
			fmt.Println("Solution Not Found")
			fmt.Println("Best:", max)
		} else {
			fmt.Println("Solution:", max)
		}
	}()

	// Run:
	// Collect and print stats and evaluate the termination conditions.
	for {
		count.Lock()
		n := count.n
		count.Unlock()
		stats := pop.Stats()

		// "\x1b[2K" is the escape code to clear the line
		fmt.Printf("\x1b[2K\rCount: %7d | Max: %2.0f | Min: %2.0f | Mean: %3.1f | SD: %6.6g",
			n,
			stats.Max(),
			stats.Min(),
			stats.Mean(),
			stats.StdDeviation())

		// We've found the solution when max is dim*2+2
		if stats.Max() == dim*2+2 {
			return
		}

		// We've converged once the deviation is less than 0.01
		if stats.StdDeviation() < 1e-2 {
			return
		}

		// Force stop after 2,000,000 fitness evaluations
		if n > 2e6 {
			return
		}
	}
}
