package fsm

import (
	"fmt"
	"strings"
)

type State struct {
	Id  int
	In  []*Transition
	Out []*Transition
}

func NewState(id int) *State {
	return &State{Id: id, In: make([]*Transition, 0), Out: make([]*Transition, 0)}
}

func (n *State) AddIn(e *Transition) {
	n.In = append(n.In, e)
}

func (n *State) PopIn() *Transition {
	e := n.In[len(n.In)-1]
	n.In = n.In[:len(n.In)-1]
	return e
}

func (n *State) AddOut(e *Transition) {
	n.Out = append(n.Out, e)
}

func (n *State) PopOut() *Transition {
	e := n.Out[len(n.Out)-1]
	n.Out = n.Out[:len(n.Out)-1]
	return e
}

func (n *State) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("State q%v\n", n.Id))
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
