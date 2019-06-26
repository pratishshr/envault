package exit

import (
	"fmt"
	"os"
)

// Error exits the application by showing the given message.
func Error(message string) {
	fmt.Println("Error: " + message)
	os.Exit(1)
}
