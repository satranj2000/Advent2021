package main

import "testing"

func TestReadinputfile6(t *testing.T) {
	expectedoutputs := []int{3, 4, 3, 1, 2}

	output := readInput6File("input6_sample1.txt")

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

func TestPuzzle11_v2(t *testing.T) {
	inVals := []int{18, 80}
	results := []int{26, 5934}

	for i := 0; i < 2; i++ {
		res := Puzzle11_v2("input6_sample1.txt", inVals[i])
		if res != results[i] {
			t.Errorf("Expected value of %d, got %d", results[i], res)
		}
	}
}
