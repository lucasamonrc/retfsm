package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/goccy/go-graphviz"
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
		line = strings.TrimSpace(line)

		l := lexer.NewLexer(line)
		p := parser.NewParser(l)

		machine := p.Parse()

		output := machine.ToBytes(graphviz.XDOT)
		io.WriteString(out, output.String())
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

		output := machine.ToBytes(graphviz.XDOT)
		io.WriteString(out, output.String())
		io.WriteString(out, "\n")
	}
}
