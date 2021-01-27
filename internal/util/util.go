package util

// StringExistsInSlice will take a string slice and search it for the input string, returning true if it succeeds.
func StringExistsInSlice(slice []string, item string) bool {
  for _, v := range slice {
    if item == v {
      return true
    }
  }
  return false
}