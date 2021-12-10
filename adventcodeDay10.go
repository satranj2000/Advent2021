package main

import "sathish.com/adventcode/utils"

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
