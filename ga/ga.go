package ga

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/cbarrick/evo"
	"github.com/cbarrick/evo/perm"
	"github.com/cbarrick/evo/pop/gen"
	"github.com/cbarrick/evo/sel"
)

// An instance represents a problem instance being solved. This type carries the
// global data used for the optimization. The schema is the partially filled
// square that we're trying to complete, where values greater than or equal to
// 0 appearing in the schema must appear in all genes at the same position, and
// values less than 0 in the schema represent variable values in the genes.
type instance struct {
	sync.Mutex
	order  int   // order of the square being solved
	schema []int // schema of the square being solved, -1 for variables
	count  int   // number of fitness evaluations
}

// The Magic type specifies our genome. We use a permutation of [0,order^2) to
// represent the square. Note that Wikipedia specifies magic squares as
// permutations of (0,order^2]. This difference may have subtle implications on
// the theory.
type Magic struct {
	gene []int // the square, interpreted as the concatination of the rows
	inst *instance
}

// String returns a representation of the square. The output is a valid Prolog
// term, and the numeric values are incremented by 1 to match the theory on
// Wikipedia. Thus the output can be verified directly by the Prolog solver.
func (m *Magic) String() (s string) {
	order := m.inst.order
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
		n = (n + 1) % order
		if n == 0 {
			s += "]"
		}
	}
	s += "]"
	s += fmt.Sprintf("@%v", m.Fitness())
	return s
}

// Close does nothing because we use no closable resources.
func (m *Magic) Close() {}

// Fitness returns the number of rows, columns, and diagonal that do not contain
// conflicts. The fitness is only computed once across all calls to the method
// by using a sync.Once.
func (m *Magic) Fitness() (fitness float64) {
	const (
		penalty = 0.25
	)

	var (
		order = m.inst.order

		// the target that each row, col, diag should sum to
		target = order * (order*order - 1) / 2

		// the actual value of the row, col, diag
		actual = 0
	)

	// rows
	for i := 0; i < order; i++ {
		actual = 0
		for j := 0; j < order; j++ {
			actual += m.gene[order*i+j]
		}
		if actual == target {
			fitness += float64(target)
		} else {
			fitness -= penalty * math.Abs(float64((actual-target)))
		}
	}

	// columns
	for i := 0; i < order; i++ {
		actual = 0
		for j := 0; j < order; j++ {
			actual += m.gene[order*j+i]
		}
		if actual == target {
			fitness += float64(target)
		} else {
			fitness -= penalty * math.Abs(float64((actual-target)))
		}
	}

	// diagonals
	actual = 0
	for i := 0; i < order; i++ {
		actual += m.gene[order*i+i]
	}
	if actual == target {
		fitness += float64(target)
	} else {
		fitness -= penalty * math.Abs(float64((actual-target)))
	}
	actual = 0
	for i := 0; i < order; i++ {
		actual += m.gene[order*(i+1)-i-1]
	}
	if actual == target {
		fitness += float64(target)
	} else {
		fitness -= penalty * math.Abs(float64((actual-target)))
	}

	// increment the counter of fitness evals
	m.inst.Lock()
	m.inst.count++
	m.inst.Unlock()

	return fitness
}

// Evolve specifies the inner loop of our evolutionary algorithm.
func (m *Magic) Evolve(suitors ...evo.Genome) evo.Genome {
	order := m.inst.order
	size := order * order

	// Creation:
	// We allocate a new genome for the child on the heap
	child := &Magic{gene: make([]int, order*order), inst: m.inst}

	// Crossover:
	// We're using a diffusion model. Thus we pick one suitor to cross with the
	// current genome in this spot. We pick the mate using a single random
	// binary tournament. We use cycle crossover to inherit positioning.
	mate := sel.BinaryTournament(suitors...).(*Magic)
	perm.CycleX(child.gene, m.gene, mate.gene)

	// Mutation:
	// 20% chance of a single random swap, 4% chance of two swaps, 0.8% chance
	// of three swaps, etc.
	for {
		chance := rand.Float64()
		if chance < 0.2 {
			var i, j int
			p := perm.New(size)
			for i = 0; i < size; i++ {
				if m.inst.schema[p[i]] < 0 {
					break
				}
			}
			for j = i + 1; j < size; j++ {
				if m.inst.schema[p[j]] < 0 {
					break
				}
			}
			if i != size && j != size {
				child.gene[p[i]], child.gene[p[j]] = child.gene[p[j]], child.gene[p[i]]
			}
		} else if chance < 0.4 {
			smartMutation(child.gene, child.inst.schema)
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

func smartMutation(gene, schema []int) bool {
	const (
		VERT  = 0
		HORTZ = 1
	)

	var (
		chosenIdx   int
		findIdx     int
		targetValue int
		diff        int
		dir         = rand.Intn(2)
		dim         = sqrt(len(gene))
		cTL         = 0
		cTR         = dim - 1
		cBL         = (dim*dim - dim)
		cBR         = (dim*dim - 1)
	)

	switch rand.Intn(4) {
	case 0:
		chosenIdx = cTL
	case 1:
		chosenIdx = cTR
	case 2:
		chosenIdx = cBL
	case 3:
		chosenIdx = cBR
	}

	if dir == HORTZ {
		if chosenIdx == cTL || chosenIdx == cTR {
			diff = gene[cBL] - gene[cBR]
			if chosenIdx == cTL {
				targetValue = gene[cTR] + diff
			} else {
				targetValue = gene[cTL] - diff
			}
		} else {
			diff = gene[cTL] - gene[cTR]
			if chosenIdx == cBL {
				targetValue = gene[cBR] + diff
			} else {
				targetValue = gene[cBL] - diff
			}
		}
	} else {
		if chosenIdx == cTL || chosenIdx == cBL {
			diff = gene[cTR] - gene[cBR]
			if chosenIdx == cTL {
				targetValue = gene[cBL] + diff
			} else {
				targetValue = gene[cTL] - diff
			}
		} else {
			diff = gene[cTL] - gene[cBL]
			if chosenIdx == cTR {
				targetValue = gene[cBR] + diff
			} else {
				targetValue = gene[cTR] - diff
			}
		}
	}

	if targetValue < 0 {
		targetValue += dim * dim
	} else if targetValue >= dim*dim {
		targetValue -= dim * dim
	}

	findIdx = perm.Search(gene, targetValue)
	if schema[findIdx] < 0 && schema[chosenIdx] < 0 {
		gene[chosenIdx], gene[findIdx] = gene[findIdx], gene[chosenIdx]
		return true
	}
	return false
}

// sqrt returns the integer part of the square root of n
func sqrt(n int) int {
	n64 := int64(n)
	res := int64(0)
	one := int64(1) << 62
	for n64 < one {
		one >>= 2
	}
	for one != 0 {
		if res+one <= n64 {
			n64 -= res + one
			res += one << 1
		}
		res >>= 1
		one >>= 2
	}
	return int(res)
}

// seed creates a random set of n Magic genomes to be used as an initial
// population for the given problem instance
func seed(inst *instance, n int) []evo.Genome {
	init := make([]evo.Genome, n)
	size := len(inst.schema)
	for i := range init {
		init[i] = &Magic{
			gene: perm.New(size),
			inst: inst,
		}
		for j := range inst.schema {
			if inst.schema[j] < 0 {
				continue
			}
			gene := init[i].(*Magic).gene
			k := perm.Search(gene, inst.schema[j])
			gene[j], gene[k] = gene[k], gene[j]
		}
	}
	return init
}

// Solve returns a Magic representing the solution to a magic square schema. The
// schema is a partially filled magic square, interpreted as the concatination
// of the rows of the square. Values greater than or equal to 0 appearing in the
// schema must appear in the solution at the same position, and values less than
// 0 in the schema represent unknown values.
//
// The population is reseeded occasionally if the solution is not found. The
// return value rsc is the "reseed count" giving the number of times reseeding
// occurs.
func Solve(schema []int) (soln *Magic, rsc int) {
	var (
		// the interface into the GA loop
		pop evo.Population

		// the number of values in the square
		size = len(schema)

		// the order of the problem
		order = sqrt(size)

		// magic constant for this order
		target = order * (order*order - 1) / 2

		// the size of the population
		popz = 10 * order

		// the problem instance bein solved
		inst = &instance{
			order:  order,
			schema: schema,
		}

		// variables used to reseed the population
		reseedT = 4 * time.Millisecond
		reseedC = time.After(reseedT)
		reseed  = func() {
			rsc++
			pop.Close()
			s := seed(inst, popz)
			s[popz-1] = evo.Max(pop)
			pop = gen.New(s)
			pop.Start()
			reseedC = time.After(reseedT)
		}

		timeout = time.After(30 * time.Second)
	)

	// start the GA in the background
	pop = gen.New(seed(inst, popz))
	pop.Start()

	for {
		select {
		case <-timeout:
			pop.Close()
			for i := pop.Iter(); i.Value() != nil; i.Next() {
				fmt.Println(i.Value())
			}
			fmt.Println(pop.Stats().RSD())
			return soln, rsc

		// reseed the population periodically with exponential backoff
		case <-reseedC:
			reseed()
			reseedT *= 2

		// stop when we find a solution
		// or reseed if we converge prematurely
		default:
			s := pop.Stats()
			if s.Max() == float64(target*(order*2+2)) {
				pop.Close()
				soln = evo.Max(pop).(*Magic)
				return soln, rsc
			}
			if s.RSD() < 0.3 {
				reseed()
			}
		}
	}
}

func Generate(ord int) (soln *Magic, rsc int) {
	schema := make([]int, ord*ord)
	for i := range schema {
		schema[i] = -1
	}
	return Solve(schema)
}
