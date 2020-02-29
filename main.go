package main

import (
	"fmt"

	"github.com/fr3fou/sudogo/sudoku"
)

func main() {
	board := sudoku.Sudoku{
		{0, 0, 8 /**/, 0, 0, 2 /**/, 0, 5, 0},
		{0, 4, 0 /**/, 0, 0, 5 /**/, 0, 0, 8},
		{0, 3, 5 /**/, 6, 0, 0 /**/, 2, 0, 7},
		/*									*/
		{3, 0, 1 /**/, 0, 2, 0 /**/, 0, 0, 0},
		{5, 0, 0 /**/, 0, 7, 0 /**/, 0, 0, 1},
		{0, 0, 0 /**/, 5, 9, 0 /**/, 8, 0, 3},
		/*									*/
		{7, 0, 3 /**/, 0, 0, 4 /**/, 0, 9, 0},
		{0, 0, 0 /**/, 2, 0, 0 /**/, 0, 8, 0},
		{0, 5, 0 /**/, 0, 0, 0 /**/, 6, 0, 0},
	}

	solver := sudoku.Solver{
		Board: &board,
		Cells: make(chan sudoku.Cell),
	}

	ch := make(chan sudoku.Sudoku)
	go func() {
		ch <- solver.Solve()
	}()

	for _ = range solver.Cells {
		// fmt.Printf("%d at %d x %d\n", cell.Num, cell.X, cell.Y)
	}

	fmt.Println(<-ch)
}
