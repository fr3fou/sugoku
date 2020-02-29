package sudoku

type Cell struct {
	X   int
	Y   int
	Num int
}

type Cells chan Cell
