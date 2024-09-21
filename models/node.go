package models

import (
	"fmt"
	"strings"
)

type Node struct {
	Id  int
	In  []*Edge
	Out []*Edge
}

func NewNode(id int) *Node {
	return &Node{Id: id, In: make([]*Edge, 0), Out: make([]*Edge, 0)}
}

func (n *Node) AddIn(e *Edge) {
	n.In = append(n.In, e)
}

func (n *Node) PopIn() *Edge {
	e := n.In[len(n.In)-1]
	n.In = n.In[:len(n.In)-1]
	return e
}

func (n *Node) AddOut(e *Edge) {
	n.Out = append(n.Out, e)
}

func (n *Node) PopOut() *Edge {
	e := n.Out[len(n.Out)-1]
	n.Out = n.Out[:len(n.Out)-1]
	return e
}

func (n *Node) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Node q%v\n", n.Id))
	sb.WriteString("Parents:\n")
	for _, e := range n.In {
		sb.WriteString(fmt.Sprintf("  q%v --%s-->\n", e.From.Id, e.Label))
	}

	sb.WriteString("Children:\n")
	for _, e := range n.Out {
		sb.WriteString(fmt.Sprintf("  --%s--> q%v\n", e.Label, e.To.Id))
	}

	return sb.String()
}
