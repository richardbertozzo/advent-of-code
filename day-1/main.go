package main

import (
	"math"
	"sort"
)

// day 1 challenge: https://adventofcode.com/2024/day/1
func distanceBetweenList(l1, l2 []int) int {
	sort.Ints(l1)
	sort.Ints(l2)

	distance := 0
	for i := 0; i < len(l1); i++ {
		distance += int(math.Abs(float64(l2[i] - l1[i])))
	}

	return distance
}

func appearanceOnTheSecondList(l1, l2 []int) int {
	numAppearance := make(map[int]int, len(l2))

	for _, n2 := range l2 {
		if _, ok := numAppearance[n2]; ok {
			numAppearance[n2]++
		} else {
			numAppearance[n2] = 1
		}
	}

	score := 0
	for _, n1 := range l1 {
		found := numAppearance[n1]
		score += n1 * found
	}

	return score
}
