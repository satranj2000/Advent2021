package main

import (
	"sort"

	"sathish.com/adventcode/utils"
)

func PuzzleDay10_Part1(filename string) int {

	chunks := readfile2(filename)

	corruptList := make([]string, 0, len(chunks))
	for _, chunk := range chunks {
		posvalue, isCorrupted := CheckIsCorrupted(chunk)
		if isCorrupted {
			corruptList = append(corruptList, posvalue)
		}
	}

	sum := 0
	for i := 0; i < len(corruptList); i++ {
		if corruptList[i] == ")" {
			sum += 3
		}
		if corruptList[i] == "]" {
			sum += 57
		}
		if corruptList[i] == "}" {
			sum += 1197
		}
		if corruptList[i] == ">" {
			sum += 25137
		}
	}

	return sum
}

func PuzzleDay10Part2(filename string) int {
	chunks := readfile2(filename)

	incompleteList := make([]string, 0, len(chunks))
	for _, chunk := range chunks {
		_, isCorrupted := CheckIsCorrupted(chunk)
		if !isCorrupted {
			incompleteList = append(incompleteList, chunk)
		}
	}

	points := make([]int, 0, len(incompleteList))
	for i := 0; i < len(incompleteList); i++ {
		score, _ := findCompletionCharsScore(incompleteList[i])
		points = append(points, score)
	}
	sort.Ints(points)

	return points[(len(points)-1)/2]
}

func findCompletionCharsScore(chunk string) (int, string) {
	myStack := utils.Stack{}

	for i := 0; i < len(chunk); i++ {
		if string(chunk[i]) == "{" || string(chunk[i]) == "[" ||
			string(chunk[i]) == "<" || string(chunk[i]) == "(" {
			myStack.Push(string(chunk[i]))
		} else {
			poppedChar, _ := myStack.Pop()
			if string(chunk[i]) == "}" && poppedChar == "{" {
				continue
			}
			if string(chunk[i]) == "]" && poppedChar == "[" {
				continue
			}
			if string(chunk[i]) == ">" && poppedChar == "<" {
				continue
			}
			if string(chunk[i]) == ")" && poppedChar == "(" {
				continue
			}
		}
	}

	sum := 0
	completedChars := ""
	for !myStack.IsEmpty() {
		poppedChar, _ := myStack.Pop()
		if poppedChar == "{" {
			completedChars += "}"
		}
		if poppedChar == "[" {
			completedChars += "]"
		}
		if poppedChar == "<" {
			completedChars += ">"
		}
		if poppedChar == "(" {
			completedChars += ")"
		}
	}

	chars := []string{")", "]", "}", ">"}
	for i := 0; i < len(completedChars); i++ {
		pos := -1
		for j := 0; j < len(chars); j++ {
			if chars[j] == string(completedChars[i]) {
				pos = j + 1
				break
			}
		}

		sum = (5 * sum) + pos
	}
	return sum, completedChars
}

func CheckIsCorrupted(chunk string) (string, bool) {
	myStack := utils.Stack{}
	for i := 0; i < len(chunk); i++ {
		if string(chunk[i]) == "{" || string(chunk[i]) == "[" ||
			string(chunk[i]) == "<" || string(chunk[i]) == "(" {
			myStack.Push(string(chunk[i]))
		} else {
			if myStack.IsEmpty() {
				return string(chunk[i]), true
			}
			poppedChar, _ := myStack.Pop()
			if string(chunk[i]) == "}" && poppedChar != "{" {
				return "}", true
			}
			if string(chunk[i]) == "]" && poppedChar != "[" {
				return "]", true
			}
			if string(chunk[i]) == ">" && poppedChar != "<" {
				return ">", true
			}
			if string(chunk[i]) == ")" && poppedChar != "(" {
				return ")", true
			}
		}
	}
	return "", false
}
