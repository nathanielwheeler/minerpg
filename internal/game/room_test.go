package game

import (
	"testing"

	"github.com/matryer/is"
)

func TestGenerateRooms(t *testing.T) {
	is := is.New(t)

  g := Game{}
  g.GenerateRooms()
  testRoom := g.Rooms[0]

  is.Equal(testRoom.Name, "test room")
  is.Equal(testRoom.Description, "A bland, featureless room.  There is a DOOR labelled 'door that connects to this very room.'")
}

func TestGetRoomDescription(t *testing.T) {
  is := is.New(t)

  g := NewGame()
  desc := g.GetRoomDescription(0)
  is.Equal(desc, "A bland, featureless room.  There is a DOOR labelled 'door that connects to this very room.'")
}

func TestGoDirection(t *testing.T) {
  is := is.New(t)

  g := NewGame()
  res := g.GoDirection("south")
  is.Equal(res, "")
  res = g.GoDirection("north")
  is.Equal(res, "You are in a dusty break room containing an ancient card TABLE.  There is an exit to the SOUTH.")
}
