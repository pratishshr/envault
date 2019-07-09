package cli

import (
	"fmt"

	"github.com/pratishshr/envault/config"
	"github.com/pratishshr/envault/platform/aws"
	"github.com/pratishshr/envault/util/system/exit"
)

// List all environment from Secrets Manager
func List(secretName string) {
	if secretName == "" {
		exit.Error("Secret Name is required to list environments. Set -secret flag.")
	}

	conf := config.GetConfig()
	secrets := aws.GetSecrets(conf.Profile, conf.Region, secretName)

	for key, value := range secrets {
		fmt.Println(key + "=" + value)
	}
}
