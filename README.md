# fem-library
Finite Element Library Implemented in Go
This library is designed for the application of the finite element method in structural engineering problems. It's a work in progress, intended to offer robust tools for analyzing structural behavior.

Usage from the command line:

go run main.go 1  examples/D1/example3.json

In this command:

    The first parameter represents the dimension of the analysis.
    The second parameter is the path to the JSON file describing the problem to be solved.

Additionally, examples are provided, including validation problems sourced from the literature, which can be used to test the library's solutions.
