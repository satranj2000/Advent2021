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

func TestPuzzleDay10_Part2(t *testing.T) {

	if PuzzleDay10Part2("inputs/input10_sample1.txt") != 288957 {
		t.Errorf("Expected 288957, got something else ")
	}
}

func TestFindCompletionCharsScore(t *testing.T) {
	input := "{[]<>(()"
	expected := ")}"
	_, out := findCompletionCharsScore(input)
	if out != expected {
		t.Errorf("Expected %v, got %v", expected, out)
	}
}

func TestFindCompletionCharsScore2(t *testing.T) {
	input := "<{([{{}}[<[[[<>{}]]]>[]]"
	expected := "])}>"
	expectedpoints := 294
	points, out := findCompletionCharsScore(input)
	if out != expected {
		t.Errorf("Expected %v, got %v", expected, out)
	}

	if points != expectedpoints {
		t.Errorf("Expected %v, got %v", expectedpoints, points)
	}
}

func TestFindCompletionCharsScore3(t *testing.T) {
	input := "[({(<(())[]>[[{[]{<()<>>"
	expected := "}}]])})]"
	expectedpoints := 288957
	points, out := findCompletionCharsScore(input)
	if out != expected {
		t.Errorf("Expected %v, got %v", expected, out)
	}

	if points != expectedpoints {
		t.Errorf("Expected %v, got %v", expectedpoints, points)
	}
}

func TestFindCompletionCharsScore4(t *testing.T) {
	inputs := []string{"[(()[<>])]({[<{<<[]>>(",
		"(((({<>}<{<{<>}{[]{[]{}",
		"{<[[]]>}<{[{[{[]{()[[[]"}
	expected := []string{")}>]})", "}}>}>))))", "]]}}]}]}>"}
	expectedpoints := []int{5566, 1480781, 995444}

	for i := 0; i < 3; i++ {
		points, out := findCompletionCharsScore(inputs[i])
		if out != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], out)
		}
		if points != expectedpoints[i] {
			t.Errorf("Expected %v, got %v", expectedpoints[i], points)
		}
	}

}
