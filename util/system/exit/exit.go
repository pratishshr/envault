package exit

import (
	"fmt"
	"os"
)

var osExit = os.Exit

// Error exits the application by showing the given message.
func Error(err interface{}) {
	fmt.Println("Error: ")
	fmt.Println(err)

	osExit(1)
}
