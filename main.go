package main

import (
	"log"

	"github.com/fr3fou/sugoku/sudoku"
	"github.com/hajimehoshi/ebiten"
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	return nil
}

func main() {
	// board := sudoku.Sudoku{
	// 	{0, 0, 8 /**/, 0, 0, 2 /**/, 0, 5, 0},
	// 	{0, 4, 0 /**/, 0, 0, 5 /**/, 0, 0, 8},
	// 	{0, 3, 5 /**/, 6, 0, 0 /**/, 2, 0, 7},
	// 	/*									*/
	// 	{3, 0, 1 /**/, 0, 2, 0 /**/, 0, 0, 0},
	// 	{5, 0, 0 /**/, 0, 7, 0 /**/, 0, 0, 1},
	// 	{0, 0, 0 /**/, 5, 9, 0 /**/, 8, 0, 3},
	// 	/*									*/
	// 	{7, 0, 3 /**/, 0, 0, 4 /**/, 0, 9, 0},
	// 	{0, 0, 0 /**/, 2, 0, 0 /**/, 0, 8, 0},
	// 	{0, 5, 0 /**/, 0, 0, 0 /**/, 6, 0, 0},
	// }

	// ch := make(chan sudoku.Cell)
	// solve(board, ch)
	// for _ = range ch {
	// 	// fmt.Printf("%d at %d x %d\n", cell.Num, cell.X, cell.Y)
	// }
	if err := ebiten.Run(update, 900, 900, 1, "Sugoku"); err != nil {
		log.Fatal(err)
	}
}

func solve(board sudoku.Sudoku, ch chan sudoku.Cell) {
	solver := sudoku.Solver{
		Board: &board,
		Cells: ch,
	}

	go solver.Solve()
}
