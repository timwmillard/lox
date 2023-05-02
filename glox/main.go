package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		runPrompt()
	}
	if len(os.Args) == 2 {
		runFile(os.Args[1])
	}
	if len(os.Args) > 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [script]\n", os.Args[0])
		os.Exit(64)
	}
}

func runFile(filename string) {
	src, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read file %s", filename)
		os.Exit(65)
	}
	run(string(src))
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		run(scanner.Text())
		fmt.Print("> ")
	}
}

