package main

import (
	"os"

	"github.com/lucasamonrc/regex-to-fsa/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
