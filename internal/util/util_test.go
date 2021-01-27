package util

import (
	"testing"

	"github.com/matryer/is"
)

func TestIsOSSupported(t *testing.T) {
	is := is.New(t)

	is.Equal(IsOSSupported("linux"), true)
	is.Equal(IsOSSupported("windows"), false)
}

func TestStringExistsInSlice(t *testing.T) {
	is := is.New(t)

	abc := []string{"a", "b", "c"}

	is.Equal(StringExistsInSlice(abc, "b"), true)
	is.Equal(StringExistsInSlice(abc, "d"), false)
}
