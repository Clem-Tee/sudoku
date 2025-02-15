package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const size = 9

func main() {
	// Read arguments
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	grid, err := parseInput(os.Args[1:])
	if err != nil {
		fmt.Println("Error")
		return
	}

	if solveSudoku(&grid) {
		printGrid(grid)
	} else {
		fmt.Println("Error")
	}
}

func parseInput(args []string) ([size][size]int, error) {
	var grid [size][size]int
	for i, line := range args {
		if len(line) != size {
			return grid, errors.New("invalid row length")
		}
		for j, char := range line {
			if char == '.' {
				grid[i][j] = 0
			} else if char >= '1' && char <= '9' {
				grid[i][j] = int(char - '0')
			} else {
				return grid, errors.New("invalid character in input")
			}
		}
	}
	return grid, nil
}

func printGrid(grid [size][size]int) {
	for i := 0; i < size; i++ {
		fmt.Println(strings.Trim(fmt.Sprint(grid[i]), "[]"))
	}
}

func solveSudoku(grid *[size][size]int) bool {
	row, col, empty := findEmptyCell(grid)
	if !empty {
		return true
	}

	for num := 1; num <= 9; num++ {
		if isValid(grid, row, col, num) {
			grid[row][col] = num
			if solveSudoku(grid) {
				return true
			}
			grid[row][col] = 0
		}
	}
	return false
}

func findEmptyCell(grid *[size][size]int) (int, int, bool) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func isValid(grid *[size][size]int, row, col, num int) bool {
	for i := 0; i < size; i++ {
		if grid[row][i] == num || grid[i][col] == num {
			return false
		}
	}
	boxRow, boxCol := (row/3)*3, (col/3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[boxRow+i][boxCol+j] == num {
				return false
			}
		}
	}
	return true
}
