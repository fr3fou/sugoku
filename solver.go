package main

import "fmt"

// Sudoku ...
type Sudoku [9][9]int

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

// ValidNums returns the valid nums
func (s Sudoku) ValidNums(x, y int) []int {
	// cols := []int{}
	// rows := []int{}
	// sqr := []int{}
	digits := [10]bool{}

	// cols
	for i := 0; i < 9; i++ {
		val := s[i][x]         // will be 1..9
		digits[val] = val != 0 // if it's 0, it's valid
	}

	// rows
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
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			// we don't care
			if s[x][y] != 0 {
				continue
			}

			num := s.solve(x, y)
			s[x][y] = num
		}
	}

	return s
}

func (s Sudoku) solve(x, y int) int {
	nums := s.ValidNums(x, y)

	// no valid nums
	if len(nums) == 0 {
		return -1
	}

	num := -1

outer:
	for _, num = range nums {
		// assume it's the current number
		s[x][y] = num
		for i := x; i < 9; i++ {
			for j := y + 1; j < 9; j++ {
				// we don't care
				if s[i][j] != 0 {
					continue
				}

				guess := s.solve(i, j)
				if guess == -1 {
					// we fucked up
					continue outer
				}
			}
		}
	}

	return num
}
