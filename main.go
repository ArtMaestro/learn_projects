package main

import "fmt"

const size = 9

func solveSudoku(board [][]int) bool {
	row, col := 0, 0
	if !findEmptyCell(board, &row, &col) {
		return true
	}

	for num := 1; num <= size; num++ {
		if isSafe(board, row, col, num) {
			board[row][col] = num

			if solveSudoku(board) {
				return true
			}

			board[row][col] = 0
		}
	}

	return false
}

func findEmptyCell(board [][]int, row *int, col *int) bool {
	for *row = 0; *row < size; *row++ {
		for *col = 0; *col < size; *col++ {
			if board[*row][*col] == 0 {
				return true
			}
		}
	}
	return false
}

func isSafe(board [][]int, row int, col int, num int) bool {
	return !usedInRow(board, row, num) &&
		!usedInColumn(board, col, num) &&
		!usedInBox(board, row-row%3, col-col%3, num)
}

func usedInRow(board [][]int, row int, num int) bool {
	for col := 0; col < size; col++ {
		if board[row][col] == num {
			return true
		}
	}
	return false
}

func usedInColumn(board [][]int, col int, num int) bool {
	for row := 0; row < size; row++ {
		if board[row][col] == num {
			return true
		}
	}
	return false
}

func usedInBox(board [][]int, startRow int, startCol int, num int) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if board[row+startRow][col+startCol] == num {
				return true
			}
		}
	}
	return false
}

func printBoard(board [][]int) {
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			fmt.Printf("%d ", board[row][col])
		}
		fmt.Println()
	}
}

func main() {
	board := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	fmt.Println("Исходная судоку:")
	printBoard(board)

	if solveSudoku(board) {
		fmt.Println("\nРешенная судоку:")
		printBoard(board)
	} else {
		fmt.Println("\nСудоку нерешаемо!")
	}
}
