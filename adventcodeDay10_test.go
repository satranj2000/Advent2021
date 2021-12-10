package main

import "testing"

func TestCheckIsCorrupted(t *testing.T) {

	inputs := []string{"{([(<{}[<>[]}>{[]{[(<()>",
		"[<(<(<(<{}))><([]([]()",
		"[[[]]",
		"[[](){{{}}}]{"}

	outputs := []string{"}", ")", "", ""}
	outputs2 := []bool{true, true, false, false}

	for i, input := range inputs {
		out1, out2 := CheckIsCorrupted(input)
		if out1 != outputs[i] || out2 != outputs2[i] {
			t.Errorf("Expected %v, got %v", outputs[i], out1)
		}
	}

}

func TestPuzzleDay10_Part1(t *testing.T) {

	if PuzzleDay10_Part1("inputs/input10_sample1.txt") != 26397 {
		t.Errorf("Expected 26397, got something else ")
	}
}
