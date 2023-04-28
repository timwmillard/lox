package main

import (
"os"
"fmt"
)

func main() {
	
	if len(os.Args) < 2 {
		runPrompt();
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
	fmt.Println("runFile", filename)
}

func runPrompt() {
	fmt.Println("runPrompt")
}


