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
			if err := renderCell(r, &sudoku.Cell{Num: 0, X: x, Y: y}); err != nil {
				return err
			}
		}
	}

	return nil
}

func renderCell(r *sdl.Renderer, cell *sudoku.Cell) error {
	if err := r.SetDrawColor(0, 0, 0, 0); err != nil {
		return err
	}

	if err := r.DrawLine(
		int32(cell.X*CellSize),
		int32(cell.Y*CellSize),
		int32(cell.X*CellSize+CellSize),
		int32(cell.Y*CellSize),
	); err != nil {
		return err
	}

	if err := r.DrawLine(
		int32(cell.X*CellSize),
		int32(cell.Y*CellSize),
		int32(cell.X*CellSize),
		int32(cell.Y*CellSize+CellSize),
	); err != nil {
		return err
	}

	// horizontal edge
	if cell.Y%3 == 0 {
		if err := r.DrawLine(
			int32(cell.X*CellSize),
			int32(cell.Y*CellSize+1),
			int32(cell.X*CellSize+CellSize),
			int32(cell.Y*CellSize+1),
		); err != nil {
			return err
		}

		if err := r.DrawLine(
			int32(cell.X*CellSize),
			int32(cell.Y*CellSize-1),
			int32(cell.X*CellSize+CellSize),
			int32(cell.Y*CellSize-1),
		); err != nil {
			return err
		}
	}

	// vertical edge
	if cell.X%3 == 0 {
		if err := r.DrawLine(
			int32(cell.X*CellSize+1),
			int32(cell.Y*CellSize),
			int32(cell.X*CellSize+1),
			int32(cell.Y*CellSize+CellSize),
		); err != nil {
			return err
		}

		if err := r.DrawLine(
			int32(cell.X*CellSize-1),
			int32(cell.Y*CellSize),
			int32(cell.X*CellSize-1),
			int32(cell.Y*CellSize+CellSize),
		); err != nil {
			return err
		}
	}

	return nil
}
