package main

import (
	"strings"
	"testing"
)

func TestCreateList(t *testing.T) {

	head := createList("ABCD")

	expected := []string{"A", "B", "C", "D"}
	arr := head.ToArray()

	if strings.Join(expected, ",") != strings.Join(arr, ",") {
		t.Errorf("Expected %v, got %v", expected, arr)
	}

}

func TestInsertMoves(t *testing.T) {
	head := createList("ABCD")
	moves := [][]string{
		{"AB", "X"},
		{"CD", "Y"},
	}

	expectedArr := []string{"A", "X", "B", "C", "Y", "D"}

	out := InsertmovestoNodes(moves, head)
	outarr := out.ToArray()

	for i := 0; i < len(expectedArr); i++ {
		if expectedArr[i] != outarr[i] {
			t.Errorf("Expected %v, got %v", expectedArr[i], outarr[i])
		}
	}

}

func TestInsertMoves2(t *testing.T) {
	head := createList("ABCD")
	moves := [][]string{
		{"AB", "X"},
		{"CD", "Y"},
		{"BC", "A"},
	}

	expectedArr := []string{"A", "X", "B", "A", "C", "Y", "D"}

	out := InsertmovestoNodes(moves, head)
	outarr := out.ToArray()

	for i := 0; i < len(expectedArr); i++ {
		if expectedArr[i] != outarr[i] {
			t.Errorf("Expected %v, got %v", expectedArr[i], outarr[i])
		}
	}

}

func TestInsertMoves3(t *testing.T) {
	head := createList("ABCDABCD")
	moves := [][]string{
		{"AB", "X"},
		{"CD", "Y"},
		{"BC", "A"},
		{"DA", "E"},
	}

	expectedArr := []string{"A", "X", "B", "A", "C", "Y", "D", "E", "A", "X", "B", "A", "C", "Y", "D"}

	out := InsertmovestoNodes(moves, head)
	outarr := out.ToArray()

	for i := 0; i < len(expectedArr); i++ {
		if expectedArr[i] != outarr[i] {
			t.Errorf("Expected %v, got %v", expectedArr[i], outarr[i])
		}
	}

}

func TestGiveCounts(t *testing.T) {

	input := []string{"C", "A", "A", "B", "B", "A"}

	out := giveCounts(input)

	if out[2].value != "C" && out[2].cnt != 1 {
		t.Error("Error")
	}
}

func TestPuzzleDay14Part1(t *testing.T) {

	cnt := PuzzleDay14Part1("inputs/input14_sample1.txt", 1)
	if cnt != 1 {
		t.Errorf("Expected 1, got %v", cnt)
	}
}

func TestPuzzleDay14Part1_2(t *testing.T) {

	cnt := PuzzleDay14Part1("inputs/input14_sample1.txt", 10)
	if cnt != 1588 {
		t.Errorf("Expected 1588, got %v", cnt)
	}
}

func TestPuzzleDay14Part1_3(t *testing.T) {

	cnt := PuzzleDay14Part1("inputs/input14_sample1.txt", 40)
	if cnt != 2188189693529 {
		t.Errorf("Expected 1588, got %v", cnt)
	}
}
