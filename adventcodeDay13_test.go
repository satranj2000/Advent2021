package main

import "testing"

func TestFilltransparentSheet(t *testing.T) {

	inputarr := [][]int{
		{6, 10},
		{0, 14},
		{9, 10},
		{0, 3},
		{10, 4},
		{4, 11},
		{6, 0},
		{6, 12},
		{4, 1},
		{0, 13},
		{10, 12},
		{3, 4},
		{3, 0},
		{8, 4},
		{1, 10},
		{2, 14},
		{8, 10},
		{9, 0},
	}

	maxx := 10
	maxy := 14

	trans := filltransparentSheet(inputarr, maxx, maxy)

	if trans[10][6] != 1 || trans[1][4] != 1 {
		t.Errorf("values not proper ")
	}
}

func TestFold1(t *testing.T) {

	input := [][]int{
		{1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 1},
	}

	expectedoutput := [][]int{
		{1, 1},
		{1, 1},
		{1, 1},
		{1, 0},
	}

	out := fold(input, "x", 2)

	if !compareMatrix(out, expectedoutput) {
		t.Errorf("Expected fold is wrong. It is showing as %v", out)
	}

}

func TestFold2(t *testing.T) {

	input := [][]int{
		{1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 1},
	}

	expectedoutput := [][]int{
		{1, 0, 0, 1, 1},
		{1, 1, 0, 1, 1},
	}

	out := fold(input, "y", 2)

	if !compareMatrix(out, expectedoutput) {
		t.Errorf("Expected fold is wrong. It is showing as %v", out)
	}

}

func TestFold3(t *testing.T) {

	input := [][]int{
		{1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 1},
		{1, 1, 0, 0, 0},
	}

	expectedoutput := [][]int{
		{1, 1, 0, 0, 0},
		{1, 0, 0, 1, 1},
		{1, 1, 0, 1, 1},
	}

	out := fold(input, "y", 2)

	if !compareMatrix(out, expectedoutput) {
		t.Errorf("Expected fold is wrong. It is showing as %v", out)
	}

}

func TestFold4(t *testing.T) {

	input := [][]int{
		{1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 0, 0, 0},
	}

	expectedoutput := [][]int{
		{1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 1, 1},
	}

	out := fold(input, "y", 3)

	if !compareMatrix(out, expectedoutput) {
		t.Errorf("Expected fold is wrong. It is showing as %v", out)
	}

}

func TestFold5(t *testing.T) {

	input := [][]int{
		{1, 0, 0, 1, 0, 0},
		{1, 1, 0, 0, 1, 1},
		{0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 1, 1},
	}

	expectedoutput := [][]int{
		{0, 1, 1},
		{1, 1, 1},
		{0, 1, 1},
		{1, 1, 0},
	}

	out := fold(input, "x", 2)

	if !compareMatrix(out, expectedoutput) {
		t.Errorf("Expected fold is wrong. It is showing as %v", out)
	}

}

func TestFold6(t *testing.T) {

	input := [][]int{
		{1, 1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0, 1},
		{0, 0, 0, 0, 1, 1},
		{1, 0, 0, 0, 0, 1},
	}

	expectedoutput := [][]int{
		{1, 1, 1},
		{0, 1, 1},
		{0, 1, 1},
		{1, 1, 0},
	}

	out := fold(input, "x", 3)

	if !compareMatrix(out, expectedoutput) {
		t.Errorf("Expected fold is wrong. It is showing as %v", out)
	}

}

func TestPuzzleDay13Part1(t *testing.T) {

	out := PuzzleDay13Part1("inputs/input13_sample1.txt", 1)

	if out != 17 {
		t.Errorf("Expected 17 got %v", out)
	}
}

func TestAddRowsBottom(t *testing.T) {

	mat := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	expectedout := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {0, 0, 0, 0}, {0, 0, 0, 0}}

	outmat := addRowsBottom(mat, 2)

	if !compareMatrix(outmat, expectedout) {
		t.Errorf("matrix values are not same. outmatrix is %v", outmat)
	}
}

func TestAddRowsTop(t *testing.T) {

	mat := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	expectedout := [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {1, 2, 3, 4}, {5, 6, 7, 8}}

	outmat := addRowsTop(mat, 2)

	if !compareMatrix(outmat, expectedout) {
		t.Errorf("matrix values are not same. outmatrix is %v", outmat)
	}
}

func TestAddColsBegin(t *testing.T) {

	mat := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	expectedout := [][]int{{0, 0, 1, 2, 3, 4}, {0, 0, 5, 6, 7, 8}}

	outmat := addColsBegin(mat, 2)

	if !compareMatrix(outmat, expectedout) {
		t.Errorf("matrix values are not same. outmatrix is %v", outmat)
	}

}

func TestAddColsEnd(t *testing.T) {

	mat := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	expectedout := [][]int{{1, 2, 3, 4, 0, 0}, {5, 6, 7, 8, 0, 0}}

	outmat := addColsEnd(mat, 2)

	if !compareMatrix(outmat, expectedout) {
		t.Errorf("matrix values are not same. outmatrix is %v", outmat)
	}

}
