package models

import (
	"strings"
)

type Adjacency struct {
	To      *Node
	Through string
}

type Graph struct {
	Nodes          []*Node
	Edges          []*Edge
	AdejacencyList map[*Node]Adjacency
}

func NewGraph() *Graph {
	return &Graph{Nodes: make([]*Node, 0), Edges: make([]*Edge, 0), AdejacencyList: make(map[*Node]Adjacency)}
}

func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) PopNode() *Node {
	n := g.Nodes[len(g.Nodes)-1]
	g.Nodes = g.Nodes[:len(g.Nodes)-1]
	return n
}

func (g *Graph) AddEdge(e *Edge) {
	g.Edges = append(g.Edges, e)
	g.AdejacencyList[e.From] = Adjacency{To: e.To, Through: e.Label}
}

func (g *Graph) PopEdge() *Edge {
	e := g.Edges[len(g.Edges)-1]
	g.Edges = g.Edges[:len(g.Edges)-1]
	return e
}

func (g *Graph) String() string {
	var sb strings.Builder

	sb.WriteString("Graph\n")
	for _, n := range g.Nodes {
		sb.WriteString("--" + n.String() + "\n")
	}

	return sb.String()
}
