package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		desc     string
		input    Sudoku
		expected Sudoku
	}{
		{
			desc: "simple",
			input: Sudoku{
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
			},
			expected: Sudoku{
				{9, 7, 8 /**/, 1, 4, 2 /**/, 3, 5, 6},
				{2, 4, 6 /**/, 7, 3, 5 /**/, 9, 1, 8},
				{1, 3, 5 /**/, 6, 8, 9 /**/, 2, 4, 7},
				/*									*/
				{3, 8, 1 /**/, 4, 2, 6 /**/, 5, 7, 9},
				{5, 9, 2 /**/, 3, 7, 8 /**/, 4, 6, 1},
				{4, 6, 7 /**/, 5, 9, 1 /**/, 8, 2, 3},
				/*									*/
				{7, 2, 3 /**/, 8, 6, 4 /**/, 1, 9, 5},
				{6, 1, 9 /**/, 2, 5, 3 /**/, 7, 8, 4},
				{8, 5, 4 /**/, 9, 1, 7 /**/, 6, 3, 2},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			solver := SolverDefault{&tt.input}
			assert.Equal(t, tt.expected, solver.Solve())
		})
	}
}
