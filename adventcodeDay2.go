package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// read day 2 input data
func Puzzle3() int {
	file, err := os.Open("input2.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	forwardMove := 0
	depthMove := 0
	for scanner.Scan() {
		//println(scanner.Text())
		action := strings.Split(scanner.Text(), " ")
		val, _ := strconv.Atoi(action[1])
		if action[0] == "forward" {
			forwardMove += val
		}
		if action[0] == "down" {
			depthMove += val
		}
		if action[0] == "up" {
			depthMove -= val
		}
	}
	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	return forwardMove * depthMove
}

//https://adventofcode.com/2021/day/2#part2
func Puzzle4() int {
	file, err := os.Open("input2.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	forwardMove := 0
	depthMove := 0
	aim := 0
	for scanner.Scan() {
		//println(scanner.Text())
		action := strings.Split(scanner.Text(), " ")
		val, _ := strconv.Atoi(action[1])
		if action[0] == "forward" {
			forwardMove += val
			depthMove += val * aim
		}
		if action[0] == "down" {
			aim += val
		}
		if action[0] == "up" {
			aim -= val
		}

	}
	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	return forwardMove * depthMove
}
