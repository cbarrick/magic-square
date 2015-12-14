# Magic Square Solver

The goal of this project is to develop a genetic algorithm for solving [normal magic squares][1]. This repository contains an experiment comparing our genetic algorithm, written in Go, to a constraint satisfaction problem solver, written in Prolog.

[1]: https://en.wikipedia.org/wiki/Magic_square


## Source Layout

- `ga/` and `csp/`: These directories contain the source code for the genetic algorithm and constraint satisfaction solvers respectively.
- `ga_trial.go` and `csp_trial.pl`: These programs provide entry points to run each solver against specific trials. They accept the same command-line flags, given below, and their output is a timing of the trial.
	- `-g <generator>`: The generator identifies the algorithm that generated the trial to be run.
	- `-o <order>`: The order identifies the size of the trial to be run. Only certain combinations of order and generator are valid. See the `trials.pl` file for valid combinations.
	- `-d <difficulty>`: The difficulty of the problem must be one of `hard`, `med`, or `easy`.
	- `-n <count>`: The trial is run repeatedly for some count of iterations and the average time taken is reported.
	- `-t <timeout>`: If any particular iteration takes longer than the timeout, the program halts and "timeout" is printed.
- `trials.pl`: This file contains the trials to run in the experiment.
- `main.sh`: This script runs each of the solvers on each of the trials and prints a comma-separated table of the average time taken for each trial.
- `data.csv`: This is the experimental data collected from `main.sh` and referenced in the report.
