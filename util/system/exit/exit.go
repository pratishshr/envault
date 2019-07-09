package exit

import (
	"fmt"
	"os"
)

// Error exits the application by showing the given message.
func Error(err interface{}) {
	fmt.Println("Error:")
	fmt.Println(err)
	os.Exit(1)
}
