package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

// https://adventofcode.com/2024/day/3

func removeChars(s string) [][]string {
	rx := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)`)
	return rx.FindAllStringSubmatch(s, -1)
}

func removeNumberFromStr(s string) (int, int) {
	rx := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)`)
	match := rx.FindStringSubmatch(s)
	num1, _ := strconv.Atoi(match[1])
	num2, _ := strconv.Atoi(match[2])
	return num1, num2
}

// consider do() and don't() instructions
func removeCharsWithEnables(s string) []string {
	rx := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)|do\(\)|don't\(\)`)
	return rx.FindAllString(s, -1)
}

func calculate() int64 {
	content := readFileContent()
	matches := removeChars(string(content))

	var total int64
	for _, match := range matches {
		if len(match) == 3 { // Ensure two capture groups were found
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])

			total += int64(num1 * num2)
		}
	}

	return total
}

func calculateWithDoDont() int64 {
	content := readFileContent()
	matches := removeCharsWithEnables(string(content))

	var total int64
	shouldCount := true
	for _, match := range matches {
		if match == "don't()" {
			shouldCount = false
			continue
		} else if match == "do()" {
			shouldCount = true
			continue
		}

		if shouldCount {
			// convert the string to get the numbers
			num1, num2 := removeNumberFromStr(match)
			total += int64(num1 * num2)
		}
	}

	return total
}

func readFileContent() []byte {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(pwd + "/day-3/input")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = file.Close()
	}()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

func main() {
	fmt.Println("total", calculate())
	fmt.Println("total", calculateWithDoDont())
}
