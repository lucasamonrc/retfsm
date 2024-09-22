package fsm

import "fmt"

type Transition struct {
	Label string
	To    *State
	From  *State
}

func NewTransition(label string, from *State, to *State) *Transition {
	return &Transition{Label: label, To: to, From: from}
}

func (e *Transition) String() string {
	return fmt.Sprintf("q%v --%s--> q%v", e.From.Id, e.Label, e.To.Id)
}
