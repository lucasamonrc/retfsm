package main

import (
	"os"

	"github.com/lucasamonrc/regex-to-fsm/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
