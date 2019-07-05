package aws

import (
	"bufio"
	"os"
	"strings"

	"github.com/pratishshr/envault/utils/system/exit"
)

func handle(err error) {
	if err != nil {
		exit.Error(err)
	}
}

func closeFile(file *os.File) {
	err := file.Close()

	handle(err)
}

// GetProfiles returns aws profiles stored in ~/.aws/credentials
func GetProfiles() []string {
	var profiles []string

	homeDir, err := os.UserHomeDir()
	handle(err)

	awsConfigPath := homeDir + "/.aws/credentials"

	file, err := os.Open(awsConfigPath)
	handle(err)

	defer closeFile(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(text, "[") {
			profiles = append(profiles, strings.Trim(text, "[]"))
		}
	}

	return profiles
}
