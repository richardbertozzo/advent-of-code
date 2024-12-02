package main

import (
	"fmt"
	"math"
)

// https://adventofcode.com/2024/day/2

func countSafeReports(reports [][]int) int {
	total := 0

	for x, report := range reports {

		valid := true
		head := report[0]
		isDec := head < report[1]

		for i := 1; i < len(report); i++ {
			if report[i] == head ||
				math.Abs(float64(head-report[i])) > 3 ||
				isDec && report[i] < head ||
				!isDec && report[i] > head {

				valid = false
				break
			}

			head = report[i]
		}

		if valid {
			fmt.Println("line", x, report)
			total++
		}
	}

	return total
}
