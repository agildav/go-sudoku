package main

import (
	"fmt"
	s "go-sudoku/sudoku"
	"time"
)

func main() {
	beg := time.Now()
	sudoku := new(s.Sudoku)
	var n = 9
	sudoku.Init(n, n)

	fmt.Print(": Initial Sudoku : \n")
	sudoku.PrintSudoku()

	sudoku.ResolveSudoku()
	end := time.Since(beg)
	fmt.Println("\n: Solution in", end, ":")
	sudoku.PrintSudoku()
}
