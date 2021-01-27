package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
    os.Exit(1)
	}
}

// glistening-mine : 
func run() error {
	fmt.Println("glistening-mine!")
	return nil
}
