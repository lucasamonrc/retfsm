package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

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
		io.WriteString(out, machine.String()+"\n")
	}
}
