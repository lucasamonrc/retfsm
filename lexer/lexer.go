package lexer

import (
	"unicode"

	symbol "github.com/lucasamonrc/regex-to-fsa/symbol"
)

type Lexer struct {
	input   string
	pos     int
	readPos int
	ch      byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() symbol.Symbol {
	var sym symbol.Symbol

	switch l.ch {
	case '*':
		if l.peekChar() == '*' {
			sym = newSymbol(symbol.ILLEGAL, l.ch)
		} else {
			sym = newSymbol(symbol.KSTAR, l.ch)
		}
	case 0:
		sym.Literal = ""
		sym.Type = symbol.EOF
	default:
		if isAlphaNumeric(l.ch) {
			sym = newSymbol(symbol.LITERAL, l.ch)
		} else {
			sym = newSymbol(symbol.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return sym
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos++
}

func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPos]
	}
}

func isAlphaNumeric(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || unicode.IsDigit(rune(ch))
}

func newSymbol(symType symbol.SymbolType, ch byte) symbol.Symbol {
	return symbol.Symbol{Type: symType, Literal: string(ch)}
}
