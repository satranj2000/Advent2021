package main

import (
	"sort"
	"strconv"
	"strings"
)

func Puzzle15(filename string) int {
	_, rside := readfileday8(filename)
	cnt := 0
	for i := 0; i < len(rside); i++ {
		digits := strings.Split(rside[i], " ")
		for j := 0; j < len(digits); j++ {
			digit := strings.Trim(digits[j], " ")
			// len of digits for numbers 1 , 4, 7, 8
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				cnt++
			}
		}
	}
	return cnt
}

func Puzzle16(filename string) int {
	lsides, rsides := readfileday8(filename)

	sum := 0

	for i := 0; i < len(lsides); i++ {
		// find pattern for each row
		patterns := findDigitPattern(strings.Split(lsides[i], " "))

		// find the positional value of each of right sides.
		word := make([]int, 4)
		// split the right side
		valstring := strings.Split(strings.Trim(rsides[i], " "), " ")
		for j := 0; j < 4; j++ {
			word[j] = findPosition(patterns, sortString(valstring[j]))
		}
		val, _ := strconv.Atoi(IntArrToString(word))
		sum += val

	}

	return sum
}

func findPosition(digits []string, s string) int {
	for i, d := range digits {
		if strings.Compare(d, s) == 0 {
			return i
		}
	}
	return -1
}

func findDigitPattern(lsides []string) []string {
	patterns := make([]string, 10)
	// first set the easy to find patterns.
	for _, digit := range lsides {
		if len(digit) == 2 {
			patterns[1] = sortString(digit)
		}
		if len(digit) == 4 {
			patterns[4] = sortString(digit)
		}
		if len(digit) == 3 {
			patterns[7] = sortString(digit)
		}
		if len(digit) == 7 {
			patterns[8] = sortString(digit)
		}
	}

	// harder to find digits
	for _, digit := range lsides {
		if len(digit) == 6 {
			if IsCharsPresent(digit, patterns[7]) && IsCharsPresent(digit, patterns[4]) {
				patterns[9] = sortString(digit)
			} else {
				if IsCharsPresent(digit, patterns[1]) && IsCharsPresent(digit, patterns[7]) {
					patterns[0] = sortString(digit)
				} else {
					patterns[6] = sortString(digit)
				}
			}
		}
	}

	// might have to loop until all 3 , 5 and 2 values are found.
	for patterns[3] == "" || patterns[5] == "" || patterns[2] == "" {
		for _, digit := range lsides {
			if len(digit) == 5 {
				if IsCharsPresent(digit, patterns[7]) {
					patterns[3] = sortString(digit)
				} else {
					if IsCharsPresent(patterns[9], digit) {
						patterns[5] = sortString(digit)
					} else if patterns[9] != "" {
						patterns[2] = sortString(digit)
					}
				}
			}
		}

	}
	return patterns
}

func IsCharsPresent(s string, chk string) bool {
	for _, c := range chk {
		if !strings.Contains(s, string(c)) {
			return false
		}
	}
	return true
}

func sortString(s string) string {
	w := []rune(s)
	sort.Slice(w, func(i int, j int) bool { return w[i] < w[j] })
	return string(w)
}

func IntArrToString(arr []int) string {
	str := ""
	for _, v := range arr {
		str += strconv.Itoa(v)
	}
	return str
}
