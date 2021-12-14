package utils

//defines one vertex and its edges/adjacency list or nodes that are linked to this node
type GraphNode struct {
	Val   string
	edges []string
}

// group of Graph Nodes make a graph
type Graph struct {
	nodes map[string]*GraphNode
}

func (g *Graph) Count() int {
	return len(g.nodes)
}

func (g *Graph) AddNode(graph *GraphNode) {
	if g.nodes == nil {
		g.nodes = make(map[string]*GraphNode)
	}
	if !g.IsNodePresent(graph) {
		g.nodes[graph.Val] = graph
	}
}

func (gnode *GraphNode) AddEdge(e string) bool {
	if !gnode.IsEdgePresent(e) {
		gnode.edges = append(gnode.edges, e)
		return true
	}
	return false
}

func (g *Graph) IsNodePresent(graph *GraphNode) bool {

	if _, ok := g.nodes[graph.Val]; ok {
		return true
	}
	return false
}

func (g *Graph) IsNodeValuePresent(s string) (bool, *GraphNode) {

	if _, ok := g.nodes[s]; ok {
		return true, g.nodes[s]
	}
	return false, nil
}

func (gnode *GraphNode) IsEdgePresent(e string) bool {

	for i := 0; i < len(gnode.edges); i++ {
		if gnode.edges[i] == e {
			return true
		}
	}
	return false
}

func (g *Graph) GetGraphNodes() []string {

	list := make([]string, len(g.nodes))

	i := 0
	for k, _ := range g.nodes {
		list[i] = k
		i++
	}

	return list

}
