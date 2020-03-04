package main

import (
	"github.com/fr3fou/sudogo/sudoku"
	"github.com/veandco/go-sdl2/sdl"
)

func renderBg(r *sdl.Renderer) {
	err := r.SetDrawColor(0, 0, 0, 0)
	if err != nil {
		panic(err)
	}

	err = r.FillRect(&sdl.Rect{
		X: 0,
		Y: 0,
		W: Width,
		H: Height,
	})
	if err != nil {
		panic(err)
	}
}

func renderBoard(r *sdl.Renderer, board sudoku.Sudoku) {
	err := r.SetDrawColor(255, 255, 255, 255)
	if err != nil {
		panic(err)
	}

	for x, line := range board {
		for y := range line {
			rect := sdl.Rect{
				X: int32(x*CellSize) + 1,
				Y: int32(y*CellSize) + 1,
				W: CellSize - 2,
				H: CellSize - 2,
			}
			err = r.FillRect(&rect)
			if err != nil {
				panic(err)
			}
		}
	}
}
