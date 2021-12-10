package main

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
