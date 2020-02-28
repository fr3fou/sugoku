package sudoku

import (
	"fmt"
	"strings"
)

type SolverDefault struct {
	Board *Sudoku
}

// Print prints the sudoku
func (s *SolverDefault) String() string {
	b := &strings.Builder{}
	for i, line := range s.Board {
		if i%3 == 0 {
			fmt.Fprintln(b, "+---------+---------+---------+")
		}
		for j, num := range line {
			if j%3 == 0 {
				fmt.Fprint(b, "|")
			}
			if num == 0 {
				fmt.Fprint(b, " . ")
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
func (s *SolverDefault) ValidNums(x, y int) []int {
	digits := [10]bool{}

	// check horizontal
	for i := 0; i < 9; i++ {
		val := s.Board[i][y]   // will be 1..9
		digits[val] = val != 0 // if it's 0, it's valid
	}

	for i := 0; i < 9; i++ {
		val := s.Board[x][i]   // will be 1..9
		digits[val] = val != 0 // if it's 0, it's valid
	}

	// square
	topX := x / 3 * 3
	topY := y / 3 * 3

	for i := topX; i < topX+3; i++ {
		for j := topY; j < topY+3; j++ {
			val := s.Board[i][j]   // will be 1..9
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

// Solve solves the sudoku
func (s *SolverDefault) Solve() Sudoku {
	return *s.solve(0, 0)
}

func (s *SolverDefault) solve(x, y int) *Sudoku {
	nums := s.ValidNums(x, y)

	// base case (bottom right corner)
	if x == 8 && y == 8 && len(nums) == 1 {
		s.Board[x][y] = nums[0] // write the last num
		return s.Board          // success!
	}

	for _, num := range nums {
		s.Board[x][y] = num // assume it's the correct one

		n := s.solve(s.next(x, y)) // recur
		if n != nil {
			return n // return if not nil (has reached the base case)
		}
		s.Board[x][y] = 0 // fix it if it's not (ADDED because impl is now mutable)
	}

	return nil // we failed
}

// next finds the next number that is available
func (s *SolverDefault) next(x, y int) (int, int) {
	for s.Board[x][y] != 0 {
		y++
		if y > 8 { // if we have reached the end of the row
			y = 0 // go to the next row
			x++
		}
	}

	return x, y
}
