package sudoku

// Sudoku ...
type Sudoku [9][9]int

type Solver interface {
	String() string
	ValidNums(int, int) string
	Solve() Sudoku
}
