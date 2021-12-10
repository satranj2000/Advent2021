package main

import "testing"

func TestPuzzleDay9Part1(t *testing.T) {

	res := PuzzleDay9Part1("inputs/input9_sample1.txt")
	if res != 15 {
		t.Errorf("Expected %v, got %v", 15, res)
	}
}
