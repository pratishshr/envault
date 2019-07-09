package cli

import (
	"fmt"

	"github.com/pratishshr/envault/config"
	"github.com/pratishshr/envault/platform/aws"
	"github.com/pratishshr/envault/util/system/exit"
)

// List all environment from Secrets Manager
func List(secretName string, env string) {
	conf := config.GetConfig()

	if secretName == "" && env == "" {
		exit.Error("Secret Name is required to list environments. Set -secret flag.")
	}

	if secretName == "" && env != "" {
		if _, ok := conf.Environments[env]; !ok {
			exit.Error("Environment '" + env + "' does not exist.")
		}

		secretName = conf.Environments[env]
	}

	secrets := aws.GetSecrets(conf.Profile, conf.Region, secretName)

	for key, value := range secrets {
		fmt.Println(key + "=" + value)
	}
}
