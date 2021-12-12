package main

import "testing"

func TestDecorateMatrix(t *testing.T) {
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := [][]int{{-10, -10, -10, -10, -10}, {-10, 1, 2, 3, -10}, {-10, 4, 5, 6, -10}, {-10, 7, 8, 9, -10}, {-10, -10, -10, -10, -10}}

	output := decorateMatrix(input)

	if !compareMatrix(expected, output) {
		t.Error("The output matrix is not equal to expected matrix")
	}

}

func TestIncrementbyOne(t *testing.T) {
	input := [][]int{{-10, -10, -10, -10, -10}, {-10, 1, 2, 3, -10}, {-10, 4, 5, 6, -10}, {-10, 7, 8, 9, -10}, {-10, -10, -10, -10, -10}}

	expected := [][]int{{-10, -10, -10, -10, -10}, {-10, 2, 3, 4, -10},
		{-10, 5, 6, 7, -10}, {-10, 8, 9, 10, -10}, {-10, -10, -10, -10, -10}}

	incrementbyOne(input)

	if !compareMatrix(input, expected) {
		t.Error("input and expected are different")
	}
}

func TestPuzzleDay11Part1(t *testing.T) {

	out := PuzzleDay11Part1("inputs/input11_sample1.txt", 10)
	if out != 204 {
		t.Errorf("Expected number of flashes to be 204. got %v", out)
	}

}

func TestPuzzleDay11Part1_100(t *testing.T) {
	out := PuzzleDay11Part1("inputs/input11_sample1.txt", 100)
	if out != 1656 {
		t.Errorf("Expected number of flashes to be 1656. got %v", out)
	}
}

func TestPuzzleDay11Part2(t *testing.T) {
	out := PuzzleDay11Part2("inputs/input11_sample1.txt")
	if out != 195 {
		t.Errorf("Expected number of runs when flashes are all lit is to be 195. But, got %v", out)
	}
}

func TestFlash(t *testing.T) {
	input := [][]int{{1, 1, 1, 1, 1}, {1, 9, 9, 9, 1}, {1, 9, 1, 9, 1}, {1, 9, 9, 9, 1}, {1, 1, 1, 1, 1}}

	expected := [][]int{{-10, -10, -10, -10, -10, -10, -10},
		{-10, 3, 4, 5, 4, 3, -10}, {-10, 4, 0, 0, 0, 4, -10},
		{-10, 5, 0, 0, 0, 5, -10},
		{-10, 4, 0, 0, 0, 4, -10},
		{-10, 3, 4, 5, 4, 3, -10},
		{-10, -10, -10, -10, -10, -10, -10}}

	input2 := decorateMatrix(input)
	incrementbyOne(input2)
	flash(input2)
	resetMatrix(input2)

	if !compareMatrix(input2, expected) {
		t.Error("output is not same as expected")
	}
}

func TestFlash1(t *testing.T) {
	input := [][]int{{1, 1, 1, 1, 1}, {1, 9, 9, 9, 1}, {1, 9, 1, 9, 1}, {1, 9, 9, 9, 1}, {1, 1, 1, 1, 1}}

	expected := [][]int{{-10, -10, -10, -10, -10, -10, -10},
		{-10, 4, 5, 6, 5, 4, -10},
		{-10, 5, 1, 1, 1, 5, -10},
		{-10, 6, 1, 1, 1, 6, -10},
		{-10, 5, 1, 1, 1, 5, -10},
		{-10, 4, 5, 6, 5, 4, -10},
		{-10, -10, -10, -10, -10, -10, -10}}

	input2 := decorateMatrix(input)
	incrementbyOne(input2)
	// call flash twice
	flash(input2)
	resetMatrix(input2)
	incrementbyOne(input2)
	flash(input2)

	if !compareMatrix(input2, expected) {
		t.Error("output is not same as expected")
	}
}

func compareMatrix(lside [][]int, rside [][]int) bool {
	if len(lside) != len(rside) {
		return false
	}

	for i := 0; i < len(lside); i++ {
		for j := 0; j < len(lside[i]); j++ {
			if lside[i][j] != rside[i][j] {
				return false
			}
		}
	}

	return true
}
