package main

import "testing"

func TestPuzzle13(t *testing.T) {

	results := []int{37}

	for i := 0; i < len(results); i++ {
		res := Puzzle13("input7_sample1.txt")
		if res != results[i] {
			t.Errorf("Expected value of %d, got %d", results[i], res)
		}
	}
}

func TestPuzzle14(t *testing.T) {

	results := []int{168}

	for i := 0; i < len(results); i++ {
		res := Puzzle14("input7_sample1.txt")
		if res != results[i] {
			t.Errorf("Expected value of %d, got %d", results[i], res)
		}
	}
}
