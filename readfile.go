package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

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
