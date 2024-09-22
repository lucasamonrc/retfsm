package fsm

import "strings"

type FSM struct {
	States      []*State
	Transitions []*Transition
}

func NewFSM() *FSM {
	return &FSM{States: make([]*State, 0), Transitions: make([]*Transition, 0)}
}

func (g *FSM) AddState(n *State) {
	g.States = append(g.States, n)
}

func (g *FSM) PopState() *State {
	n := g.States[len(g.States)-1]
	g.States = g.States[:len(g.States)-1]
	return n
}

func (g *FSM) AddTransition(e *Transition) {
	g.Transitions = append(g.Transitions, e)
}

func (g *FSM) PopTransition() *Transition {
	e := g.Transitions[len(g.Transitions)-1]
	g.Transitions = g.Transitions[:len(g.Transitions)-1]
	return e
}

func (g *FSM) String() string {
	var sb strings.Builder

	sb.WriteString("FSM\n")
	for _, n := range g.States {
		sb.WriteString("--" + n.String() + "\n")
	}

	return sb.String()
}
