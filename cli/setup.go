package cli

import (
	"github.com/pratishshr/envault/internal/cli/setup"
)

// Setup prompts user for required settings and creates a envault.json file
func Setup() {
	setup.Run()
}
