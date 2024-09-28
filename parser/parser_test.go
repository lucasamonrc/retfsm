package parser

import (
	"testing"

	"github.com/lucasamonrc/retfsm/fsm"
	"github.com/lucasamonrc/retfsm/lexer"
)

func TestParse(t *testing.T) {
	input := "ab*c"
	l := lexer.NewLexer(input)
	p := NewParser(l)

	expectedMachine := prepareMachine()
	actualMachine := p.Parse()

	if actualMachine.String() != expectedMachine.String() {
		t.Fatalf("expected machine %v, got %v", expectedMachine, actualMachine)
	}

}

func prepareMachine() *fsm.FSM {
	expectedMachine := fsm.NewFSM()

	q0 := fsm.NewState(0)
	q1 := fsm.NewState(1)
	q2 := fsm.NewState(2)

	tA := fsm.NewTransition("a", q0, q1)
	tB := fsm.NewTransition("b", q1, q1)
	tC := fsm.NewTransition("c", q1, q2)

	q0.AddOut(tA)
	q1.AddOut(tB)
	q1.AddOut(tC)

	q1.AddIn(tA)
	q1.AddIn(tB)
	q2.AddIn(tC)

	expectedMachine.AddState(q0)
	expectedMachine.AddState(q1)
	expectedMachine.AddState(q2)

	expectedMachine.AddTransition(tA)
	expectedMachine.AddTransition(tB)
	expectedMachine.AddTransition(tC)

	return expectedMachine
}
