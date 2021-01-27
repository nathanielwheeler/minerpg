package util

// IsOSSupported will return false if the OS is not on the supported list.
func IsOSSupported(goos string) bool {
	supported := []string{
		"darwin",
		"freebsd",
		"linux",
		"netbsd",
		"openbsd",
	}
	if !StringExistsInSlice(supported, goos) {
		return false
	}
	return true
}

// StringExistsInSlice will take a string slice and search it for the input string, returning true if it succeeds.
func StringExistsInSlice(slice []string, item string) bool {
	for _, v := range slice {
		if item == v {
			return true
		}
	}
	return false
}
