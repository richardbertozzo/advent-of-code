package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/richardbertozzo/adventofcode-2024/file"
)

func countXmasWords(matrix [][]rune) int {
	total := 0

	for i, line := range matrix {
		for y, element := range line {
			if element == 'X' {
				// horizontal
				if horizontal(line, y) {
					total++
				}

				// horizontal backwards
				if horizontalBackwards(line, y) {
					total++
				}

				// vertical
				if vertical(matrix, i, y) {
					total++
				}

				// vertical backwards
				if verticalBackwards(matrix, i, y) {
					total++
				}

				// diagonal down left
				if diagonalDownLeft(matrix, i, y) {
					total++
				}

				// diagonal down right
				if diagonalDownRight(matrix, i, y) {
					total++
				}

				// diagonal up right
				if diagonalUpRight(matrix, i, y) {
					total++
				}

				// diagonal up left
				if diagonalUpLeft(matrix, i, y) {
					total++
				}
			}
		}
	}

	return total
}

func countMasInXWords(matrix [][]rune) int {
	total := 0

	for i, line := range matrix {
		for y, element := range line {
			if element == 'M' {
				// diagonal down left
				if masX(matrix, i, y) {
					total++
				}
			}
		}
	}

	return total
}

func masX(matrix [][]rune, lineIdx, columnIdx int) bool {
	if lineIdx+2 > len(matrix) && lineIdx-2 < 0 {
		if matrix[lineIdx+2][columnIdx] != 'M' {
			return false
		}
	}

	if columnIdx+2 >= len(matrix[lineIdx]) && columnIdx-2 < 0 {
		if matrix[lineIdx][columnIdx+2] != 'M' {
			return false
		}
	}

	// down right
	downRight := checkDiagonal(matrix, lineIdx, columnIdx, 1, 1)
	if downRight {
		downLeft := checkDiagonal(matrix, lineIdx, columnIdx+2, 1, -1)
		upRight := checkDiagonal(matrix, lineIdx+2, columnIdx, -1, 1)
		return downLeft || upRight
	}

	// up left
	upLeft := checkDiagonal(matrix, lineIdx, columnIdx, -1, -1)
	if upLeft {
		upRight := checkDiagonal(matrix, lineIdx, columnIdx-2, -1, 1)
		downLeft := checkDiagonal(matrix, lineIdx-2, columnIdx, 1, -1)
		return upRight || downLeft
	}

	return false
}

func checkDiagonal(matrix [][]rune, lineIdx, columnIdx, rowStep, colStep int) bool {
	// Check if the movement stays within bounds
	for step := 1; step <= 2; step++ {
		newRow := lineIdx + step*rowStep
		newCol := columnIdx + step*colStep
		if newRow < 0 || newRow >= len(matrix) || newCol < 0 || newCol >= len(matrix[lineIdx]) {
			return false
		}

		// Match the pattern 'A', 'S' in respective steps
		expected := map[int]rune{1: 'A', 2: 'S'}[step]
		if matrix[newRow][newCol] != expected {
			return false
		}
	}

	return true
}

func horizontal(line []rune, idx int) bool {
	if idx+3 > len(line)-1 {
		return false
	}

	if line[idx+1] != 'M' {
		return false
	} else if line[idx+2] != 'A' {
		return false
	} else if line[idx+3] != 'S' {
		return false
	}

	return true
}

func horizontalBackwards(line []rune, idx int) bool {
	if idx-3 < 0 {
		return false
	}

	if line[idx-1] != 'M' {
		return false
	} else if line[idx-2] != 'A' {
		return false
	} else if line[idx-3] != 'S' {
		return false
	}

	return true
}

func vertical(matrix [][]rune, lineIdx, columnIdx int) bool {
	if lineIdx+3 > len(matrix) {
		return false
	}

	if matrix[lineIdx+1][columnIdx] != 'M' {
		return false
	} else if matrix[lineIdx+2][columnIdx] != 'A' {
		return false
	} else if matrix[lineIdx+3][columnIdx] != 'S' {
		return false
	}

	return true
}

func verticalBackwards(matrix [][]rune, lineIdx, columnIdx int) bool {
	if lineIdx-3 < 0 {
		return false
	}

	if matrix[lineIdx-1][columnIdx] != 'M' {
		return false
	} else if matrix[lineIdx-2][columnIdx] != 'A' {
		return false
	} else if matrix[lineIdx-3][columnIdx] != 'S' {
		return false
	}

	return true
}

func diagonalDownRight(matrix [][]rune, lineIdx, columnIdx int) bool {
	if lineIdx+3 > len(matrix)-1 || columnIdx+3 > len(matrix[lineIdx])-1 {
		return false
	}

	if matrix[lineIdx+1][columnIdx+1] != 'M' {
		return false
	} else if matrix[lineIdx+2][columnIdx+2] != 'A' {
		return false
	} else if matrix[lineIdx+3][columnIdx+3] != 'S' {
		return false
	}

	return true
}

func diagonalUpRight(matrix [][]rune, lineIdx, columnIdx int) bool {
	if lineIdx-3 < 0 || columnIdx+3 > len(matrix[lineIdx])-1 {
		return false
	}

	if matrix[lineIdx-1][columnIdx+1] != 'M' {
		return false
	} else if matrix[lineIdx-2][columnIdx+2] != 'A' {
		return false
	} else if matrix[lineIdx-3][columnIdx+3] != 'S' {
		return false
	}

	return true
}

func diagonalUpLeft(matrix [][]rune, lineIdx, columnIdx int) bool {
	if lineIdx-3 < 0 || columnIdx-3 < 0 {
		return false
	}

	if matrix[lineIdx-1][columnIdx-1] != 'M' {
		return false
	} else if matrix[lineIdx-2][columnIdx-2] != 'A' {
		return false
	} else if matrix[lineIdx-3][columnIdx-3] != 'S' {
		return false
	}

	return true
}

func diagonalDownLeft(matrix [][]rune, lineIdx, columnIdx int) bool {
	if lineIdx+3 > len(matrix)-1 || columnIdx-3 < 0 {
		return false
	}

	if matrix[lineIdx+1][columnIdx-1] != 'M' {
		return false
	} else if matrix[lineIdx+2][columnIdx-2] != 'A' {
		return false
	} else if matrix[lineIdx+3][columnIdx-3] != 'S' {
		return false
	}

	return true
}

func convertInputToSlice(s string) [][]rune {
	lines := strings.Split(s, "\n")

	result := make([][]rune, len(lines))
	for i, line := range lines {
		result[i] = []rune(line)
	}

	return result
}

func main() {
	content, err := file.ReadFileContent("/day-4/input")
	if err != nil {
		log.Fatal(err)
	}
	input := convertInputToSlice(string(content))

	c := countXmasWords(input) // 2397
	fmt.Println("total", c)

	c2 := countMasInXWords(input)
	fmt.Println("total part 2", c2) // 1824
}
