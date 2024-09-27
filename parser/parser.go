package parser

import (
	"errors"
	"os"

	"github.com/lucasamonrc/retfsm/fsm"
	"github.com/lucasamonrc/retfsm/lexer"
	"github.com/lucasamonrc/retfsm/symbol"
	"github.com/lucasamonrc/retfsm/util"
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	currentSymbol  symbol.Symbol
	previousSymbol symbol.Symbol
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{l: l, errors: []string{}}

}

func (p *Parser) nextSymbol() {
	p.previousSymbol = p.currentSymbol
	p.currentSymbol = p.l.NextSymbol()
}

func (p *Parser) currentSymbolIs(t symbol.SymbolType) bool {
	return p.currentSymbol.Type == t
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) Parse() *fsm.FSM {
	machine := fsm.NewFSM()
	stateId := 0

	initial := fsm.NewState(stateId)
	machine.AddState(initial)

	var tmp *fsm.State
	var prev *fsm.State

	current := initial

	for p.nextSymbol(); !p.currentSymbolIs(symbol.EOF); p.nextSymbol() {
		if p.currentSymbolIs(symbol.ILLEGAL) {
			err := errors.New(p.currentSymbol.Literal)
			util.LogError("illegal symbol", err)
			os.Exit(1)
		}

		if p.currentSymbolIs(symbol.KSTAR) {
			tmp = nil
			current = prev

			machine.PopTransition()
			machine.PopState()
			current.PopOut()

			transition := fsm.NewTransition(p.previousSymbol.Literal, current, current)

			current.AddOut(transition)
			current.AddIn(transition)

			machine.AddTransition(transition)

			prev = nil
			continue
		}

		if tmp == nil {
			tmp = fsm.NewState(-1)

			transition := fsm.NewTransition(p.currentSymbol.Literal, current, tmp)

			current.AddOut(transition)
			tmp.AddIn(transition)

			machine.AddTransition(transition)
			machine.AddState(tmp)
		} else {
			stateId++
			current.Id = stateId
			tmp = fsm.NewState(-1)

			transition := fsm.NewTransition(string(p.currentSymbol.Literal), current, tmp)

			current.AddOut(transition)
			tmp.AddIn(transition)

			machine.AddTransition(transition)
			machine.AddState(tmp)
		}

		prev = current
		current = tmp
	}

	if p.currentSymbolIs(symbol.EOF) {
		stateId++
		current.Id = stateId
	}

	return machine
}
