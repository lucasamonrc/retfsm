package main

import (
	"fmt"
	"os"

	"github.com/lucasamonrc/regex-to-fsm/cmd"
	"github.com/lucasamonrc/regex-to-fsm/repl"
)

func main() {
	if len(os.Args) < 2 {
		repl.Start(os.Stdin, os.Stdout)
	}

	command := os.Args[1]

	switch command {
	case cmd.Draw:
		cmd.RunDraw()
	default:
		fmt.Fprintf(os.Stderr, "\033[31mError: unknown command \033[4m%v\033[0m\033[0m\n", command)
		cmd.Usage()
		os.Exit(1)
	}
}
