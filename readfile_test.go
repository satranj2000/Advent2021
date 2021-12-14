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

func TestReadfileDay13(t *testing.T) {

	expectedArr := [][]int{
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

	expectedStr := []string{
		"y,7", "x,5",
	}

	out1, out2, outx, outy := readfileday13("inputs/input13_sample1.txt")

	if !compareMatrix(expectedArr, out1) {
		t.Error("return array of ints did not match")
	}

	for i := 0; i < 2; i++ {
		if out2[i] != expectedStr[i] {
			t.Errorf("wrong values ")
		}
	}

	if outx != 10 {
		t.Errorf("Expected 10 got %v as max x value", outx)
	}

	if outy != 14 {
		t.Errorf("Expected 14 got %v as max y value", outy)
	}
}
