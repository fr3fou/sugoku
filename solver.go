package main

import "fmt"

// Sudoku ...
type Sudoku [][]int

// Print prints the sudoku
func (s Sudoku) Print() {
	for _, line := range s {
		for _, num := range line {
			if num == 0 {
				fmt.Print("- ")
			} else {
				fmt.Printf("%d ", num)
			}
		}
		fmt.Println()
	}
}

// Solve ...
func (s Sudoku) Solve() Sudoku {
	if len(s) != len(s[0]) {
		panic("must be square")
	}

	return s
}
