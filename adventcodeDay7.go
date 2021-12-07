package main

import (
	"math"
	"sort"
)

func Puzzle13(filename string) int {
	nPositions := readfile(filename)
	// sort elements in their positions.
	sort.Ints(nPositions)
	// find max value.
	maxValue := nPositions[len(nPositions)-1]
	minChange := maxValue * len(nPositions)
	for i := 0; i < maxValue; i++ {
		currChange := 0
		for j := 0; j < len(nPositions); j++ {
			diff := int(math.Abs(float64(nPositions[j]) - float64(i)))
			currChange += diff
		}
		if currChange < minChange {
			minChange = currChange
		}
	}

	return minChange

}

func Puzzle14(filename string) int {
	nPositions := readfile(filename)
	// sort elements in their positions.
	sort.Ints(nPositions)
	// find max value.
	maxValue := nPositions[len(nPositions)-1]
	minChange := math.MaxInt64
	for i := 0; i < maxValue; i++ {
		currChange := 0
		for j := 0; j < len(nPositions); j++ {
			diff := int(math.Abs(float64(nPositions[j]) - float64(i)))
			diffsum := 0
			for k := 0; k <= diff; k++ {
				diffsum += k
			}
			currChange += diffsum
		}
		if currChange < minChange {
			minChange = currChange
		}
	}

	return minChange

}
