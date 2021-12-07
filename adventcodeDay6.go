package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput6File(filename string) []int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	line := scanner.Text()
	strLaternFishs := strings.Split(line, ",")

	LaternFishs := make([]int, 0, 1000)

	for i := 0; i < len(strLaternFishs); i++ {
		val, _ := strconv.Atoi(strLaternFishs[i])
		LaternFishs = append(LaternFishs, val)
	}

	return LaternFishs
}

func Puzzle11(file string, cnt int) int {
	initiallaternFishs := readInput6File(file)
	laternfishs := make([]int, len(initiallaternFishs), 10000)
	copy(laternfishs, initiallaternFishs)
	for i := 0; i < cnt; i++ {
		l := len(laternfishs)
		for j := 0; j < l; j++ {
			if laternfishs[j] == 0 {
				laternfishs = append(laternfishs, 8)
				laternfishs[j] = 6
			} else {
				laternfishs[j]--
			}

		}
	}

	return len(laternfishs)
}

// another version of puzzle11 and 12 without using a large slice.
func Puzzle11_v2(file string, cnt int) int {
	initiallaternFishs := readInput6File(file)

	// keep a map of number of fishes per count.
	laterncnt := len(initiallaternFishs)
	mapLtCnt := make(map[int]int, 9)
	for i := 0; i < laterncnt; i++ {
		mapLtCnt[initiallaternFishs[i]]++
	}
	// this gives a map of 9 - each with a laternfish count per number.

	for i := 0; i < cnt; i++ {
		// in every loop; find out how many zeros are present.
		zeroCnt := mapLtCnt[0]
		// move cnt in each interval to one level below.
		for j := 1; j <= 8; j++ {
			mapLtCnt[j-1] = mapLtCnt[j]
		}
		// move the zeros count to 8th
		// add zeros count to whatever is already present in 6th interal.
		mapLtCnt[8] = zeroCnt
		mapLtCnt[6] += zeroCnt
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += mapLtCnt[i]
	}
	return sum
}
