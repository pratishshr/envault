package cli

import (
	"fmt"

	"github.com/pratishshr/envault/internal/secrets"
)

// List all environment from Secrets Manager
func List(secretName string, env string, region string, profile string) {
	for key, value := range secrets.GetSecrets(secretName, env, region, profile) {
		fmt.Println(key + "=" + value)
	}
}
