package game

import (
	"fmt"
	"strings"
)

// Game is the overall structure, holding entities, locations, methods, etc.
type Game struct {
	CurrentRoomID uint
	Rooms         map[uint]Room
}

// NewGame will construct a game
func NewGame() *Game {
	g := Game{CurrentRoomID: uint(1)}
	g.GenerateRooms()

  fmt.Println(g.GetRoomDescription(g.CurrentRoomID))

	return &g
}

// Execute takes in a command, passes it to internal logic, and returns whatever the response is.
func (g *Game) Execute(command string) string {
	// TODO Put a proper validation layer in?
	if command == "" {
		return "Type something in, yo."
	}
	command = strings.ToLower(command) // normalize to lowercase

	fields := strings.Fields(command)
	switch fields[0] {
	case "go":
    if len(fields) < 2 {
      return "Go where?"
    }
    return g.GoDirection(fields[1])
  case "look":
    desc := g.GetRoomDescription(g.CurrentRoomID)
		return desc
	case "something":
		return "Don't you sass me."
	default:
		return fmt.Sprintf("'%s' does not compute.", command)
	}
}
