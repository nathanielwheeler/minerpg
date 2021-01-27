package main

import (
	"errors"
	"fmt"
	"os"
  "runtime"
  
  "minerpg/internal/util"
)

// minerpg will run a cui that will start the game loop.  Only supports linux and darwin.
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	if !util.IsOSSupported(runtime.GOOS) {
		return errors.New("operating system not supported")
	}

	return nil
}
