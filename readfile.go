package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

//read file with has an list of comma separated numbers.
func readfile(filename string) []int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	line := scanner.Text()
	values := strings.Split(line, ",")
	iValues := make([]int, 0, len(values))
	for i := 0; i < len(values); i++ {
		val, _ := strconv.Atoi(values[i])
		iValues = append(iValues, val)
	}

	return iValues
}

// read the file and return 2 arrays of strings
func readfileday8(filename string) ([]string, []string) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	//scanner.Scan()
	leftside := make([]string, 0, 250)
	rightside := make([]string, 0, 250)
	for scanner.Scan() {
		val := scanner.Text()
		sides := strings.Split(val, "|")
		leftside = append(leftside, sides[0])
		rightside = append(rightside, sides[1])
	}

	return leftside, rightside
}

func readfileday9(filename string) [][]int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	matrix := make([][]int, 0, 100)
	for scanner.Scan() {
		line := make([]int, 0, 100)
		val := scanner.Text()

		for _, v := range val {
			intval, _ := strconv.Atoi(string(v))
			line = append(line, intval)
		}
		matrix = append(matrix, line)
	}
	return matrix
}
