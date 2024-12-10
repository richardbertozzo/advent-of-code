package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/richardbertozzo/adventofcode-2024/file"
)

var (
	XMAS = map[int]rune{1: 'M', 2: 'A', 3: 'S'}
	MAS  = map[int]rune{1: 'A', 2: 'S'}
)

func countXmasWords(matrix [][]rune) int {
	total := 0

	for i, line := range matrix {
		for y, element := range line {
			if element == 'X' {
				// horizontal
				if matchesPatternInDirection(XMAS, matrix, i, y, 0, 1) {
					total++
				}

				// horizontal backwards
				if matchesPatternInDirection(XMAS, matrix, i, y, 0, -1) {
					total++
				}

				// vertical
				if matchesPatternInDirection(XMAS, matrix, i, y, 1, 0) {
					total++
				}

				// vertical backwards
				if matchesPatternInDirection(XMAS, matrix, i, y, -1, 0) {
					total++
				}

				// diagonal down left
				if matchesPatternInDirection(XMAS, matrix, i, y, 1, -1) {
					total++
				}

				// diagonal down right
				if matchesPatternInDirection(XMAS, matrix, i, y, 1, 1) {
					total++
				}

				// diagonal up right
				if matchesPatternInDirection(XMAS, matrix, i, y, -1, 1) {
					total++
				}

				// diagonal up left
				if matchesPatternInDirection(XMAS, matrix, i, y, -1, -1) {
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
	downRight := matchesPatternInDirection(MAS, matrix, lineIdx, columnIdx, 1, 1)
	if downRight {
		downLeft := matchesPatternInDirection(MAS, matrix, lineIdx, columnIdx+2, 1, -1)
		upRight := matchesPatternInDirection(MAS, matrix, lineIdx+2, columnIdx, -1, 1)
		return downLeft || upRight
	}

	// up left
	upLeft := matchesPatternInDirection(MAS, matrix, lineIdx, columnIdx, -1, -1)
	if upLeft {
		upRight := matchesPatternInDirection(MAS, matrix, lineIdx, columnIdx-2, -1, 1)
		downLeft := matchesPatternInDirection(MAS, matrix, lineIdx-2, columnIdx, 1, -1)
		return upRight || downLeft
	}

	return false
}

func matchesPatternInDirection(steps map[int]rune, matrix [][]rune, lineIdx, columnIdx, rowStep, colStep int) bool {
	// Check if the movement stays within bounds
	for step := 1; step <= len(steps); step++ {
		newRow := lineIdx + step*rowStep
		newCol := columnIdx + step*colStep
		if newRow < 0 || newRow >= len(matrix) || newCol < 0 || newCol >= len(matrix[lineIdx]) {
			return false
		}

		// Match the pattern 'M', 'A', 'S' in respective steps
		expected := steps[step]
		if matrix[newRow][newCol] != expected {
			return false
		}
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
