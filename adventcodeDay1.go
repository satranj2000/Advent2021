package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// file open and scan
func Puzzle1() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	//var text []string

	totcnt := 0
	prev := 1000000 // very high value , so that it ignores the first row.
	//curr := 0
	for scanner.Scan() {
		//println(scanner.Text())
		curr, _ := strconv.Atoi(scanner.Text())
		if curr > prev {
			totcnt = totcnt + 1
		}
		prev = curr
	}
	println(totcnt)
	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()
}

func Readinputdata() []int {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var depths []int
	for scanner.Scan() {
		//println(scanner.Text())
		val, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, val)
	}
	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()
	return depths
}

// read in group of 3 and compare the outputs
//https://adventofcode.com/2021/day/1#part2
func Puzzle2(depths []int) int {
	totcnt := 0
	depthcnt := depths[0] + depths[1] + depths[2]
	for i := 1; i+2 < len(depths); i++ {
		depth2 := depths[i] + depths[i+1] + depths[i+2]
		if depth2 > depthcnt {
			totcnt++
		}
		depthcnt = depth2
	}
	return totcnt
}
