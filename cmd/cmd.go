package cmd

import (
	"fmt"
	"os"
)

const (
	Draw = "draw"
	Help = "help"
)

func UsageFromError() {
	fmt.Fprintln(os.Stderr, "\033[31mUsage:\033[0m")
	fmt.Fprintln(os.Stderr, "  \033[31m- retfsm\033[0m")
	fmt.Fprintln(os.Stderr, "  \033[31m- retfsm help\033[0m")
	fmt.Fprintln(os.Stderr, "  \033[31m- retfsm draw <input>\033[0m")
}

func Usage() {
	fmt.Println("Usage:")
	fmt.Println("  - retfsm")
	fmt.Println("  - retfsm help")
	fmt.Println("  - retfsm draw <input>")
	os.Exit(0)
}
