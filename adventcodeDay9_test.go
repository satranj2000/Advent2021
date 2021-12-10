package main

import "testing"

func TestPuzzleDay9Part1(t *testing.T) {

	res := PuzzleDay9Part1("inputs/input9_sample1.txt")
	if res != 15 {
		t.Errorf("Expected %v, got %v", 15, res)
	}
}

func TestPuzzleDay9Part2(t *testing.T) {

	res := PuzzleDay9Part2("inputs/input9_sample1.txt")
	if res != 1134 {
		t.Errorf("Expected %v, got %v", 1134, res)
	}
}

func TestBFSUntil9(t *testing.T) {
	inputmat := [][]int{{1, 3, 3},
		{9, 4, 5},
		{1, 9, 9}}

	expected := 5

	if BFSUntil9Reached(inputmat, 0, 0) != expected {
		t.Errorf("Did not get expected value of 5")
	}
}

func TestBFSUntil9_2(t *testing.T) {
	inputmat := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	if BFSUntil9Reached(inputmat, 0, 0) != 3 {
		t.Errorf("Did not get expected value of 3")
	}
}

func TestBFSUntil9_3(t *testing.T) {
	inputmat := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	if BFSUntil9Reached(inputmat, 0, 9) != 9 {
		t.Errorf("Did not get expected value of 9")
	}
}

func TestBFSUntil9_4(t *testing.T) {
	inputmat := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	if BFSUntil9Reached(inputmat, 2, 2) != 14 {
		t.Errorf("Did not get expected value of 14")
	}
}

func TestFindlowpoints(t *testing.T) {
	inputmat := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}
	expected := [][2]int{{0, 0}, {1, 9}, {2, 2}, {4, 6}}

	outputs := findlowpoints(inputmat)

	for i := 0; i < len(outputs); i++ {
		if outputs[i][0] != expected[i][0] && outputs[i][1] != expected[i][1] {
			t.Errorf("Expected %v, %v, got %v,%v", expected[i][0], expected[i][0],
				outputs[i][0], outputs[i][0])
		}
	}
}
