package fsm

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/goccy/go-graphviz"
	"github.com/lucasamonrc/retfsm/util"
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

func (m *FSM) ToBytes(outType graphviz.Format) bytes.Buffer {
	g := graphviz.New()
	graph, err := g.Graph()

	defer func() {
		if err := graph.Close(); err != nil {
			util.LogError("could not close graph", err)
		}
		g.Close()
	}()

	if err != nil {
		util.LogError("could not create graph", err)
	}

	for _, transition := range m.Transitions {
		from, err := graph.Node(fmt.Sprintf("q%v", transition.From.Id))

		if err != nil {
			util.LogError("could not find node", err)
		}

		if from == nil {
			from, err = graph.CreateNode(fmt.Sprintf("q%v", transition.From.Id))
			if err != nil {
				util.LogError("could not create node", err)
			}
			from.SetLabel(fmt.Sprintf("q%v", transition.From.Id))
			from.SetShape("circle")
		}

		to, err := graph.Node(fmt.Sprintf("q%v", transition.To.Id))

		if err != nil {
			util.LogError("could not find node", err)
		}

		if to == nil {
			to, err = graph.CreateNode(fmt.Sprintf("q%v", transition.To.Id))
			if err != nil {
				util.LogError("could not create node", err)
			}
			to.SetLabel(fmt.Sprintf("q%v", transition.To.Id))

			if transition.To.Id == m.States[len(m.States)-1].Id {
				to.SetShape("doublecircle")
			} else {
				to.SetShape("circle")
			}
		}

		t, err := graph.CreateEdge(transition.Label, from, to)
		if err != nil {
			util.LogError("could not create edge", err)
		}
		t.SetLabel(transition.Label)
	}

	graph.SetRankDir("LR")
	graph.SetSize(8, 5)

	var buf bytes.Buffer
	g.Render(graph, outType, &buf)
	return buf
}
