package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"

	"github.com/cbarrick/evo"
	"github.com/cbarrick/evo/perm"
	"github.com/cbarrick/evo/pop/graph"
)

const (
	dim = 4
)

var (
	count struct {
		sync.Mutex
		n int
	}
)

type magic struct {
	gene    []int
	fitness float64
	once    sync.Once
}

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

func (m *magic) Close() {}

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

func (m *magic) Evolve(suitors ...evo.Genome) evo.Genome {
	// Creation:
	child := &magic{gene: make([]int, dim*dim)}

	// Crossover:
	mom := suitors[rand.Intn(len(suitors))].(*magic)
	dad := suitors[rand.Intn(len(suitors))].(*magic)
	perm.OrderX(child.gene, mom.gene, dad.gene)

	// Mutation:
	switch n := rand.Float64(); {
	case n < 0.1:
		perm.RandInvert(child.gene)
	case n < 0.2:
		perm.RandSwap(child.gene)
	}

	// Replacement:
	if child.Fitness() < m.Fitness() {
		return m
	}
	return child
}

func main() {
	fmt.Println("Dimension:", dim)

	// Setup:
	init := make([]evo.Genome, 1024)
	for i := range init {
		init[i] = &magic{gene: rand.Perm(dim*dim)}
	}
	pop := graph.Grid(init)
	pop.Start()

	// Tear-down:
	defer func() {
		pop.Close()
		max := evo.Max(pop)
		if max.Fitness() != dim*2+2 {
			fmt.Println("\nSolution Not Found")
		} else {
			fmt.Println("\nSolution:", evo.Max(pop))
		}
	}()

	// Run:
	for {
		count.Lock()
		n := count.n
		count.Unlock()
		stats := pop.Stats()

		// "\x1b[2K" is the escape code to clear the line
		fmt.Printf("\x1b[2K\rCount: %7d | Max: %4.0f | Min: %4.0f | SD: %6.6g",
			n,
			stats.Max(),
			stats.Min(),
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
