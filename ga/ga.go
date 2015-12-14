package ga

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	"github.com/cbarrick/evo"
	"github.com/cbarrick/evo/pop/gen"
	"github.com/cbarrick/evo/sel"
)

// Genome is the implementation of the evo.Genome interface for the GA.
// It embeds the Siam type which embeds the Square type.
// (Embeding is Go's version of inheritance)
type Genome struct {
	Siam
}

func (*Genome) Close() {}

func (g *Genome) Fitness() float64 {
	g.Express()
	return float64(len(g.GoodSet()))
}

func (g *Genome) Evolve(suitors ...evo.Genome) evo.Genome {
	// Creation:
	child := &Genome{}

	// Crossover:
	mom := sel.BinaryTournament(suitors...).(*Genome)
	dad := sel.BinaryTournament(suitors...).(*Genome)
	child.UniformX(&mom.Siam, &dad.Siam)

	// Mutation:
	for chance := rand.Float64(); chance < 0.2; chance = rand.Float64() {
		child.Mutate()
	}

	if child.Fitness() == float64(child.order*child.order+2) {
		fmt.Println(child)
	}

	// Replacement:
	return child
}

// Solve returns a Magic representing the solution to a magic square schema. The
// schema is a partially filled magic square, interpreted as the concatination
// of the rows of the square. Values greater than or equal to 0 appearing in the
// schema must appear in the solution at the same position, and values less than
// 0 in the schema represent unknown values.
//
// The value is returned on a channel allocated by the caller.
//
// The population is reseeded occasionally if the solution is not found. The
// return value rsc is the "reseed count" giving the number of times reseeding
// occurs.
func Solve(schema Square, ret chan Square) {
	var (
		// the number of values in the square
		order = len(schema.Sq)

		// the size of the population
		popz = 10 * order

		// the interface into the GA loop
		pop evo.Population

		// the initial population
		seed = make([]evo.Genome, popz)
	)

	// Print the population when receiving a SIGQUIT.
	// (press ctrl-/ at the terminal)
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGQUIT)
		<-sigs
		stats := pop.Stats()
		fmt.Println(stats.SD(), stats.RSD())
		for i := pop.Iter(); i.Value() != nil; i.Next() {
			fmt.Println(i.Value().(*Genome).Express())
		}
		os.Exit(2)
	}()

	// start the GA in the background
	for i := range seed {
		seed[i] = &Genome{RandSiam(schema)}
	}
	pop = gen.New(seed)
	pop.Start()

	for {
		best := evo.Max(pop)
		stats := pop.Stats()

		// halt when we find a solution
		if best.Fitness() == float64(order*2+2) {
			ret <- best.(*Genome).Square
			pop.Close()
			return
		}

		// reseed
		if stats.RSD() < 0.3 {
			pop.Close()
			for i := range seed {
				seed[i] = &Genome{RandSiam(schema)}
			}
			seed[0] = best
			pop = gen.New(seed)
			pop.Start()
		}
	}
}
