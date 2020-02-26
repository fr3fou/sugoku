package main

import (
	"fmt"
	"strings"
)

// Sudoku ...
type Sudoku [9][9]int

// Print prints the sudoku
func (s Sudoku) String() string {
	b := &strings.Builder{}
	for i, line := range s {
		if i%3 == 0 {
			fmt.Fprintln(b, "+---------+---------+---------+")
		}
		for j, num := range line {
			if j%3 == 0 {
				fmt.Fprint(b, "|")
			}
			if num == 0 {
				fmt.Fprint(b, " - ")
			} else {
				fmt.Fprintf(b, " %d ", num)
			}
			if j == 8 {
				fmt.Fprint(b, "|")
			}
		}
		if i == 8 {
			fmt.Fprint(b, "\n+---------+---------+---------+")
		}
		fmt.Fprintln(b)
	}

	return b.String()
}

// ValidNums returns the valid nums
func (s Sudoku) ValidNums(x, y int) []int {
	digits := [10]bool{}

	// check horizontal
	for i := 0; i < 9; i++ {
		val := s[i][y]         // will be 1..9
		digits[val] = val != 0 // if it's 0, it's valid
	}

	for i := 0; i < 9; i++ {
		val := s[x][i]         // will be 1..9
		digits[val] = val != 0 // if it's 0, it's valid
	}

	// square
	topX := x / 3 * 3
	topY := y / 3 * 3

	for i := topX; i < topX+3; i++ {
		for j := topY; j < topY+3; j++ {
			val := s[i][j]         // will be 1..9
			digits[val] = val != 0 // if it's 0, it's valid
		}
	}

	nums := []int{}
	for i := 1; i < len(digits); i++ {
		if !digits[i] {
			nums = append(nums, i)
		}
	}

	return nums
}

// Solve ...
func (s Sudoku) Solve() Sudoku {
	return s.solve(0, 0)
}

func (s Sudoku) solve(x, y int) Sudoku {
	nums := s.ValidNums(x, y)

	// no valid nums
	if len(nums) == 0 {
		s[x][y] = -1
		return s
	}

	og := s[x][y]

outer:
	for _, num := range nums {
		// assume it's the current number
		s[x][y] = num
		for i := x; i < 9; i++ {
			for j := y + 1; j < 9; j++ {
				// we don't care
				if s[i][j] != 0 {
					continue
				}

				if s.solve(i, j)[i][j] == -1 {
					// we fucked up
					// we gotta clean up
					s[i][j] = og
					continue outer
				}
			}
		}
	}

	return s
}
