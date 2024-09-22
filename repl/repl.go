package repl

import (
	"bufio"
	"fmt"
	"io"

	lexer "github.com/lucasamonrc/regex-to-fsa/lexer"
	symbol "github.com/lucasamonrc/regex-to-fsa/symbol"
)

const PROMPT = "regex> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)

		for tok := l.NextSymbol(); tok.Type != symbol.EOF; tok = l.NextSymbol() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
