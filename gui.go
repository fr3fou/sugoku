package main

import (
	"strconv"

	"github.com/fr3fou/sugoku/sudoku"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// RenderBG renders the background
func RenderBG(r *sdl.Renderer) error {
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

// RenderLine renders a line
func RenderLine(r *sdl.Renderer, x0, y0, x1, y1 int) error {
	return r.DrawLine(
		int32(x0),
		int32(y0),
		int32(x1),
		int32(y1),
	)
}

// RenderBoard renders only the grid
func RenderBoard(r *sdl.Renderer, font *ttf.Font, board sudoku.Sudoku) error {
	for x, line := range board {
		for y, num := range line {
			if err := RenderCell(r, x, y); err != nil {
				return err
			}
			if err := RenderNum(r, font, x, y, num); err != nil {
				return err
			}
		}
	}

	return nil
}

// RenderCell renders a single square
func RenderCell(r *sdl.Renderer, x, y int) error {
	if err := r.SetDrawColor(0, 0, 0, 0); err != nil {
		return err
	}

	if err := RenderLine(r,
		x*CellSize,
		y*CellSize,
		x*CellSize+CellSize,
		y*CellSize,
	); err != nil {
		return err
	}

	if err := RenderLine(r,
		x*CellSize,
		y*CellSize,
		x*CellSize,
		y*CellSize+CellSize,
	); err != nil {
		return err
	}

	// horizontal edge
	if y%3 == 0 {
		if err := RenderLine(r,
			x*CellSize,
			y*CellSize+1,
			x*CellSize+CellSize,
			y*CellSize+1,
		); err != nil {
			return err
		}

		if err := RenderLine(r,
			x*CellSize,
			y*CellSize-1,
			x*CellSize+CellSize,
			y*CellSize-1,
		); err != nil {
			return err
		}
	}

	// vertical edge
	if x%3 == 0 {
		if err := RenderLine(r,
			x*CellSize+1,
			y*CellSize,
			x*CellSize+1,
			y*CellSize+CellSize,
		); err != nil {
			return err
		}

		if err := RenderLine(r,
			x*CellSize-1,
			y*CellSize,
			x*CellSize-1,
			y*CellSize+CellSize,
		); err != nil {
			return err
		}
	}

	// bottom line
	if y == 8 {
		if err := RenderLine(r,
			x*CellSize,
			(y+1)*CellSize-1,
			x*CellSize+CellSize,
			(y+1)*CellSize-1,
		); err != nil {
			return err
		}

		if err := RenderLine(r,
			x*CellSize,
			(y+1)*CellSize-2,
			x*CellSize+CellSize,
			(y+1)*CellSize-2,
		); err != nil {
			return err
		}
	}

	// left line
	if x == 8 {
		if err := RenderLine(r,
			(x+1)*CellSize-1,
			y*CellSize,
			(x+1)*CellSize-1,
			y*CellSize+CellSize,
		); err != nil {
			return err
		}

		if err := RenderLine(r,
			(x+1)*CellSize-2,
			y*CellSize,
			(x+1)*CellSize-2,
			y*CellSize+CellSize,
		); err != nil {
			return err
		}
	}

	return nil
}

// RenderNum renders only a number
func RenderNum(r *sdl.Renderer, f *ttf.Font, x, y, num int) error {
	if num == 0 {
		return nil
	}

	if err := r.SetDrawColor(0, 0, 0, 0); err != nil {
		return err
	}

	s, err := f.RenderUTF8Blended(strconv.Itoa(num), sdl.Color{})
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

	return r.Copy(t, nil, &sdl.Rect{
		X: int32(x*CellSize) + (CellSize-clip.W)/2, // center horizontally (relies on clip width being accurate)
		Y: int32(y*CellSize) + (CellSize-clip.H)/2, // center vertically (relies on clip height being accurate)
		W: clip.W,
		H: clip.H, // apparently 94px
	})
}
