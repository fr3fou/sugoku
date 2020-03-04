package main

import (
	"github.com/fr3fou/sudogo/sudoku"
	"github.com/veandco/go-sdl2/sdl"
)

func renderBg(r *sdl.Renderer) error {
	err := r.SetDrawColor(255, 255, 255, 255)
	if err != nil {
		return err
	}

	return r.FillRect(&sdl.Rect{
		X: 0,
		Y: 0,
		W: Width,
		H: Height,
	})
}

func renderBoard(r *sdl.Renderer, board sudoku.Sudoku) error {
	for x, line := range board {
		for y := range line {
			err := renderCell(r, &sudoku.Cell{Num: 0, X: x, Y: y})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func renderCell(r *sdl.Renderer, cell *sudoku.Cell) error {
	err := r.SetDrawColor(0, 0, 0, 0)
	if err != nil {
		return err
	}

	err = r.DrawLine(
		int32(cell.X*CellSize),
		int32(cell.Y*CellSize),
		int32(cell.X*CellSize+CellSize),
		int32(cell.Y*CellSize),
	)
	if err != nil {
		return err
	}

	err = r.DrawLine(
		int32(cell.X*CellSize),
		int32(cell.Y*CellSize),
		int32(cell.X*CellSize),
		int32(cell.Y*CellSize+CellSize),
	)
	if err != nil {
		return err
	}

	return nil
}
