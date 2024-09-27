package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/lucasamonrc/retfsm/lexer"
	"github.com/lucasamonrc/retfsm/parser"
)

const PROMPT = "retfsm> "

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
		p := parser.NewParser(l)

		machine := p.Parse()

		io.WriteString(out, machine.String())
		io.WriteString(out, "\n")
		io.WriteString(out, machine.ToDOT())
		io.WriteString(out, "\n")
	}
}

func Debug(in io.Reader, out io.Writer) {
	line := "ab*c"

	for {
		fmt.Fprint(out, PROMPT)

		l := lexer.NewLexer(line)
		p := parser.NewParser(l)

		machine := p.Parse()

		io.WriteString(out, machine.String())
		io.WriteString(out, "\n")
		io.WriteString(out, machine.ToDOT())
		io.WriteString(out, "\n")
	}
}
