package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	// "minerpg/internal/game"
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

  // Start game loop
  
	util.ClearConsole()
  fmt.Println("\tMineRPG!")
  reader := bufio.NewReader(os.Stdin)

	if g, err := game.NewGame(); err != nil {
	  return err
  }
  
	// Take input and await response
	for {
		fmt.Print("==> ")
		req, _ := reader.ReadString('\n')
    req = strings.Replace(req, "\n", "", -1) // Convert CRLF to LF
    
		res, err := g.Execute(req)
		if err != nil {
			return err
    }

		util.ClearConsole()
		fmt.Printf("%s\n", res)
	}
}