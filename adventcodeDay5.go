package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func Puzzle9() int {
	file, err := os.Open("input5.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	pointsVal := make(map[string]int)

	var x1, y1, x2, y2 int
	for scanner.Scan() {
		val := scanner.Text()
		fmt.Sscanf(val, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		// if either of the two coordinates is same
		if x1 == x2 || y1 == y2 {
			if x1 == x2 {
				start := int64(math.Min(float64(y1), float64(y2)))
				end := int64(math.Max(float64(y1), float64(y2)))
				for i := start; i <= end; i++ {
					key := fmt.Sprintf("%d,%d", x1, i)
					pointsVal[key]++
				}
			}
			if y1 == y2 {
				start := int64(math.Min(float64(x1), float64(x2)))
				end := int64(math.Max(float64(x1), float64(x2)))
				for i := start; i <= end; i++ {
					key := fmt.Sprintf("%d,%d", i, y1)
					pointsVal[key]++
				}
			}
		}
	}

	cnt := 0

	for _, v := range pointsVal {
		if v > 1 {
			cnt++
		}
	}

	return cnt

}
