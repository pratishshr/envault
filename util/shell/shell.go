package shell

import (
	"os"
	"os/exec"

	"github.com/pratishshr/envault/util/system/exit"
)

func convertToSlice(secrets map[string]string) []string {
	var secretsSlice []string

	for key, value := range secrets {
		secretsSlice = append(secretsSlice, key+"="+value)
	}

	return secretsSlice
}

var throwError = exit.Error

// Execute runs the process with the supplied environment.
func Execute(command string, secrets map[string]string) {
	env := append(os.Environ(), convertToSlice(secrets)...)

	cmd := exec.Command("sh", "-c", command)
	cmd.Env = append(env, "IS_ENVAULT=true")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		throwError("Script Error: " + err.Error())
	}
}
