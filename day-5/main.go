package main

import (
	"errors"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/richardbertozzo/adventofcode-2024/file"
)

func main() {
	orderRules, err := file.ReadFileContent("/day-5/input")
	if err != nil {
		log.Fatal(err)
	}

	rules, err := convertInputToOrderRules(string(orderRules))
	if err != nil {
		log.Fatal(err)
	}

	rows, err := file.ReadFileContent("/day-5/input_order")
	if err != nil {
		log.Fatal(err)
	}

	inputsRows, err := convertInputsRows(string(rows))
	if err != nil {
		log.Fatal(err)
	}

	total := sumMiddlePageNumberValidRows(inputsRows, rules)
	log.Println("total", total)
}

func sumMiddlePageNumberValidRows(rows [][]int, rules map[int][]int) int {
	total := 0
	for _, row := range rows {
		if isValidRow(row, rules) {
			total += row[getMiddlePageNumber(row)]
		}
	}

	return total

}

func isValidRow(row []int, rules map[int][]int) bool {
	prev := row[0]
	valid := true

	for i := 1; i < len(row) && valid; i++ {
		next := row[i]
		nextNums := rules[prev]

		if !slices.Contains(nextNums, next) {
			valid = false
		}

		prev = next
	}

	return valid
}

func getMiddlePageNumber(row []int) int {
	return len(row) / 2
}

func convertInputToOrderRules(s string) (map[int][]int, error) {
	lines := strings.Split(s, "\n")

	result := make(map[int][]int)
	for _, line := range lines {
		split := strings.Split(line, "|")

		if len(split) != 2 {
			return nil, errors.New("input order rule size different than 2")
		}

		var before, next int
		var err error
		if before, err = strconv.Atoi(split[0]); err != nil {
			return nil, err
		}
		if next, err = strconv.Atoi(split[1]); err != nil {
			return nil, err
		}

		if r, found := result[before]; found {
			result[before] = append(r, next)
		} else {
			result[before] = []int{next}
		}
	}

	return result, nil
}

func convertInputsRows(s string) ([][]int, error) {
	lines := strings.Split(s, "\n")

	result := make([][]int, len(lines))
	for i, line := range lines {
		valuesStr := strings.Split(line, ",")

		values := make([]int, len(valuesStr))
		for x, str := range valuesStr {
			num, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			values[x] = num
		}

		result[i] = values
	}

	return result, nil
}
