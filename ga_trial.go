package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/cbarrick/magic-square/ga"
)

var synopsis = `synopsis:
	This program runs the genetic algorithm experiment. A file named
	'trials.pl' containing the possible trials to run must exist in the current
	working directory.
`

// value of command line arguments
var (
	gen     = flag.String("g", "siamese1", "the generator type of the trial.")
	order   = flag.Int("o", 3, "the order of the trial.")
	dif     = flag.String("d", "easy", "the dificulty of the trial.")
	count   = flag.Int("n", 1, "the number of times to run the trial.")
	timeout = flag.Int("t", 30, "timeout in seconds")
	help    = flag.Bool("h", false, "show this help message.")
)

// printHelp prints a usage message
func printHelp() {
	name := filepath.Base(os.Args[0])
	fmt.Fprintln(os.Stderr, "usage:\n\t", name, "<args>\n")
	fmt.Fprintln(os.Stderr, synopsis)
	fmt.Fprintln(os.Stderr, "arguments:")
	flag.PrintDefaults()
}

// main fetches a schema from trials.pl and tries to solve it with the GA.
func main() {
	// read the command line arguments and get the schema
	flag.Parse()
	if *help {
		printHelp()
		return
	}
	schema := getSchema()

	// run the ga and print the average time of the trials
	var avg float64
	ret := make(chan ga.Square, 1)
	stop := time.After(time.Duration(*timeout) * time.Second)
	start := time.Now()
	for i := 0; i < *count; i++ {
		go ga.Solve(schema, ret)
		select {
		case <-ret:
		case <-stop:
			fmt.Print("timeout")
			return
		}
	}
	avg = float64(time.Since(start).Nanoseconds()) / 1e9 / float64(*count)
	fmt.Println(avg)
}

// Parser
// --------------------------------------------------
// Everything below this point is helper code to read trials from the file.
// Use getSchema to retreive a specific trial.

type trial struct {
	order  int       // order of the square
	gen    string    // algorithm generating the square
	dif    string    // difficulty of the schema
	schema ga.Square // decoded schema
}

type unexpectedErr struct{}

func (unexpectedErr) Error() string {
	return "unexpected input"
}

func getSchema() ga.Square {
	f, err := os.Open("./trials.pl")
	defer f.Close()
	if err != nil {
		panic(err.Error())
	}
	input := bufio.NewReader(f)
	for {
		t, err := readTrial(input)
		if err != nil {
			panic(err.Error())
		}
		if t.order == *order && t.gen == *gen && t.dif == *dif {
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
		t.schema, err = readSquare(r, t.order)
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

func readSquare(r *bufio.Reader, order int) (schema ga.Square, err error) {
	err = expect(r, "[")
	if err != nil {
		return schema, err
	}
	depth := 1
	schema = ga.NewSquare(order)
	i, j := 0, 0
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
			i++
			j = 0
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
			j++
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
			schema.Sq[i][j] = n
			j++
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
