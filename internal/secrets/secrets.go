package secrets

import (
	"github.com/pratishshr/envault/config"
	"github.com/pratishshr/envault/platform/aws"
	"github.com/pratishshr/envault/util/system/exit"
)

// GetSecrets sets appropriate config and fetches secrets from aws.
func GetSecrets(secretName string, env string, region string, profile string) map[string]string {
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

	if region == "" {
		region = conf.Region
	}

	if profile == "" {
		profile = conf.Profile
	}

	return aws.GetSecrets(profile, region, secretName)
}
