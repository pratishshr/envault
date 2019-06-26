package cli

import (
	"fmt"

	"github.com/pratishshr/envault/platforms/aws"
	"github.com/pratishshr/envault/utils/system/exit"
)

// List all environment from Secrets Manager
func List(secretName string) {
	if secretName == "" {
		exit.Error("Secret Name is required to list environments. Set -secret flag.")
	}

	secrets := aws.GetSecrets(secretName)

	for key, value := range secrets {
		fmt.Println(key + "=" + value)
	}
}
