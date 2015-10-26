package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/cbarrick/evo"
	"github.com/cbarrick/evo/perm"
	"github.com/cbarrick/evo/pop/graph"
	"github.com/cbarrick/evo/sel"
)

type square struct {
	genes   []int
	fitness float64   // cache of the fitness
	once    sync.Once // used to sync fitness computations
}

func (s *square) Close() {}

func (s *square) Fitness() float64 {
	s.once.Do(func() {
		// compuation of Fitness

		// s.fitness = ?
	})
	return s.fitness
}

func (s *square) Evolve(matingPool ...evo.Genome) evo.Genome {

}

//go fmt <>// it would format file
