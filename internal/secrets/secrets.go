package secrets

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pratishshr/envault/config"
	"github.com/pratishshr/envault/platform/aws"
	"github.com/pratishshr/envault/util/system/exit"
)

// GetSecrets sets appropriate config and fetches secrets from aws.
func GetSecrets(secretName string, env string, region string, profile string, envFile string) map[string]string {
	if envFile != "" {
		secrets, err := godotenv.Read(envFile)

		if err != nil {
			exit.Error("Could not read env file " + envFile)
		}

		return secrets
	}

	conf := config.GetConfig()

	if env == "" {
		env = os.Getenv("ENVIRONMENT")
	}

	if env == "" {
		env = conf.DefaultEnvironment
	}

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
