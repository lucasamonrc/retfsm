package fsm

import (
	"fmt"
	"strings"
)

type FSM struct {
	States      []*State
	Transitions []*Transition
}

func NewFSM() *FSM {
	return &FSM{States: make([]*State, 0), Transitions: make([]*Transition, 0)}
}

func (m *FSM) AddState(n *State) {
	m.States = append(m.States, n)
}

func (m *FSM) PopState() *State {
	n := m.States[len(m.States)-1]
	m.States = m.States[:len(m.States)-1]
	return n
}

func (m *FSM) AddTransition(e *Transition) {
	m.Transitions = append(m.Transitions, e)
}

func (m *FSM) PopTransition() *Transition {
	e := m.Transitions[len(m.Transitions)-1]
	m.Transitions = m.Transitions[:len(m.Transitions)-1]
	return e
}

func (m *FSM) String() string {
	var sb strings.Builder

	sb.WriteString("FSM\n")
	for _, n := range m.States {
		sb.WriteString("--" + n.String() + "\n")
	}

	return sb.String()
}

func (m *FSM) ToDOT() string {
	lastId := m.States[len(m.States)-1].Id

	dot := fmt.Sprintf(`
digraph finite_state_machine {
    rankdir=LR;
    size="8,5";

    node [shape = doublecircle]; q%v;
    node [shape = circle];

`, lastId)

	for _, state := range m.States {
		for _, transition := range state.Out {
			dot += fmt.Sprintf("    q%v -> q%v [ label = \"%s\" ];\n", transition.From.Id, transition.To.Id, transition.Label)
		}
	}

	dot += "}"

	return dot
}
