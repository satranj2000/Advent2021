package main

import (
	"fmt"
	"sort"

	"sathish.com/adventcode/utils"
)

type nodeInserter struct {
	nodesInserts []*utils.Node
	insertvals   string
}

type CountNodes struct {
	value string
	cnt   int
}

func PuzzleDay14Part1(filename string, run int) int {

	template, moves := readfileday14(filename)

	head := createList(template)

	for i := 0; i < run; i++ {
		InsertmovestoNodes(moves, head)
	}

	finalnodes := head.ToArray()
	cnts := giveCounts(finalnodes)

	max := cnts[0].cnt
	min := cnts[0].cnt

	for i := 0; i < len(cnts); i++ {
		if cnts[i].cnt > max {
			max = cnts[i].cnt
		}
		if cnts[i].cnt < min {
			min = cnts[i].cnt
		}
	}

	return max - min
}

func giveCounts(nodes []string) []CountNodes {

	if len(nodes) <= 0 {
		return nil
	}
	sort.Strings(nodes)
	cntnodes := make([]CountNodes, 0)

	runningcnt := 0
	currval := nodes[0]
	for i := 0; i < len(nodes); i++ {
		if nodes[i] == currval {
			runningcnt++
		} else {
			cnode := CountNodes{}
			cnode.value = currval
			cnode.cnt = runningcnt
			cntnodes = append(cntnodes, cnode)
			currval = nodes[i]
			runningcnt = 1
		}
	}
	cnode := CountNodes{}
	cnode.value = currval
	cnode.cnt = runningcnt
	cntnodes = append(cntnodes, cnode)

	return cntnodes
}

func createList(template string) *utils.Node {

	head := utils.SetupNode(string(template[0]))
	for _, v := range template[1:] {
		head.Add(string(v))
	}

	return head
}

func InsertmovestoNodes(moves [][]string, head *utils.Node) *utils.Node {

	listnodeInserters := make([]nodeInserter, 0)

	for i := 0; i < len(moves); i++ {
		val := moves[i][0]
		var lside, rside string
		fmt.Sscanf(val, "%1s%1s", &lside, &rside)
		nodeIns := nodeInserter{}
		nodeIns.nodesInserts = head.TwonodesFound(lside, rside)
		nodeIns.insertvals = moves[i][1]
		listnodeInserters = append(listnodeInserters, nodeIns)
	}

	for i := 0; i < len(listnodeInserters); i++ {
		lnodes := listnodeInserters[i].nodesInserts
		for j := 0; j < len(lnodes); j++ {
			lnodes[j].AddAfter(listnodeInserters[i].insertvals)
		}
	}

	return head
}
