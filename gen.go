package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
)

// square
// -------------------------

type square struct {
	m    [][]int // the square
	mask [][]bool
	d    density
	alg  string // the algorithm used to generate the square
}

type density float64

const (
	FULL density = 0
	HIGH density = 0.25
	MID  density = 0.5
	LOW  density = 0.75
)

func (sq *square) Init(n int) {
	sq.m = make([][]int, n)
	sq.mask = make([][]bool, n)
	m := make([]int, n*n)
	mask := make([]bool, n*n)
	for i := range sq.m {
		sq.m[i] = m[n*i : n*(i+1)]
		sq.mask[i] = mask[n*i : n*(i+1)]
	}
}

func (sq *square) String() string {
	var d string
	switch sq.d {
	case FULL:
		d = "soln"
	case HIGH:
		d = "easy"
	case MID:
		d = "med"
	case LOW:
		d = "hard"
	}
	n := len(sq.m)
	var buf bytes.Buffer
	buf.WriteString("[")
	for i := range sq.m {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString("[")
		for j := range sq.m[i] {
			if j != 0 {
				buf.WriteString(",")
			}
			if sq.mask[i][j] {
				buf.WriteString("_")
			} else {
				buf.WriteString(strconv.Itoa(sq.m[i][j]))
			}
		}
		buf.WriteString("]")
	}
	buf.WriteString("]")
	return fmt.Sprintf("test_case(%v, %v, %v, %v).", n, sq.alg, d, buf.String())
}

func (sq *square) ReflectX() {
	n := len(sq.m)
	for y := 0; y < n; y++ {
		for x := 0; x < n/2; x++ {
			sq.m[y][x], sq.m[y][n-x-1] = sq.m[y][n-x-1], sq.m[y][x]
		}
	}
}

func (sq *square) ReflectY() {
	n := len(sq.m)
	for y := 0; y < n/2; y++ {
		for x := 0; x < n; x++ {
			sq.m[y][x], sq.m[n-y-1][x] = sq.m[n-y-1][x], sq.m[y][x]
		}
	}
}

func (sq *square) Transpose() {
	n := len(sq.m)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sq.m[i][j], sq.m[j][i] = sq.m[j][i], sq.m[i][j]
		}
	}
}

func (sq *square) Shuffle() {
	if rand.Float64() < 0.5 {
		sq.ReflectX()
	}
	if rand.Float64() < 0.5 {
		sq.ReflectY()
	}
	if rand.Float64() < 0.5 {
		sq.Transpose()
	}
}

func (sq *square) Mask(d density) {
	n := len(sq.m)
	sq.mask = make([][]bool, n)
	sq.d = d
	arr := make([]bool, n*n)
	for i := range sq.mask {
		sq.mask[i] = arr[n*i : n*(i+1)]
	}
	count := float64(0)
	for {
		for i := range arr {
			if float64(d) <= count/float64(n*n) {
				return
			}
			if !arr[i] && rand.Float64() < 0.5 {
				arr[i] = true
				count++
			}
		}
	}
}

// cursor
// -------------------------

type cursor struct {
	*square
	x, y int
}

func (c *cursor) Read() int {
	return c.m[c.y][c.x]
}

func (c *cursor) Set(n int) {
	c.m[c.y][c.x] = n
}

func (c *cursor) SetPos(x, y int) {
	c.x, c.y = x, y
}

func (c *cursor) Move(dx, dy int) {
	dim := len(c.m)
	c.x += dx
	c.y += dy
	for c.x < 0 {
		c.x += dim
	}
	for c.y < 0 {
		c.y += dim
	}
	c.x %= dim
	c.y %= dim
}

// generators
// -------------------------

func Siamese1(n int) (sq square) {
	if n%2 == 0 {
		panic("Siamese1 must be called with odd order")
	}
	sq.Init(n)
	sq.alg = "siamese1"
	c := cursor{square: &sq}
	c.SetPos(0, n/2)
	for i := 0; i < n*n; i++ {
		c.Set(i + 1)
		c.Move(-1, +1)
		if c.Read() != 0 {
			c.Move(+2, -1)
		}
	}
	return sq
}

func Siamese2(n int) (sq square) {
	if n%2 == 0 {
		panic("Siamese2 must be called with odd order")
	}
	sq.Init(n)
	sq.alg = "siamese2"
	c := cursor{square: &sq}
	c.SetPos(n/2-1, n/2)
	for i := 0; i < n*n; i++ {
		c.Set(i + 1)
		c.Move(-1, +1)
		if c.Read() != 0 {
			c.Move(-1, -1)
		}
	}
	return sq
}

func Fours(n int) (sq square) {
	if n%4 != 0 {
		panic("Fours must called with doubly even order")
	}
	sq.Init(n)
	sq.alg = "mystic"
	c := cursor{square: &sq}
	v := 1
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			c.SetPos(y, x)
			if ((y%4 == 0 || y%4 == 3) && (x%4 == 0 || x%4 == 3)) ||
				((y%4 == 1 || y%4 == 2) && (x%4 == 1 || x%4 == 2)) {
				c.Set(n*n + 1 - v)
			} else {
				c.Set(v)
			}
			v++
		}
	}
	return sq
}

func main() {
	var sq square

	sq = Siamese1(3)
	sq.Shuffle()
	fmt.Println(sq.String())
	sq.Mask(HIGH)
	fmt.Println(sq.String())
	sq.Mask(MID)
	fmt.Println(sq.String())
	sq.Mask(LOW)
	fmt.Println(sq.String())
	fmt.Println()

	sq = Siamese1(5)
	sq.Shuffle()
	fmt.Println(sq.String())
	sq.Mask(HIGH)
	fmt.Println(sq.String())
	sq.Mask(MID)
	fmt.Println(sq.String())
	sq.Mask(LOW)
	fmt.Println(sq.String())
	fmt.Println()

	sq = Siamese1(7)
	sq.Shuffle()
	fmt.Println(sq.String())
	sq.Mask(HIGH)
	fmt.Println(sq.String())
	sq.Mask(MID)
	fmt.Println(sq.String())
	sq.Mask(LOW)
	fmt.Println(sq.String())
	fmt.Println()

	sq = Siamese2(3)
	sq.Shuffle()
	fmt.Println(sq.String())
	sq.Mask(HIGH)
	fmt.Println(sq.String())
	sq.Mask(MID)
	fmt.Println(sq.String())
	sq.Mask(LOW)
	fmt.Println(sq.String())
	fmt.Println()

	sq = Siamese2(5)
	sq.Shuffle()
	fmt.Println(sq.String())
	sq.Mask(HIGH)
	fmt.Println(sq.String())
	sq.Mask(MID)
	fmt.Println(sq.String())
	sq.Mask(LOW)
	fmt.Println(sq.String())
	fmt.Println()

	sq = Siamese2(7)
	sq.Shuffle()
	fmt.Println(sq.String())
	sq.Mask(HIGH)
	fmt.Println(sq.String())
	sq.Mask(MID)
	fmt.Println(sq.String())
	sq.Mask(LOW)
	fmt.Println(sq.String())
	fmt.Println()

	sq = Fours(4)
	sq.Shuffle()
	fmt.Println(sq.String())
	sq.Mask(HIGH)
	fmt.Println(sq.String())
	sq.Mask(MID)
	fmt.Println(sq.String())
	sq.Mask(LOW)
	fmt.Println(sq.String())
	fmt.Println()

	sq = Fours(8)
	sq.Shuffle()
	fmt.Println(sq.String())
	sq.Mask(HIGH)
	fmt.Println(sq.String())
	sq.Mask(MID)
	fmt.Println(sq.String())
	sq.Mask(LOW)
	fmt.Println(sq.String())
	fmt.Println()
}
