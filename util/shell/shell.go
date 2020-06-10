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

// Execute runs the process with the supplied environment.
func Execute(command string, secrets map[string]string) {
	env := append(os.Environ(), convertToSlice(secrets)...)

	cmd := exec.Command("sh", "-c", command)
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		exit.Error("Script Error: " + err.Error())
	}
}
