package main

import "testing"

func TestReadfile(t *testing.T) {
	expectedoutputs := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	output := readfile("input7_sample1.txt")

	for i := 0; i < len(expectedoutputs); i++ {
		if output[i] != expectedoutputs[i] {
			t.Errorf("Expected %d, got %d", expectedoutputs[i], output[i])
		}
	}
}
