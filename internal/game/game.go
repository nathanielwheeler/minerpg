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
func NewGame() (*Game, error) {
	g := Game{CurrentRoomID: uint(1)}
	g.GenerateRooms()

  fmt.Println(g.GetRoomDescription(g.CurrentRoomID))

	return &g, nil
}

// Execute takes in a command, passes it to internal logic, and returns whatever the response is.
func (g *Game) Execute(command string) (string, error) {
	// TODO Put a proper validation layer in?
	if command == "" {
		return "Type something in, yo.", nil
	}
	command = strings.ToLower(command) // normalize to lowercase

	// Turn command into fields
	fields := strings.Fields(command)
	// args := fields[:2]

	switch fields[0] {
	case "go":
		// TODO refactor for simplicity, this is unwieldy in the long run
    // I could use util.StringExistsInSlice on the location directions, keeping my switch statements a bit more maintainable
    if len(fields) < 2 {
      return "Go where?", nil
    }
    return g.GoDirection(fields[1]), nil
  case "look":
    desc := g.GetRoomDescription(g.CurrentRoomID)
		return desc, nil
	case "something":
		return "Don't you sass me.", nil
	default:
		return fmt.Sprintf("'%s' does not compute.", command), nil
	}
}
