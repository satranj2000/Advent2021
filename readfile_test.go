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
