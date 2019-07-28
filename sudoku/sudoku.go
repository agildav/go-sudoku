package sudoku

import "fmt"

// Sudoku board
type Sudoku struct {
	board   [][]int
	empty   int
	rows    int
	columns int
}

// Init initializes a new sudoku board
func (s *Sudoku) Init(r int, c int) {
	s.rows = r
	s.columns = c
	s.empty = 0

	s.board = make([][]int, r)
	for row := 0; row < r; row++ {
		s.board[row] = make([]int, c)
	}
}

// PrintSudoku prints the sudoku board
func (s *Sudoku) PrintSudoku() {
	rows := s.rows
	columns := s.columns

	for r := 0; r < rows; r++ {
		if r%3 == 0 && r != 0 {
			fmt.Println("---------------------")
		}

		for c := 0; c < columns; c++ {
			if c%3 == 0 && c != 0 {
				fmt.Print("| ")
			}

			fmt.Print(s.board[r][c], " ")
		}

		fmt.Println()
	}
}

func (s *Sudoku) isValidInRow(r int, n int) bool {
	for c := 0; c < s.columns; c++ {
		if s.board[r][c] == n {
			return false
		}
	}
	return true
}

func (s *Sudoku) isValidInColumn(c int, n int) bool {
	for r := 0; r < s.rows; r++ {
		if s.board[r][c] == n {
			return false
		}
	}
	return true
}

func (s *Sudoku) isValidInBlock(r int, c int, n int) bool {
	rlimit := r - r%3
	climit := c - c%3

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {

			if s.board[r+rlimit][c+climit] == n {
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) isNumberAllowed(r int, c int, n int) bool {
	return s.isValidInRow(r, n) && s.isValidInColumn(c, n) && s.isValidInBlock(r, c, n)
}

// ResolveSudoku solves the sudoku
func (s *Sudoku) ResolveSudoku() bool {
	rows := s.rows
	columns := s.columns
	e := s.empty

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if s.board[r][c] == e {
				for n := 1; n <= 9; n++ {
					if s.isNumberAllowed(r, c, n) {
						s.board[r][c] = n
						if s.ResolveSudoku() {
							return true
						}
						s.board[r][c] = e
					}

				}
				return false
			}
		}
	}
	return true
}
