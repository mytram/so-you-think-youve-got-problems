# So you think you've got problems?

`So you think you've got problems?` is a puzzle book by Alex Bellos. Here are the notes and thoughts I take whilst trying to solve the puzzles.

## Problem 001 - Number conundrums 1) on page 7

Let us label the six small white circles from top to bottom and from left to right as `a`, `b`, `c`, `d`, `e`, `f`, it gives us a set of equations and the puzzle becomes an algebraic problem.

```
a + c + e + d = 14
d + b + c + f = 14
a + b + e + f = 14

```

where `a`, `b`, `c`, `d`, `e`, and `f` are each a number from `1-6`. One solution can be quickly found out by trial-and-error.

However, it will not be surprising to see that there are multiple solutions to the puzzle. It will be easier to use a programme to find out all the solutions.

A brute-force algorithm is to generate the permutations of the array `[1, 2, 3, 4, 5, 6]`. Each combination represents a possible configuration of `a, b, e, d, e, f`. When a configuration solves the aforementioned equations, it is a solution.

In total, the programme calculates 48 solutions.
