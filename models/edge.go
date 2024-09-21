package models

import "fmt"

type Edge struct {
	Label string
	To    *Node
	From  *Node
}

func NewEdge(label string, from *Node, to *Node) *Edge {
	return &Edge{Label: label, To: to, From: from}
}

func (e *Edge) String() string {
	return fmt.Sprintf("q%v --%s--> q%v", e.From.Id, e.Label, e.To.Id)
}
