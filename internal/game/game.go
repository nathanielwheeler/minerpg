package game

import "fmt"

// Game is the overall structure, holding entities, locations, methods, etc.
type Game struct{}

// NewGame will construct a game
func NewGame() *Game {
	g := Game{}

	fmt.Println("You are in a glistening mine.  There is an exit to the North.")

	return &g
}

// Execute takes in a command, passes it to internal logic, and returns whatever the response is.
func (g *Game) Execute(command string) (string, error) {
  return fmt.Sprintf("You just input: '%s'", command), nil
}