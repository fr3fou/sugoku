package main

func main() {
	var s Sudoku = [][]int{
		{1, 0, 3, 4, 5},
		{1, 2, 3, 0, 0},
		{1, 0, 3, 0, 5},
		{1, 2, 3, 4, 0},
		{1, 2, 0, 4, 5},
	}

	s.Print()
}
