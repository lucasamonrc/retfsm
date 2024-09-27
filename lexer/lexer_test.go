package lexer

import (
	"testing"

	"github.com/lucasamonrc/retfsm/symbol"
)

func TestNextSymbol(t *testing.T) {
	input := "ab*c"

	tests := []struct {
		expectedType    symbol.SymbolType
		expectedLiteral string
	}{
		{symbol.LITERAL, "a"},
		{symbol.LITERAL, "b"},
		{symbol.KSTAR, "*"},
		{symbol.LITERAL, "c"},
		{symbol.EOF, ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		sym := l.NextSymbol()

		if sym.Type != tt.expectedType {
			t.Fatalf("tests[%d] - symboltype wrong. expected=%q, got=%q",
				i, tt.expectedType, sym.Type)
		}

		if sym.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, sym.Literal)
		}
	}
}
