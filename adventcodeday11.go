package main

import (
	"fmt"
)

func PuzzleDay11Part1(filename string, run int) int {

	mat := readfileday9(filename) // returns a grid/matrix
	// decorate it with an extra row and col around the grid
	mat = decorateMatrix(mat)

	flashcnt := 0
	for i := 0; i < run; i++ {
		incrementbyOne(mat)
		flash(mat)
		flashcnt += resetMatrix(mat)
	}

	return flashcnt
}

func PuzzleDay11Part2(filename string) int {

	mat := readfileday9(filename) // returns a grid/matrix
	// decorate it with an extra row and col around the grid
	mat = decorateMatrix(mat)

	flashcnt := 0
	run := 0
	for flashcnt != 100 {
		incrementbyOne(mat)
		flash(mat)
		flashcnt = resetMatrix(mat)
		run++
	}
	return run
}

func incrementbyOne(mat [][]int) {

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] != -10 {
				mat[i][j]++
			}
		}
	}

}

func flash(mat [][]int) {
	myQueue := Queue{}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] > 9 {
				val := fmt.Sprintf("%d,%d", i, j)
				myQueue.Enqueue(val)
			}
		}
	}

	// array to indicate how to check for left right, down, up and diagonal
	move := [][]int{{1, 0}, {1, -1}, {1, 1}, {0, -1}, {0, 1}, {-1, 0}, {-1, -1}, {-1, 1}}
	for !myQueue.IsEmpty() {
		val, _ := myQueue.Dequeue()
		var r, c int
		fmt.Sscanf(val, "%d,%d", &r, &c)

		for i := 0; i < len(move); i++ {
			newr := r + move[i][0]
			newc := c + move[i][1]
			if mat[newr][newc] != -10 {
				mat[newr][newc]++
				if mat[newr][newc] > 9 {
					newvaltoQueue := fmt.Sprintf("%d,%d", newr, newc)
					if !myQueue.IsPresent(newvaltoQueue) {
						myQueue.Enqueue(newvaltoQueue)
					}
				}
			}
		}
	}
}

func resetMatrix(mat [][]int) int {
	cnt := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] != -10 {
				if mat[i][j] > 9 {
					mat[i][j] = 0
					cnt++
				}
			}
		}
	}
	return cnt
}

func decorateMatrix(mat [][]int) [][]int {
	newMat := make([][]int, 0, len(mat)+2)

	// first row
	newRow := make([]int, len(mat[0])+2)
	for i := 0; i < len(mat[0])+2; i++ {
		newRow[i] = -10
	}
	newMat = append(newMat, newRow)
	for i := 0; i < len(mat); i++ {
		row := make([]int, 0, len(mat[0])+2)
		row = append(row, -10)
		row = append(row, mat[i]...)
		row = append(row, -10)
		newMat = append(newMat, row)
	}
	// last row
	newRow = make([]int, len(mat[0])+2)
	for i := 0; i < len(mat[0])+2; i++ {
		newRow[i] = -10
	}
	newMat = append(newMat, newRow)
	return newMat
}
