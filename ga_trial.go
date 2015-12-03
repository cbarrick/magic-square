package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cbarrick/magic-square/ga"
)

// command line arguments
var (
	spec  = flag.String("t", "siamese1,easy", "gives trial to run")
	order = flag.Int("o", 3, "gives the order of the trial")
	count = flag.Int("n", 10, "the number of times to run the trial")
)

func main() {
	flag.Parse()
	var avg float64
	t := strings.Split(*spec, ",")
	if t[0] == "generate" {
		start := time.Now()
		for i := 0; i < *count; i++ {
			ga.Generate(*order)
		}
		avg = float64(time.Since(start).Nanoseconds()) / 1000000000 / float64(*count)
	} else {
		schema := getSchema(t[0], *order, t[1])
		start := time.Now()
		for i := 0; i < *count; i++ {
			ga.Solve(schema)
		}
		avg = float64(time.Since(start).Nanoseconds()) / 1000000000 / float64(*count)
	}
	fmt.Printf("%v",  avg)
}

// parser
// -------------------------
// Everything below this point is helper code to read trials.
// Use getSchema to retreive a specific trial from the input.

type trial struct {
	order  int    // order of the square
	gen    string // algorithm generating the square
	dif    string // difficulty of the schema
	schema []int  // decoded schema
}

type unexpectedErr struct{}

func (unexpectedErr) Error() string {
	return "unexpected input"
}

var (
	trials = make([]trial, 0, 40)
	input  *bufio.Reader
)

func getSchema(gen string, order int, dif string) []int {
	if input == nil {
		f, err := os.Open("./trials.pl")
		if err != nil {
			panic(err.Error())
		}
		input = bufio.NewReader(f)
	}
	for _, t := range trials {
		if t.order == order && t.gen == gen && t.dif == dif {
			return t.schema
		}
	}
	for {
		t, err := readTrial(input)
		if err != nil {
			panic(err.Error())
		}
		trials = append(trials, t)
		if t.order == order && t.gen == gen && t.dif == dif {
			return t.schema
		}
	}
}

func readTrial(r *bufio.Reader) (t trial, err error) {
	char := make([]byte, 1)
	char, err = r.Peek(1)
	if err != nil {
		return t, err
	}

	switch string(char) {

	// skip comments
	case "%":
		err = skipline(r)
		if err != nil {
			return t, err
		}
		return readTrial(r)

	// skip blank lines
	case "\n":
		err = expect(r, "\n")
		if err != nil {
			return t, err
		}
		return readTrial(r)

	// start parsing a trial
	case "t":
		err = expect(r, "trial(")
		if err != nil {
			return t, err
		}
		t.order, t.gen, t.dif, err = readMetadata(r)
		if err != nil {
			return t, err
		}
		t.schema, err = readSquare(r, t.order*t.order)
		if err != nil {
			return t, err
		}
		err = skipline(r)
		if err != nil {
			return t, err
		}
		return t, nil

	default:
		return t, unexpectedErr{}
	}
}

func skipline(r *bufio.Reader) (err error) {
	var char rune
	for {
		char, _, err = r.ReadRune()
		if err != nil {
			return err
		}
		if char == '\n' {
			return nil
		}
	}
}

func readMetadata(r *bufio.Reader) (order int, gen, dif string, err error) {
	// order
	str, err := r.ReadString(byte(','))
	if err != nil {
		return order, gen, dif, err
	}
	order, err = strconv.Atoi(str[:len(str)-1])
	if err != nil {
		return order, gen, dif, err
	}

	// generator
	str, err = r.ReadString(byte(','))
	gen = str[:len(str)-1]
	if err != nil {
		return order, gen, dif, err
	}

	// difficulty
	str, err = r.ReadString(byte(','))
	dif = str[:len(str)-1]
	if err != nil {
		return order, gen, dif, err
	}

	return order, gen, dif, nil
}

func readSquare(r *bufio.Reader, n int) (schema []int, err error) {
	err = expect(r, "[")
	if err != nil {
		return schema, err
	}
	depth := 1
	schema = make([]int, 0, n)
	for {
		char, err := r.Peek(1)
		if err != nil {
			return schema, err
		}
		switch string(char) {
		case "[":
			err = expect(r, "[")
			if err != nil {
				return schema, err
			}
			depth++
		case "]":
			err = expect(r, "]")
			if err != nil {
				return schema, err
			}
			depth--
			if depth == 0 {
				return schema, nil
			}
		case ",":
			err = expect(r, ",")
			if err != nil {
				return schema, err
			}
		case "_":
			err = expect(r, "_")
			if err != nil {
				return schema, err
			}
			schema = append(schema, -1)
		default:
			bytes := make([]byte, 0, 3)
			for {
				char, err := r.Peek(1)
				if err != nil {
					return schema, err
				}
				if byte('0') <= char[0] && char[0] <= byte('9') {
					err = expect(r, string(char))
					if err != nil {
						return schema, err
					}
					bytes = append(bytes, char[0])
				} else {
					break
				}
			}
			n, err := strconv.Atoi(string(bytes))
			if err != nil {
				return schema, err
			}
			schema = append(schema, n-1)
		}
	}
}

func expect(r *bufio.Reader, str string) (err error) {
	n := len(str)
	bytes := make([]byte, n)
	_, err = r.Read(bytes)
	if err != nil {
		return err
	}
	if string(bytes) != str {
		panic(fmt.Sprintf("expected '%v', found '%v'", str, bytes))
	}
	return nil
}
