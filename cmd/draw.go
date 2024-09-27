package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/lucasamonrc/retfsm/lexer"
	"github.com/lucasamonrc/retfsm/parser"
)

func RunDraw() {
	args := os.Args[2:]

	if len(args) == 0 && !isInputRedirected() {
		fmt.Fprintln(os.Stderr, "\033[31mError: no input provided for draw command\033[0m")
		os.Exit(1)
	}

	if args[0] == "help" {
		fmt.Println("Usage: retfsm draw <input>")
		fmt.Println("       retfsm draw <input> <output>")
		os.Exit(0)
	}

	output := len(args) == 2

	var input string

	if isInputRedirected() {
		stdInput, err := io.ReadAll(os.Stdin)

		if err != nil {
			fmt.Fprintln(os.Stderr, "\033[31mError: could not read input from stdin\033[0m")
			fmt.Fprintf(os.Stderr, "\033[31m%v\033[0m\n", err)
			os.Exit(1)
		}

		input = string(stdInput)
	} else if len(args) > 0 {
		input = args[0]

		if isFilePath(input) {
			fileContent, err := os.ReadFile(input)

			if err != nil {
				fmt.Fprintln(os.Stderr, "\033[31mError: could not read file\033[0m")
				fmt.Fprintf(os.Stderr, "\033[31m%v\033[0m\n", err)
				os.Exit(1)
			}

			input = string(fileContent)
		} else {
			input = strings.Trim(input, `"'`)
		}
	}

	l := lexer.NewLexer(input)
	p := parser.NewParser(l)

	machine := p.Parse()
	dot := strings.TrimLeft(machine.ToDOT(), "\n")

	outputFile := "a.dot"

	if output {
		outputFile = args[1]
	}

	err := os.WriteFile(outputFile, []byte(dot), 0644)

	if err != nil {
		fmt.Fprintln(os.Stderr, "\033[31mError: could not write output file\033[0m")
		fmt.Fprintf(os.Stderr, "\033[31m%v\033[0m\n", err)
		os.Exit(1)
	}
}

func isInputRedirected() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error checking stdin:", err)
		os.Exit(1)
	}
	return (info.Mode() & os.ModeCharDevice) == 0
}

func isFilePath(input string) bool {
	_, err := os.Stat(input)
	return !os.IsNotExist(err)
}
