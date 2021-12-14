package main

import (
	"fmt"
)

func PuzzleDay13Part1(filename string, run int) int {

	pos, folds, maxx, maxy := readfileday13(filename)

	trans := filltransparentSheet(pos, maxx, maxy)

	// assumption is run is less than or equal to number of folds read from file
	for i := 0; i < run; i++ {
		var coord string
		var foldv int
		fmt.Sscanf(folds[i], "%1s,%d", &coord, &foldv)
		trans = fold(trans, coord, foldv)
	}

	cnt := 0
	for i := 0; i < len(trans); i++ {
		for j := 0; j < len(trans[0]); j++ {
			if trans[i][j] == 1 {
				cnt++
			}
		}
	}
	return cnt
}

func PuzzleDay13Part2(filename string, run int) [][]int {

	pos, folds, maxx, maxy := readfileday13(filename)

	trans := filltransparentSheet(pos, maxx, maxy)

	// assumption is run is less than or equal to number of folds read from file
	for i := 0; i < run; i++ {
		var coord string
		var foldv int
		fmt.Sscanf(folds[i], "%1s,%d", &coord, &foldv)
		trans = fold(trans, coord, foldv)
	}

	return trans
}

func filltransparentSheet(pos [][]int, maxx int, maxy int) [][]int {

	trans := make([][]int, 0, maxy+1)
	for i := 0; i < maxy+1; i++ {
		arr := make([]int, maxx+1)
		trans = append(trans, arr)
	}

	for i := 0; i < len(pos); i++ {
		trans[pos[i][1]][pos[i][0]] = 1
	}

	return trans
}

func fold(trans [][]int, fold string, v int) [][]int {

	var maxr, maxc int
	var origmaxr, origmaxc int
	origmaxr = len(trans)
	origmaxc = len(trans[0])

	if fold == "y" {
		maxr = v
		maxc = len(trans[0])
		if v > (origmaxr - v - 1) {
			trans = addRowsBottom(trans, v-(origmaxr-v-1))
			origmaxr += (v - (origmaxr - v - 1))
		} else {
			if (origmaxr - v - 1) > v {
				trans = addRowsTop(trans, (origmaxr-v-1)-v)
				maxr = v + ((origmaxr - v - 1) - v)
				origmaxr += ((origmaxr - v - 1) - v)
			}
		}
	} else {
		maxr = len(trans)
		maxc = v
		if v > (origmaxc - v - 1) {
			trans = addColsEnd(trans, v-(origmaxc-v-1))
			origmaxc += (v - (origmaxc - v - 1))
		} else {
			if (origmaxc - v - 1) > v {
				trans = addColsBegin(trans, (origmaxc-v-1)-v)
				maxc = v + ((origmaxc - v - 1) - v)
				origmaxc += ((origmaxc - v - 1) - v)
			}
		}
	}

	newtrans := make([][]int, 0, maxr)

	for i := 0; i < maxr; i++ {
		r := make([]int, maxc)
		newtrans = append(newtrans, r)
	}

	if fold == "y" {
		for i := 0; i < maxr; i++ {
			for j := 0; j < maxc; j++ {
				newtrans[i][j] = trans[i][j] | trans[origmaxr-i-1][j]
			}
		}
	} else {
		for i := 0; i < maxr; i++ {
			for j := 0; j < maxc; j++ {
				newtrans[i][j] = trans[i][j] | trans[i][origmaxc-j-1]
			}
		}
	}

	return newtrans
}

func addRowsBottom(mat [][]int, cntrows int) [][]int {

	newrcnt := len(mat) + cntrows
	newmat := make([][]int, len(mat), newrcnt)

	copy(newmat, mat)

	for i := 0; i < cntrows; i++ {
		newmat = append(newmat, make([]int, len(mat[0])))
	}

	return newmat

}

func addRowsTop(mat [][]int, cntrows int) [][]int {

	newrcnt := len(mat) + cntrows
	newmat := make([][]int, 0, newrcnt)

	for i := 0; i < cntrows; i++ {
		newmat = append(newmat, make([]int, len(mat[0])))
	}

	for i := 0; i < len(mat); i++ {
		newr := make([]int, len(mat[0]))
		copy(newr, mat[i])
		newmat = append(newmat, newr)

	}

	return newmat

}

func addColsBegin(mat [][]int, cntcols int) [][]int {

	newccnt := len(mat[0]) + cntcols

	newmat := make([][]int, 0, len(mat))

	for i := 0; i < len(mat); i++ {
		newr := make([]int, 0, newccnt)
		newextracols := make([]int, cntcols)
		newr = append(newr, newextracols...)
		newr = append(newr, mat[i]...)
		newmat = append(newmat, newr)
	}

	return newmat
}

func addColsEnd(mat [][]int, cntcols int) [][]int {

	newccnt := len(mat[0]) + cntcols

	newmat := make([][]int, 0, len(mat))

	for i := 0; i < len(mat); i++ {
		newr := make([]int, 0, newccnt)
		newextracols := make([]int, cntcols)
		newr = append(newr, mat[i]...)
		newr = append(newr, newextracols...)
		newmat = append(newmat, newr)
	}

	return newmat
}
