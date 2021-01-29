package game

import (
	"testing"

	"github.com/matryer/is"
)

func TestNewGame(t *testing.T) {
	is := is.New(t)

	g := NewGame()

	is.Equal(g.CurrentRoomID, uint(1))
	testRoom := g.Rooms[0]
	testRoomAssertion := Room{
		Name:        "test room",
		Description: "A bland, featureless room.  There is a DOOR labelled 'door that connects to this very room.'",
		Doors:       []Door{Door{RoomID: 0, Direction: "door"}},
	}
  is.Equal(testRoom, testRoomAssertion)
}

func TestExecute(t *testing.T) {
	is := is.New(t)

	g := NewGame()

	res := g.Execute("look")
	is.Equal(res, "You are in a glistening mine.  There is a passage to the NORTH and door to the EAST.")
  
  // test invalid directions
	res = g.Execute("go")
	is.Equal(res, "Go where?")
	res = g.Execute("go south")
	is.Equal(res, "You can't go that way.")
	res = g.Execute("south")
	is.Equal(res, "'south' does not compute.")
  res = g.Execute("east")
  is.Equal(res, "The way EAST is locked.")

  // test going north, getting key, unlocking door, and going east.
  res = g.Execute("go north")
  is.Equal(res, "You are in a dusty break room containing an ancient card TABLE.  There is an exit to the SOUTH.")
  // get key
  res = g.Execute("south")
  is.Equal(res, "You are in a glistening mine.  There is a passage to the NORTH and door to the EAST.")
  // unlock door
  res = g.Execute("east")
  is.Equal(res, "You come to a deep mineshaft.  You can go back WEST or go DOWN.")
}
