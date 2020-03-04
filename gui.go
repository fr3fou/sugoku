package main

import (
	"strconv"

	"github.com/fr3fou/sudogo/sudoku"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
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

func renderLine(r *sdl.Renderer, x, y, xx, yy int) error {
	return r.DrawLine(
		int32(x),
		int32(y),
		int32(xx),
		int32(yy),
	)
}

func renderBoard(r *sdl.Renderer, f *ttf.Font, board sudoku.Sudoku) error {
	for x, line := range board {
		for y, num := range line {
			if err := renderCell(r, f, &sudoku.Cell{Num: num, X: x, Y: y}); err != nil {
				return err
			}
		}
	}

	return nil
}

func renderCell(r *sdl.Renderer, f *ttf.Font, cell *sudoku.Cell) error {
	if err := r.SetDrawColor(0, 0, 0, 0); err != nil {
		return err
	}

	if err := renderLine(r,
		cell.X*CellSize,
		cell.Y*CellSize,
		cell.X*CellSize+CellSize,
		cell.Y*CellSize,
	); err != nil {
		return err
	}

	if err := renderLine(r,
		cell.X*CellSize,
		cell.Y*CellSize,
		cell.X*CellSize,
		cell.Y*CellSize+CellSize,
	); err != nil {
		return err
	}

	// horizontal edge
	if cell.Y%3 == 0 {
		if err := renderLine(r,
			cell.X*CellSize,
			cell.Y*CellSize+1,
			cell.X*CellSize+CellSize,
			cell.Y*CellSize+1,
		); err != nil {
			return err
		}

		if err := renderLine(r,
			cell.X*CellSize,
			cell.Y*CellSize-1,
			cell.X*CellSize+CellSize,
			cell.Y*CellSize-1,
		); err != nil {
			return err
		}
	}

	// vertical edge
	if cell.X%3 == 0 {
		if err := renderLine(r,
			cell.X*CellSize+1,
			cell.Y*CellSize,
			cell.X*CellSize+1,
			cell.Y*CellSize+CellSize,
		); err != nil {
			return err
		}

		if err := renderLine(r,
			cell.X*CellSize-1,
			cell.Y*CellSize,
			cell.X*CellSize-1,
			cell.Y*CellSize+CellSize,
		); err != nil {
			return err
		}
	}
	// bottom line
	if cell.Y == 8 {
		if err := renderLine(r,
			cell.X*CellSize,
			(cell.Y+1)*CellSize-1,
			cell.X*CellSize+CellSize,
			(cell.Y+1)*CellSize-1,
		); err != nil {
			return err
		}

		if err := renderLine(r,
			cell.X*CellSize,
			(cell.Y+1)*CellSize-2,
			cell.X*CellSize+CellSize,
			(cell.Y+1)*CellSize-2,
		); err != nil {
			return err
		}
	}
	// left line
	if cell.X == 8 {
		if err := renderLine(r,
			(cell.X+1)*CellSize-1,
			cell.Y*CellSize,
			(cell.X+1)*CellSize-1,
			cell.Y*CellSize+CellSize,
		); err != nil {
			return err
		}

		if err := renderLine(r,
			(cell.X+1)*CellSize-2,
			cell.Y*CellSize,
			(cell.X+1)*CellSize-2,
			cell.Y*CellSize+CellSize,
		); err != nil {
			return err
		}
	}

	if cell.Num == 0 {
		return nil
	}

	s, err := f.RenderUTF8Blended(strconv.Itoa(cell.Num), sdl.Color{})
	if err != nil {
		return err
	}
	defer s.Free()

	clip := &sdl.Rect{}
	s.GetClipRect(clip)
	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return err
	}
	r.Copy(t, nil, &sdl.Rect{
		X: int32(cell.X*CellSize) + (CellSize-clip.W)/2, // center horizontally (relies on clip width being accurate)
		Y: int32(cell.Y*CellSize) + (CellSize-clip.H)/2, // center vertically (relies on clip height being accurate)
		W: clip.W,
		H: clip.H, // apparently 94px
	})

	return nil
}
