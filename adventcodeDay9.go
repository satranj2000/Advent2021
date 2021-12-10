package main

import (
	"sort"
	"strconv"
	"strings"
)

func PuzzleDay9Part1(filename string) int {

	mat := readfileday9(filename)
	rowlen := len(mat)
	collen := len(mat[0])
	sum := 0
	for i := 0; i < rowlen; i++ {
		for j := 0; j < collen; j++ {
			// if i ==0 and j is 0; check right and down
			if i == 0 && j == 0 {
				if mat[i][j] < mat[i][j+1] && mat[i][j] < mat[i+1][j] {
					sum += (mat[i][j] + 1)
				}
			}
			// if row is 0 and col = max; check to left and below
			if i == 0 && j == collen-1 {
				if mat[i][j] < mat[i][j-1] && mat[i][j] < mat[i+1][j] {
					sum += (mat[i][j] + 1)
				}
			}

			// if i ==max row and j is 0; check above and right
			if i == rowlen-1 && j == 0 {
				if mat[i][j] < mat[i][j+1] && mat[i][j] < mat[i-1][j] {
					sum += (mat[i][j] + 1)
				}
			}
			// if row is 0 and col = max; check for above and left
			if i == rowlen-1 && j == collen-1 {
				if mat[i][j] < mat[i][j-1] && mat[i][j] < mat[i-1][j] {
					sum += (mat[i][j] + 1)
				}
			}
			// if row == 0 and anything other than 0th or max col
			if i == 0 && j < collen-1 && j != 0 {
				if mat[i][j] < mat[i][j-1] && mat[i][j] < mat[i+1][j] && mat[i][j] < mat[i][j+1] {
					sum += (mat[i][j] + 1)
				}
			}
			// if row == max and anything other than 0th or max col
			if i == rowlen-1 && j < collen-1 && j != 0 {
				if mat[i][j] < mat[i][j-1] && mat[i][j] < mat[i-1][j] && mat[i][j] < mat[i][j+1] {
					sum += (mat[i][j] + 1)
				}
			}

			// if any row with column == 0 (except 0 and max)
			if i < rowlen-1 && i != 0 && j == 0 {
				if mat[i][j] < mat[i][j+1] && mat[i][j] < mat[i+1][j] && mat[i][j] < mat[i-1][j] {
					sum += (mat[i][j] + 1)
				}
			}
			// if any row with column == max (except 0 and max rows)
			if i < rowlen-1 && i != 0 && j == collen-1 {
				if mat[i][j] < mat[i][j-1] && mat[i][j] < mat[i-1][j] && mat[i][j] < mat[i+1][j] {
					sum += (mat[i][j] + 1)
				}
			}

			// all other cases
			if i < rowlen-1 && i != 0 && j < collen-1 && j != 0 {
				if mat[i][j] < mat[i][j-1] && mat[i][j] < mat[i][j+1] && mat[i][j] < mat[i+1][j] && mat[i][j] < mat[i-1][j] {
					sum += (mat[i][j] + 1)
				}
			}

		}
	}
	return sum
}

func PuzzleDay9Part2(filename string) int {

	mat := readfileday9(filename)
	lowpoints := findlowpoints(mat)

	basins := make([]int, 0, 500)
	for i := 0; i < len(lowpoints); i++ {
		val := BFSUntil9Reached(mat, lowpoints[i][0], lowpoints[i][1])
		basins = append(basins, val)
	}

	sort.Ints(basins)
	l := len(basins)
	return basins[l-1] * basins[l-2] * basins[l-3]

}

func findlowpoints(mat [][]int) [][2]int {
	rowlen := len(mat)
	collen := len(mat[0])
	lowpoints := make([][2]int, 0, rowlen*collen)
	for i := 0; i < rowlen; i++ {
		for j := 0; j < collen; j++ {
			// if i ==0 and j is 0; check right and down
			if i == 0 && j == 0 {
				if mat[i][j] < mat[i][j+1] && mat[i][j] < mat[i+1][j] {
					//sum += (mat[i][j] + 1)
					lowp := [2]int{i, j}
					lowpoints = append(lowpoints, lowp)
				}
			}
			// if row is 0 and col = max; check to left and below
			if i == 0 && j == collen-1 {
				if mat[i][j] < mat[i][j-1] && mat[i][j] < mat[i+1][j] {
					//sum += (mat[i][j] + 1)
					lowp := [2]int{i, j}
					lowpoints = append(lowpoints, lowp)
				}
			}

			// if i ==max row and j is 0; check above and right
			if i == rowlen-1 && j == 0 {
				if mat[i][j] < mat[i][j+1] && mat[i][j] < mat[i-1][j] {
					//sum += (mat[i][j] + 1)
					lowp := [2]int{i, j}
					lowpoints = append(lowpoints, lowp)
				}
			}
			// if row is 0 and col = max; check for above and left
			if i == rowlen-1 && j == collen-1 {
				if mat[i][j] < mat[i][j-1] && mat[i][j] < mat[i-1][j] {
					//sum += (mat[i][j] + 1)
					lowp := [2]int{i, j}
					lowpoints = append(lowpoints, lowp)
				}
			}
			// if row == 0 and anything other than 0th or max col
			if i == 0 && j < collen-1 && j != 0 {
				if mat[i][j] < mat[i][j-1] && mat[i][j] < mat[i+1][j] && mat[i][j] < mat[i][j+1] {
					//sum += (mat[i][j] + 1)
					lowp := [2]int{i, j}
					lowpoints = append(lowpoints, lowp)
				}
			}
			// if row == max and anything other than 0th or max col
			if i == rowlen-1 && j < collen-1 && j != 0 {
				if mat[i][j] < mat[i][j-1] && mat[i][j] < mat[i-1][j] && mat[i][j] < mat[i][j+1] {
					//sum += (mat[i][j] + 1)
					lowp := [2]int{i, j}
					lowpoints = append(lowpoints, lowp)
				}
			}

			// if any row with column == 0 (except 0 and max)
			if i < rowlen-1 && i != 0 && j == 0 {
				if mat[i][j] < mat[i][j+1] && mat[i][j] < mat[i+1][j] && mat[i][j] < mat[i-1][j] {
					//sum += (mat[i][j] + 1)
					lowp := [2]int{i, j}
					lowpoints = append(lowpoints, lowp)
				}
			}
			// if any row with column == max (except 0 and max rows)
			if i < rowlen-1 && i != 0 && j == collen-1 {
				if mat[i][j] < mat[i][j-1] && mat[i][j] < mat[i-1][j] && mat[i][j] < mat[i+1][j] {
					//sum += (mat[i][j] + 1)
					lowp := [2]int{i, j}
					lowpoints = append(lowpoints, lowp)
				}
			}

			// all other cases
			if i < rowlen-1 && i != 0 && j < collen-1 && j != 0 {
				if mat[i][j] < mat[i][j-1] && mat[i][j] < mat[i][j+1] && mat[i][j] < mat[i+1][j] && mat[i][j] < mat[i-1][j] {
					//sum += (mat[i][j] + 1)
					lowp := [2]int{i, j}
					lowpoints = append(lowpoints, lowp)
				}
			}

		}
	}
	return lowpoints
}

func BFSUntil9Reached(grid [][]int, startx int, starty int) int {

	// from the base point (startx, starty) - do BFS till we have visited all points
	// or reached a 9 or reached a row/col below or above the matrix ends.

	myQueue := Queue{}
	visited := make(map[string]bool)

	myQueue.Enqueue(strings.Join([]string{strconv.Itoa(startx), strconv.Itoa(starty)}, ","))
	cnt := 0
	for !myQueue.IsEmpty() {

		currval, _ := myQueue.Dequeue()
		_, found := visited[currval]
		if found {
			continue
		} // if found in the visited map, then ignore this and pick up the next item from queue
		visited[currval] = true

		currvals := strings.Split(currval, ",")
		x, _ := strconv.Atoi(currvals[0])
		y, _ := strconv.Atoi(currvals[1])
		cnt++
		if x+1 < len(grid) && grid[x+1][y] != 9 {
			myQueue.Enqueue(strings.Join([]string{strconv.Itoa(x + 1),
				strconv.Itoa(y)}, ","))
		}
		if x-1 >= 0 && grid[x-1][y] != 9 {
			myQueue.Enqueue(strings.Join([]string{strconv.Itoa(x - 1),
				strconv.Itoa(y)}, ","))
		}
		if y+1 < len(grid[0]) && grid[x][y+1] != 9 {
			myQueue.Enqueue(strings.Join([]string{strconv.Itoa(x),
				strconv.Itoa(y + 1)}, ","))
		}
		if y-1 >= 0 && grid[x][y-1] != 9 {
			myQueue.Enqueue(strings.Join([]string{strconv.Itoa(x),
				strconv.Itoa(y - 1)}, ","))
		}
	}

	return cnt
}
