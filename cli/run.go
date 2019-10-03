package cli

import (
	"github.com/pratishshr/envault/internal/secrets"
	"github.com/pratishshr/envault/util/shell"
	"github.com/pratishshr/envault/util/system/exit"
)

// Run given command with the secrets from given Secret Manager.
func Run(secretName string, command string, env string, region string) {
	if command == "" {
		exit.Error("Command to run is not specified. Add command as 'envault run [command]'")
	}

	shell.Execute(command, secrets.GetSecrets(secretName, env, region))
}
