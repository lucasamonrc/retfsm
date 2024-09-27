package cmd

import (
	"fmt"
	"os"
)

const (
	Draw = "draw"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "\033[31mUsage:\033[0m")
	fmt.Fprintln(os.Stderr, "  \033[31m- retfsm\033[0m")
	fmt.Fprintln(os.Stderr, "  \033[31m- retfsm draw <input>\033[0m")
}
