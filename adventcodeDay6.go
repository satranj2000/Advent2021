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
