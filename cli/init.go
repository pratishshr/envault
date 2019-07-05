package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/AlecAivazis/survey/v2"
	"github.com/pratishshr/envault/platforms/aws"
)

type answers struct {
	Profile string `json:"profile"`
}

func askQuestions() *answers {
	questions := []*survey.Question{
		{
			Name: "profile",
			Prompt: &survey.Select{
				Message: "Choose AWS profile:",
				Options: aws.GetProfiles(),
			},
		},
	}

	answers := &answers{}

	err := survey.Ask(questions, answers)
	if err != nil {
		fmt.Println(err)
	}

	return answers
}

// Init prompts user for required settings and creates a envault.json file
func Init() {
	answers := askQuestions()
	config, _ := json.MarshalIndent(answers, "", "  ")

	ioutil.WriteFile("envault.json", config, 0644)
}
