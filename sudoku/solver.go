package sudoku

import "errors"

type Solver struct {
	Board     *Sudoku
	Snapshots chan Sudoku
}

var (
	ErrUnsolvableSudoku = errors.New("sudoku: sudoku is unsolvable")
)

func (s *Solver) Write(x, y, num int) {
	s.Board[x][y] = num
	if s.Snapshots != nil {
		s.snapshot()
	}
}

func (s *Solver) snapshot() {
	s.Snapshots <- *s.Board
}

// ValidNums returns the valid nums
func (s *Solver) ValidNums(x, y int) []int {
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
func (s *Solver) Solve() (Sudoku, error) {
	n := s.solve(0, 0)

	if n == nil {
		return Sudoku{}, ErrUnsolvableSudoku
	}

	return *n, nil
}

func (s *Solver) solve(x, y int) *Sudoku {
	nums := s.ValidNums(x, y)

	// base case (bottom right corner)
	if x == 8 && y == 8 && len(nums) == 1 {
		s.Write(x, y, nums[0]) // write the last num
		if s.Snapshots != nil {
			close(s.Snapshots)
		}

		return s.Board // success!
	}

	for _, num := range nums {
		s.Write(x, y, num) // assume it's the correct one

		i, j := s.next(x, y)

		// no free cells available
		if i == -1 && j == -1 {
			if s.Snapshots != nil {
				close(s.Snapshots)
			}

			return s.Board
		}

		n := s.solve(i, j) // recur
		if n != nil {
			return n // return if not nil (has reached the base case)
		}

		s.Write(x, y, 0) // clean up
	}

	return nil // we failed
}

// next finds the next number that is available
// returns -1, -1 if there are no free nums left
func (s *Solver) next(x, y int) (int, int) {
	for x, line := range s.Board {
		for y := range line {
			if s.Board[x][y] == 0 {
				return x, y
			}
		}
	}

	return -1, -1
}
