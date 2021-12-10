package main

import (
	"strings"
	"testing"
)

func TestPuzzle15(t *testing.T) {
	expectedResult := 26
	res := Puzzle15("input8_sample1.txt")
	if res != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, res)
	}
}

func TestSortString(t *testing.T) {
	inputs := []string{"a",
		"afe",
		"zbeca",
		"axab"}
	expected := []string{"a",
		"aef",
		"abcez",
		"aabx"}

	for i, input := range inputs {
		ret := sortString(input)
		if ret != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], ret)
		}
	}

}

func TestFindDigitPattern(t *testing.T) {
	input := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab "
	inputs := strings.Split(input, " ")
	expected := []string{"cagedb", "ab", "gcdfa", "fbcad", "eafb", "cdfbe",
		"cdfgeb", "dab", "acedgfb", "cefabd"}
	for i := 0; i < 10; i++ {
		expected[i] = sortString(expected[i])
	}
	outputs := findDigitPattern(inputs)

	for i := 0; i < 10; i++ {
		if expected[i] != outputs[i] {
			t.Errorf("Expected %v, got %v", expected[i], outputs[i])
		}
	}
}

func TestIsCharPresent(t *testing.T) {
	inputs1 := []string{"abcedfgh",
		"ced",
		"acf"}
	inputs2 := []string{"cdh",
		"cb",
		"fc"}

	expected := []bool{true, false, true}

	for i := 0; i < len(inputs1); i++ {
		out := IsCharsPresent(inputs1[i], inputs2[i])
		if out != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], out)
		}
	}
}

func TestPuzzle16(t *testing.T) {

	res := Puzzle16("input8_sample1.txt")
	if res != 61229 {
		t.Errorf("Expected %v, got %v", 61229, res)
	}
}

func TestIntArrtoString(t *testing.T) {

	input := []int{8, 3, 9, 4}
	output := IntArrToString(input)
	if output != "8394" {
		t.Errorf("Expected %v, got %v", "8394", output)
	}
}
