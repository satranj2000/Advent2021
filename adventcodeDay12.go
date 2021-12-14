package main

import (
	"fmt"
	"strings"

	"sathish.com/adventcode/utils"
)

func PuzzleDay12Part1(filename string) int {
	links := readfile2(filename)
	g := GenerateGraphfromfile(links)

	fmt.Println(g.Count())

	return g.Count()
}

func GenerateGraphfromfile(links []string) utils.Graph {

	g := utils.Graph{}

	for _, s := range links {
		sides := strings.Split(s, "-")

		if ok, gn1 := g.IsNodeValuePresent(sides[0]); ok {
			gn1.AddEdge(sides[1])
			g.AddNode(gn1)
		} else {
			gn := &utils.GraphNode{Val: sides[0]}
			gn.AddEdge(sides[1])
			g.AddNode(gn)
		}

		if ok, gn2 := g.IsNodeValuePresent(sides[1]); ok {
			gn2.AddEdge(sides[0])
			g.AddNode(gn2)
		} else {
			gn := &utils.GraphNode{Val: sides[1]}
			gn.AddEdge(sides[0])
			g.AddNode(gn)
		}

	}

	return g
}

