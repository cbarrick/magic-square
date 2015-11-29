# Magic Square Solver

The goal of this project is to develop a genetic algorithm for solving [normal magic squares][1]. This repository contains an experiment comparing our genetic algorithm, written in Go, to a constraint satisfaction problem solver, written in Prolog.

[1]: https://en.wikipedia.org/wiki/Magic_square


## Source Layout

- `ga/` and `csp/`: These directories contain the source code for the genetic algorithm and constraint satisfaction solvers respectively.
- `gen.go`: This program is used to the generate problem instances used in the experiment. It uses a variety of algorithms to generate partial magic squares of different orders. The output is a valid Prolog knowledge base and is saved as `trials.pl`.
- `ga_trial.go` and `csp_trial.pl`: These programs provide entry points to run each solver against specific trials. They require the same command-line flags, given below:
	- `-g GENERATOR`: This flag specifies the generator of the trial. The generator is the algorithm used to generate the problem instance.
	- `-o ORDER`: This flag specifies the order of the trial. Only certain combinations of order and generator are valid for specific trials. See the `trials.pl` file for valid combinations.
	- `-d DIFFICULTY`: This flag specifies the difficulty of the problem instance and must be one of `hard`, `med`, or `easy`.
