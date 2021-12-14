package main

import (
	"sort"
	"testing"
)

func TestGenerateGraphfromfile(t *testing.T) {

	inputs := []string{"a-b", "b-c", "c-d"}

	gout := GenerateGraphfromfile(inputs)

	outnodes := []string{"a", "b", "c", "d"}
	if gout.Count() != 4 {
		t.Error("Expected 4 nodes ")
	}

	goutNodes := gout.GetGraphNodes()

	sort.Strings(goutNodes)

	for i := 0; i < len(outnodes); i++ {
		if goutNodes[i] != outnodes[i] {
			t.Errorf("Expected %v, got %v", outnodes[i], goutNodes[i])
		}
	}

}
