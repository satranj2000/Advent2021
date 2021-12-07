package main

import "testing"

func TestReadinputfile6(t *testing.T) {
	expectedoutputs := []int{3, 4, 3, 1, 2}

	output := readInput6File("input6_sample1.text")

	for i := 0; i < len(expectedoutputs); i++ {
		if output[i] != expectedoutputs[i] {
			t.Errorf("Expected %d, got %d", expectedoutputs[i], output[i])
		}
	}
}

func TestPuzzle11(t *testing.T) {
	res := Puzzle11("input6_sample1.txt", 18)
	if res != 26 {
		t.Errorf("Expected value of 26, got %d", res)
	}
}

func TestPuzzle11_2(t *testing.T) {
	res := Puzzle11("input6_sample1.txt", 80)
	if res != 5934 {
		t.Errorf("Expected value of 5934, got %d", res)
	}
}
