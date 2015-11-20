package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type square []int

func (sq square) Dim() int {
	return isqrt(len(sq))
}

// String returns the square as a Prolog term.
// 0s are replaced by placeholders.
func (sq square) String() string {
	var buffer bytes.Buffer
	_, _ = buffer.WriteString("[")
	dim := sq.Dim()
	for i, x := range sq {
		if i > 0 {
			_, _ = buffer.WriteString(", ")
		}
		if i%dim == 0 {
			_, _ = buffer.WriteString("[")
		}
		if x != 0 {
			_, _ = buffer.WriteString(strconv.Itoa(x))
		} else {
			_, _ = buffer.WriteString("_")
		}
		if i%dim == dim-1 {
			_, _ = buffer.WriteString("]")
		}
	}
	_, _ = buffer.WriteString("]")
	return buffer.String()
}

// isqrt returns the integer square root of x
func isqrt(x int) int {
	var root int64
	x64 := int64(x)
	bit := int64(1) << 62
	for bit > x64 {
		bit >>= 2
	}
	for bit != 0 {
		if x64 >= root+bit {
			x64 -= root + bit
			root >>= 1
			root += bit
		} else {
			root >>= 1
		}
		bit >>= 2
	}
	return int(root)
}

// move takes an index, i, into a square of some dimension and returns a new
// index that is offset from i along the x and y axes. The movement wraps
// around edges.
func move(i, dim, dx, dy int) int {
	for dx < 0 {
		dx += dim
	}
	for dim <= dx {
		dx -= dim
	}
	for dy < 0 {
		dy += dim
	}
	for dim <= dy {
		dy -= dim
	}
	x := (i + dx) % dim
	y := (i/dim + dy) % dim
	return x + dim*y
}

// odd1 implements the Siamese method (up & right / down)
func odd1(dim int) (sq square) {
	sq = make(square, dim*dim)
	i := dim / 2
	for n := 1; n <= dim*dim; n++ {
		sq[i] = n
		j := move(i, dim, +1, -1)
		if sq[j] != 0 {
			j = move(i, dim, 0, +1)
		}
		i = j
	}
	return sq
}

// odd2 implements a variation the Siamese method (up & right / up & up)
func odd2(dim int) (sq square) {
	sq = make(square, dim*dim)
	i := dim / 2
	for n := 1; n <= dim*dim; n++ {
		sq[i] = n
		j := move(i, dim, +1, -1)
		if sq[j] != 0 {
			j = move(i, dim, 0, -2)
		}
		i = j
	}
	return sq
}

func main() {
	fmt.Println(odd1(5))
	fmt.Println(odd2(5))
}
