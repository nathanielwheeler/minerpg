package game

// Room holds the directions and entities within the room.
type Room struct {
	Name        string
	Description string
	Doors       []Door
	Items []Item
}

// Door holds the metadata required to journey to a different room.
type Door struct {
	RoomID    uint
	Direction string
	IsLocked bool 
	// IsVisible bool
	// GoDesc string
	// LookDesc string
}

// GenerateRooms will add a list of rooms to the game
func (g *Game) GenerateRooms() {
	g.Rooms = map[uint]Room{
    0: {
      Name: "test room",
      Description: "A bland, featureless room.  There is a DOOR labelled 'door that connects to this very room.'",
      Doors: []Door{Door{RoomID:0,Direction:"door"}},
    },
		1: {
			Name:        "Glistening Mine",
			Description: "You are in a glistening mine.  There is a passage to the NORTH and door to the EAST.",
			Doors: []Door{
				Door{RoomID: 2, Direction: "north"},
				Door{RoomID: 3, Direction: "east", IsLocked: true},
			},
		},
		2: {
      Name:        "Dusty Break Room",
      Description: "You are in a dusty break room containing an ancient card TABLE.  There is an exit to the SOUTH.",
      Items: []Item{
        Item{
          Name: "TABLE",
          Description: "The TABLE looks old and dusty.  There is a KEY resting on the surface.",
          CanGet: false,
        },
        Item{
          Name: "KEY",
          Description: "The KEY has a faint layer of rust on it.",
          CanGet: true,
          UnlocksDoor: struct{RoomID uint; Direction string}{
            RoomID: uint(1),
            Direction: "east",
          },
        },
      },
		},
    3: {
      Name:        "Mineshaft",
      Description: "You come to a deep mineshaft.  You can go back WEST or go DOWN.",
      Doors: []Door{
        Door{RoomID: 1, Direction: "south"},
        Door{RoomID: 3, Direction: "down"},
      },
    },
		4: {
			Name:        "Underground Lake",
			Description: "You are next to an underground lake.  Looks like you're stuck here!",
		},
	}
}

// GetRoomDescription will return the room's description given an id
func (g *Game) GetRoomDescription(id uint) string {
	description := g.Rooms[id].Description
	return description
}

// GoDirection will check if the direction is valid given the room, change the game's CurrentRoomID, and return the new room's description.
func (g *Game) GoDirection(dir string) string {
	room := g.Rooms[g.CurrentRoomID]

	for _, door := range room.Doors {
		if dir == door.Direction {
			g.CurrentRoomID = door.RoomID
			return g.GetRoomDescription(g.CurrentRoomID)
		}
	}

	return ""
}
