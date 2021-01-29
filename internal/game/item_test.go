package game

import (
	"testing"

	"github.com/matryer/is"
)

func TestGetItem(t *testing.T) {
	is := is.New(t)

  g := NewGame()
  keyID := uint(1)

  res := g.GetItem(keyID)
  is.Equal(res, "There is no key.")
  _ = g.GoDirection("north")
  res = g.GetItem(keyID)
  is.Equal(res, "You got KEY.")
}
