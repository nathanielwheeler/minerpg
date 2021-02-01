package game

// Item holds description and other attributes of objects that can be placed in the player's inventory.
type Item struct {
  Name string
  Description string
  CanGet bool
  UnlocksDoor struct {
    RoomID uint
    Direction string
  }
}

// GetItem will check if an item is obtainable and if so put said item into the player's inventory.
func (g *Game) GetItem(id uint) string {
  return "There is no key."
}