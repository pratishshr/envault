package shell

import (
	"os"
	"os/exec"
)

func convertToSlice(secrets map[string]string) []string {
	secretsSlice := []string{}

	for key, value := range secrets {
		secretsSlice = append(secretsSlice, key+"="+value)
	}

	return secretsSlice
}

// Execute runs the process with the supplied environment.
func Execute(command string, secrets map[string]string) {
	env := append(os.Environ(), convertToSlice(secrets)...)

	cmd := exec.Command("bash", "-c", command)
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
