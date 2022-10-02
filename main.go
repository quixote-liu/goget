package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("missing URL\nTry 'gget --help` for more options")
		return
	}

	// TODO: optimize
}
