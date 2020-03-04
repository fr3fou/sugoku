package main

import (
	"github.com/fr3fou/sudogo/sudoku"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	CellSize = 100
	White    = 0xFFFFFF
	Black    = 0x000000
	Width    = 900
	Height   = 900
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Sudoku", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		Width, Height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

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

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	if err != nil {
		panic(err)
	}

	renderBg(renderer)
	renderBoard(renderer, board)
	renderer.Present()

	// ch := make(chan sudoku.Cell)
	// solve(board, ch)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
			}
		}
	}

}

func solve(board sudoku.Sudoku, ch chan sudoku.Cell) {
	solver := sudoku.Solver{
		Board: &board,
		Cells: ch,
	}

	go solver.Solve()
}
