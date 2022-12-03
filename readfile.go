package main

import (
	"bufio"
	"fmt"
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

// assumption is that file contains less than 100 rows
func readfile2(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	strArr := make([]string, 0, 100)
	for scanner.Scan() {
		line := scanner.Text()

		strArr = append(strArr, line)
	}
	return strArr
}

func readfileday13(filename string) ([][]int, []string, int, int) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	posArr := make([][]int, 0, 1000)
	foldArr := make([]string, 0, 25)
	maxx := 0
	maxy := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if len(line) > 10 {
			var axis string
			foldn := 0
			fmt.Sscanf(line, "fold along %1s=%d", &axis, &foldn)
			s := fmt.Sprintf("%s,%d", axis, foldn)
			foldArr = append(foldArr, s)
		} else {
			lr := strings.Split(line, ",")
			newr := make([]int, 2)
			newr[0], _ = strconv.Atoi(lr[0])
			newr[1], _ = strconv.Atoi(lr[1])
			if newr[0] > maxx {
				maxx = newr[0]
			}
			if newr[1] > maxy {
				maxy = newr[1]
			}
			posArr = append(posArr, newr)
		}

	}
	return posArr, foldArr, maxx, maxy
}

func readfileday14(filename string) (string, [][]string) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	firstline := scanner.Text()
	scanner.Scan()
	scanner.Text() // empty 2nd line .. ignore.

	insertpos := make([][]string, 0, 1000)
	for scanner.Scan() {
		val := scanner.Text()
		var lside, rside string
		fmt.Sscanf(val, "%s -> %s", &lside, &rside)
		temp := make([]string, 2)
		temp[0] = lside
		temp[1] = rside
		insertpos = append(insertpos, temp)
	}

	return firstline, insertpos
}
