package cli

import (
	"github.com/pratishshr/envault/config"
	"github.com/pratishshr/envault/platform/aws"
	"github.com/pratishshr/envault/util/shell"
	"github.com/pratishshr/envault/util/system/exit"
)

// Run given command with the secrets from given Secret Manager.
func Run(secretName string, command string) {
	if secretName == "" {
		exit.Error("Secret Name is required to list environments. Set -secret flag.")
	}

	if command == "" {
		exit.Error("Command to run is not specified. Add command as 'envault run [command]'")
	}

	conf := config.GetConfig()
	secrets := aws.GetSecrets(conf.Profile, conf.Region, secretName)

	shell.Execute(command, secrets)
}
