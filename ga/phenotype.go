package ga

import (
	"strconv"

	"github.com/cbarrick/evo/perm"
)

// Phenotype Representation
// --------------------------------------------------
// This file describes the phenotype representation of the GA. This
// representation can be used directly in the GA or a higher-level genotype
// representation may "express" this phenotype.

// A Square is a square matrix of ints used as the phenotype in the GA.
// When the square is being used as an individual in a population, it contains
// exactly one of all of the numbers from 1 to the size of the square. When used
// as a schema, the square may contain 0s to represent wildcards.
type Square struct {
	Arr []int   // the 1D interface to the square
	Sq  [][]int // the 2D interface to the square, shares backing array
}

// NewSquare returns a square of a given order filled with 0s.
func NewSquare(order int) (sq Square) {
	sq.Sq = make([][]int, order)
	sq.Arr = make([]int, order*order)
	for i := range sq.Sq {
		sq.Sq[i] = sq.Arr[i*order : i*order+order]
	}
	return sq
}

// RandSquare returns a randomly filled square of a given order.
func RandSquare(order int) (sq Square) {
	sq.Sq = make([][]int, order)
	sq.Arr = perm.New(order * order)
	for i := range sq.Arr {
		sq.Arr[i] = sq.Arr[i] + 1
	}
	for i := range sq.Sq {
		sq.Sq[i] = sq.Arr[i*order : i*order+order]
	}
	return sq
}

// NSquares genreates n squares matching a schema
func (sq Square) NSquares(n int, schema Square) (squares []Square) {
	squares = make([]Square, n)
	for i := range squares {
		squares[i] = RandSquare(sq.Order())
		squares[i].Repair(schema)
	}
	return squares
}

// Repair repares a square to match a schema.
func (sq Square) Repair(schema Square) {
	for i, a := range schema.Arr {
		if a != 0 && sq.Arr[i] != a {
			for j, b := range sq.Arr {
				if a == b {
					sq.Arr[i], sq.Arr[j] = sq.Arr[j], sq.Arr[i]
					break
				}
			}
		}
	}
}

// MagicConstant returns the magic constant for a square of this size.
// The magic constant is the value to which all rows/cols/diags of a magic square sum.
func (sq Square) MagicConstant() int {
	order := sq.Order()
	return order * (order*order + 1) / 2
}

// Order returns the length of a side of the square.
func (sq Square) Order() int {
	return len(sq.Sq)
}

// String returns a pretty representation of the square.
func (sq Square) String() (str string) {
	str += "["
	for i := range sq.Sq {
		if i != 0 {
			str += ","
		}
		str += "["
		for j := range sq.Sq[i] {
			if j != 0 {
				str += ","
			}
			if sq.Sq[i][j] == 0 {
				str += "_"
			} else {
				str += strconv.Itoa(sq.Sq[i][j])
			}
		}
		str += "]"
	}
	str += "]"
	return str
}

// Fitness
// --------------------------------------------------
// This section provides helpers that may be useful in computing fitness
// metrics for squares.

// GoodSet returns the set of all valid rows, columns, and diagonals.
// Lists in the set may share the backing array of the Square.
func (sq Square) GoodSet() (good [][]int) {
	order := sq.Order()
	mc := sq.MagicConstant()
	good = make([][]int, 0, order*order+2)

	// append rows and columns
	transpose := NewSquare(order)
	for i := range sq.Sq {
		for j := range sq.Sq[i] {
			transpose.Sq[j][i] = sq.Sq[i][j]
		}
	}
	for i := range sq.Sq {
		var rowSum, colSum int
		for j := range sq.Sq[i] {
			rowSum += sq.Sq[i][j]
			colSum += transpose.Sq[i][j]
		}
		if rowSum == mc {
			good = append(good, sq.Sq[i])
		}
		if colSum == mc {
			good = append(good, transpose.Sq[i])
		}
	}

	// append diags
	posDiag := make([]int, order)
	negDiag := make([]int, order)
	var posSum, negSum int
	for i := range sq.Sq {
		posDiag[i] = sq.Sq[i][i]
		negDiag[i] = sq.Sq[order-i-1][i]
		posSum += sq.Sq[i][i]
		negSum += sq.Sq[order-i-1][i]
	}
	if posSum == mc {
		good = append(good, posDiag)
	}
	if negSum == mc {
		good = append(good, negDiag)
	}

	return good
}

// FitDelta returns the degree by which the rows/cols/diags are different from
// their expected value (the magic constant).
func (sq Square) FitDelta() (delta int) {
	order := sq.Order()
	mc := sq.MagicConstant()

	// rows and columns
	transpose := NewSquare(order)
	for i := range sq.Sq {
		for j := range sq.Sq[i] {
			transpose.Sq[j][i] = sq.Sq[i][j]
		}
	}
	for i := range sq.Sq {
		var rowSum, colSum int
		for j := range sq.Sq[i] {
			rowSum += sq.Sq[i][j]
			colSum += transpose.Sq[j][i]
		}
		if rowSum < mc {
			delta += mc - rowSum
		} else if mc < rowSum {
			delta += rowSum - mc
		}
		if colSum < mc {
			delta += mc - colSum
		} else if mc < colSum {
			delta += colSum - mc
		}
	}

	// diags
	posDiag := make([]int, order)
	negDiag := make([]int, order)
	var posSum, negSum int
	for i := range sq.Sq {
		posDiag[i] = sq.Sq[i][i]
		negDiag[i] = transpose.Sq[i][i]
		posSum += sq.Sq[i][i]
		negSum += transpose.Sq[i][i]
	}
	if posSum < mc {
		delta += mc - posSum
	} else if mc < posSum {
		delta += posSum - mc
	}
	if negSum < mc {
		delta += mc - negSum
	} else if mc < negSum {
		delta += negSum - mc
	}

	return delta
}

// Helper Type: Cursors
// --------------------------------------------------

// A Cursor is a view into a square provides an API to move around squares.
type Cursor struct {
	Square
	Row int
	Col int
}

// Get returns the value under the cursor.
func (c *Cursor) Get() int {
	return c.Square.Sq[c.Row][c.Col]
}

// Set sets the value under the cursor.
func (c *Cursor) Set(n int) {
	c.Square.Sq[c.Row][c.Col] = n
}

// Move moves the cursor, wrapping around the edges.
func (c *Cursor) Move(dy, dx int) {
	order := c.Square.Order()
	c.Row += dy
	c.Col += dx
	c.Row %= order
	c.Col %= order
	for c.Row < 0 {
		c.Row += order
	}
	for c.Col < 0 {
		c.Col += order
	}
}
