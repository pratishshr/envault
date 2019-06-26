package cli

import (
	"github.com/pratishshr/envault/platforms/aws"
	"github.com/pratishshr/envault/utils/shell"
	"github.com/pratishshr/envault/utils/system/exit"
)

// Run given command with the secrets from given Secret Manager.
func Run(secretName string, command string) {
	if secretName == "" {
		exit.Error("Secret Name is required to list environments. Set -secret flag.")
	}

	if command == "" {
		exit.Error("Command to run is not specified. Add command as 'envault run [command]'")
	}

	secrets := aws.GetSecrets(secretName)

	shell.Execute(command, secrets)
}
