package ga

import (
	"math/rand"
	"sync"

	"github.com/cbarrick/evo/integer"
)

// Genotype Representation
// --------------------------------------------------

// A Siam represents a specific instance of the general "Siamese Method" for
// producing magic squares. It is used as a genotype expressing Squares.
type Siam struct {
	Square             // the expressed Square
	once     sync.Once // used to lazily compute the expressed Square
	schema   Square    // the schema restricting the genotype
	order    int       // the order of the square
	start    [2]int    // the starting position
	ordvec   [2]int    // the ordinary movement
	breakvec [2]int    // the movement in break cases
}

func RandSiam(schema Square) (s Siam) {
	order := schema.Order()
	s = Siam{
		schema:   schema,
		order:    order,
		start:    [2]int{rand.Intn(order), rand.Intn(order)},
		ordvec:   [2]int{rand.Intn(order), rand.Intn(order)},
		breakvec: [2]int{rand.Intn(order), rand.Intn(order)},
	}
	s.start = [2]int{rand.Intn(order), rand.Intn(order)}
	return s
}

func NSiam(n int, schema Square) (siams []Siam) {
	siams = make([]Siam, n)
	for i := range siams {
		siams[i] = RandSiam(schema)
	}
	return siams
}

// Repair repares a square to match a schema.
func (s Siam) Repair() {
	s.Square.Repair(s.schema)
}

// Express returns the genotype expressed.
// This involves computing the Square using the general siamese method if it has
// not been previously computed.
func (s *Siam) Express() Square {
	s.once.Do(func() {
		s.Square = NewSquare(s.order)
		c := Cursor{s.Square, s.start[0], s.start[1]}
		dim := s.order * s.order
		var loopchk [][2]int
		for i := 1; i < dim; i++ {
			c.Set(i)
			c.Move(s.ordvec[0], s.ordvec[1])
			if c.Get() != 0 {
				c.Move(-s.ordvec[0], -s.ordvec[1])
			}
			for c.Get() != 0 {
				loopchk = append(loopchk, [2]int{c.Row, c.Col})
				c.Move(s.breakvec[0], s.breakvec[1])
				var loop bool
				for _, pos := range loopchk {
					if pos == [2]int{c.Row, c.Col} {
						loop = true
						break
					}
				}
				if loop {
					for h := 0; h < s.order; h++ {
						c.Row = h
						for k := 0; k < s.order; k++ {
							c.Col = k
							if c.Get() == 0 {
								break
							}
						}
						if c.Get() == 0 {
							break
						}
					}
				}
			}
		}
		c.Set(s.order * s.order)
		s.Square.Repair(s.schema)
	})
	return s.Square
}

// Crossover
// --------------------------------------------------
// This section provides helpers for implementing crossover in the GA.

func (s *Siam) UniformX(mom, dad *Siam) {
	s.schema = mom.schema
	s.order = mom.order
	integer.UniformX(s.start[:], mom.start[:], dad.start[:])
	integer.UniformX(s.ordvec[:], mom.ordvec[:], dad.ordvec[:])
	integer.UniformX(s.breakvec[:], mom.breakvec[:], dad.breakvec[:])
}

// Mutation
// --------------------------------------------------
// This section provides helpers for implementing mutation in the GA.

func (s *Siam) Mutate() {
	s.start[0] += int(rand.NormFloat64())
	s.start[1] += int(rand.NormFloat64())
	s.ordvec[0] += int(rand.NormFloat64())
	s.ordvec[1] += int(rand.NormFloat64())
	s.breakvec[0] += int(rand.NormFloat64())
	s.breakvec[1] += int(rand.NormFloat64())
	s.start[0] %= s.order
	s.start[1] %= s.order
	for s.start[0] < 0 {
		s.start[0] += s.order
	}
	for s.start[1] < 0 {
		s.start[1] += s.order
	}
}
