# Solving Magic Squares with a Genetic Algorithm

- Term Project Proposal
- Chris Barrick and Manuel Phun

We are attempting to generate and solve so called "magic squares" using a genetic algorithm. A magic square is an arrangement of distinct integers in an n by n grid such that each row, column, and diagonal sum to the same value. When n=3, there is only 1 magic squares, excluding those created by rotation and reflection. When n=4, there are 880 different magic squares, and when n=5, there are 275305224 different magic squares. Because of the combinatorial explosion of the size of the search space, we believe that a genetic algorithm will fair better than deterministic search techniques for large n.

There are actually two different problems with magic squares. The first task is to generate a random square of a given size. In the context of a GA, this task requires generating permutations and checking their feasibility as magic squares. We want to solve the harder problem of completing a partially filled square. This second task is like the first with additional constraints that certain numbers must appear in certain positions. For a deterministic approach like constraint satisfaction, these additional constraints make the problem easier. However, our genetic algorithm would likely search the space of all permutations for feasible magic squares. These additional constraints make that problem harder in the context of a GA because we've reduced the number of solutions without reducing the search space. However, custom operators may allow us to reduce the search space much like what occurs in constraint satisfaction.

Along with the application, we are creating a general framework for evolutionary algorithms in the Go programming language, with a focus on parallelism.
