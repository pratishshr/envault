package aws

import (
	"bufio"
	"os"
	"sort"
	"strings"

	"github.com/pratishshr/envault/util/system/exit"
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

func getRegionsWithCode() map[string]string {
	return map[string]string{
		"US East (Ohio)":             "us-east-2",
		"US East (N. Virginia)":      "us-east-1",
		"US West (N. California)":    "us-west-1",
		"US West (Oregon)":           "us-west-2",
		"Asia Pacific (Hong Kong)":   "ap-east-1",
		"Asia Pacific (Mumbai)":      "ap-south-1",
		"Asia Pacific (Osaka-Local)": "ap-northeast-3",
		"Asia Pacific (Seoul)": "ap-northeast-2	",
		"Asia Pacific (Singapore)":  "ap-southeast-1",
		"Asia Pacific (Sydney)":     "ap-southeast-2",
		"Asia Pacific (Tokyo)":      "ap-northeast-1",
		"Canada (Central)":          "ca-central-1",
		"China (Beijing)":           "cn-north-1",
		"China (Ningxia)":           "cn-northwest-1",
		"EU (Frankfurt)":            "eu-central-1",
		"EU (Ireland)":              "eu-west-1",
		"EU (London)":               "eu-west-2",
		"EU (Paris)":                "eu-west-3",
		"EU (Stockholm)":            "eu-north-1",
		"South America (São Paulo)": "sa-east-1",
		"AWS GovCloud (US-East)":    "us-gov-east-1",
		"AWS GovCloud (US)":         "us-gov-west-1",
	}
}

// GetRegions returns a list of AWS region names.
func GetRegions() []string {
	var regions []string

	for key := range getRegionsWithCode() {
		regions = append(regions, key)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(regions)))

	return regions
}

// GetRegionCode returns the code of the provided AWS region.
func GetRegionCode(name string) string {
	return getRegionsWithCode()[name]
}
