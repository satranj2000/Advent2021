package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput4() ([]int, [][][5]int) {
	file, err := os.Open("input4.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	// read the first line and add the values for the bingo
	callNums := make([]int, 0, 255)
	firstlinestr := scanner.Text()
	firstLine := strings.Split(firstlinestr, ",")
	for i := 0; i < len(firstLine); i++ {
		val, _ := strconv.Atoi(firstLine[i])
		callNums = append(callNums, val)
	}
	cnt := 0
	boards := make([][][5]int, 0, 255)
	board := [][5]int{}

	for scanner.Scan() {
		// ignore the first one - empty string.
		val := scanner.Text() // ignore
		if cnt == 0 {
			//ignore
			cnt++
		} else {
			if cnt < 6 {
				// append to the board array.
				board = append(board, [5]int{})
				line := strings.Split(val, " ")
				// generally the length of line will be 5.
				// but if there single digit characters are present, it will be more than that
				// so, we need to ignore them.
				m := 0
				for i := 0; i < len(line); i++ {
					// ignore cases where there is a space (like single digit characters)
					if line[i] == "" {
						continue
					} else {
						board[cnt-1][m], _ = strconv.Atoi(line[i])
						m++
					}

				}
				cnt++
			} else { // this would also be an empty string.
				cnt = 1
				boards = append(boards, board)
				board = nil
				board = [][5]int{}
			}
		}

	}
	boards = append(boards, board)

	return callNums, boards
}

func Puzzle7() int {
	calls, boards := readInput4()
	// go through each item in the calls and check if the board is complete.
	for i := 0; i < len(calls); i++ {
		for j := 0; j < len(boards); j++ {
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					if boards[j][k][l] == calls[i] {
						boards[j][k][l] = -1
					}
				}
			}
		}
		// check if any of rows or cols in the boards are equal to -1.
		// if not continue.
		// if yes, then calculate the remaining sum and multiply by the calls[i] value
		b, _ := whichBoardComplete(boards)
		if b == -1 {
			continue
		} else {
			remSum := calculateRemainingSum(boards[b])
			return remSum * calls[i]
		}
	}
	return -1
}

func Puzzle8() int {
	calls, boards := readInput4()
	completedList := make([]int, 0, 100)
	lastval := 0
	// go through each item in the calls and check if the board is complete.
	for i := 0; i < len(calls); i++ {
		for j := 0; j < len(boards); j++ {
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					if boards[j][k][l] == calls[i] {
						boards[j][k][l] = -1
					}
				}
			}
		}
		completedList = countBoardsCompleted(boards, completedList)
		if len(completedList) == 100 {
			lastval = calls[i]
			break
		}
	}
	l := len(completedList)
	sumRem := calculateRemainingSum(boards[completedList[l-1]])
	return sumRem * lastval
}

func calculateRemainingSum(board [][5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] != -1 {
				sum += board[i][j]
			}
		}
	}
	return sum
}

func countBoardsCompleted(boards [][][5]int, completedList []int) []int {
	for j := 0; j < len(boards); j++ {
		bPresent := false
		for k := 0; k < len(completedList); k++ {
			if completedList[k] == j {
				bPresent = true
			}
		}
		if bPresent {
			continue
		}
		bComplete := false
		for r := 0; r < 5; r++ {
			if isRowComplete(boards[j], r) {
				bComplete = true
				break
			}
		}
		if bComplete {
			completedList = append(completedList, j)
			continue // ignore checking for columns
		}
		bComplete = false
		for c := 0; c < 5; c++ {
			if isColComplete(boards[j], c) {
				bComplete = true
				break
			}
		}
		if bComplete {
			completedList = append(completedList, j)
			continue // ignore checking for columns
		}
	}
	return completedList
}

func whichBoardComplete(boards [][][5]int) (int, int) {
	for j := 0; j < len(boards); j++ {
		for r := 0; r < 5; r++ {
			if isRowComplete(boards[j], r) {
				return j, r
			}
		}
		for c := 0; c < 5; c++ {
			if isColComplete(boards[j], c) {
				return j, c
			}
		}
	}
	return -1, -1

}

func isRowComplete(board [][5]int, r int) bool {
	sum := board[r][0] + board[r][1] + board[r][2] + board[r][3] + board[r][4]
	return sum == -5

}

func isColComplete(board [][5]int, c int) bool {
	sum := board[0][c] + board[1][c] + board[2][c] + board[3][c] + board[4][c]
	return sum == -5
}
