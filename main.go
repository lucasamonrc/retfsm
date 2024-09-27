package main

import (
	"fmt"
	"os"

	"github.com/lucasamonrc/retfsm/cmd"
	"github.com/lucasamonrc/retfsm/repl"
	"github.com/lucasamonrc/retfsm/util"
)

func main() {
	if len(os.Args) < 2 {
		repl.Start(os.Stdin, os.Stdout)
		os.Exit(0)
	}

	command := os.Args[1]

	switch command {
	case cmd.Draw:
		cmd.RunDraw()
	case cmd.Help:
		cmd.Usage()
	default:
		util.LogError(fmt.Sprintf("unknown command %v", command), nil)
		cmd.UsageFromError()
		os.Exit(1)
	}
}
