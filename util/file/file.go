package file

import (
	"os"
)

// Exists checks if the filepath exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}

	return true
}
