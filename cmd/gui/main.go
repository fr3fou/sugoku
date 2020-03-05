package main

import (
	"github.com/fr3fou/sugoku/gui"
	"github.com/fr3fou/sugoku/sudoku"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Sudoku", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		gui.Width, gui.Height, sdl.WINDOW_SHOWN)
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
	defer renderer.Destroy()

	if err := ttf.Init(); err != nil {
		panic(err)
	}
	defer ttf.Quit()

	font, err := ttf.OpenFont("assets/fonts/roboto.ttf", 80)
	if err != nil {
		panic(err)
	}
	defer font.Close()

	ch := make(chan sudoku.Sudoku)

	solving := false
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseButtonEvent:
				if e.State == sdl.PRESSED || !solving {
					solving = true
					go solve(board, ch)
				}
			}
		}

		renderer.Clear()

		if err := gui.RenderBG(renderer); err != nil {
			panic(err)
		}

		if solving {
			newBoard, ok := <-ch
			if ok {
				board = newBoard
			}
		}

		if err := gui.RenderBoard(renderer, font, board); err != nil {
			panic(err)
		}

		renderer.Present()
	}
}

func solve(board sudoku.Sudoku, ch chan sudoku.Sudoku) {
	ch <- board
	solver := sudoku.Solver{
		Board:     &board,
		Snapshots: ch,
	}

	go solver.Solve()
}
