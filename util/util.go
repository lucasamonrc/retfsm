package util

import (
	"fmt"
	"os"
)

func LogError(msg string, err error) {
	fmt.Fprintf(os.Stderr, "\033[31mError: %v\033[0m\n", msg)

	if err != nil {
		fmt.Fprintf(os.Stderr, "\033[31m%v\033[0m\n", err)
	}
}
