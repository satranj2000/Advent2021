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

func TestReadfileDay8(t *testing.T) {

	expected_lside := []string{"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb ",
		"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec "}

	expected_rside := []string{" fdgacbe cefdb cefbgd gcbe",
		" fcgedb cgb dgebacf gc"}

	lside, rside := readfileday8("input8_sample2.txt")

	for i := 0; i < len(expected_lside); i++ {
		if expected_lside[i] != lside[i] {
			t.Errorf("Expected left side %s, got %s", expected_lside[i], lside[i])
		}
		if expected_rside[i] != rside[i] {
			t.Errorf("Expected right side %s, got %s", expected_rside[i], rside[i])
		}
	}

}

func TestReadfileDay9(t *testing.T) {

	expected := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}
	mat := readfileday9("inputs/input9_sample1.txt")

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] != expected[i][j] {
				t.Errorf("Expected %v, got %v", mat[i], expected[i])
			}
		}
	}

}
