package shell

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestConvertToSlice(t *testing.T) {
	tests := []struct {
		Name   string
		Input  map[string]string
		Output []string
	}{
		{
			Name: "With single key value",
			Input: map[string]string{
				"secretkey": "secretvalue",
			},
			Output: []string{"secretkey=secretvalue"},
		},
		{
			Name: "With many key value",
			Input: map[string]string{
				"secretkey1": "secretvalue",
				"secretkey2": "secretvalue",
				"secretkey3": "secretvalue",
			},
			Output: []string{"secretkey1=secretvalue", "secretkey2=secretvalue", "secretkey3=secretvalue"},
		},
		{
			Name:   "With zero key value",
			Input:  map[string]string{},
			Output: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			output := convertToSlice(tt.Input)
			// convertToSlice implementation loops over map and go does no specify the
			// iteration order so had to sort the output to test it reliably
			sort.Strings(output)
			if !reflect.DeepEqual(output, tt.Output) {
				t.Fatalf("convertToSlice() error:\ngot  %v\nwant %v", output, tt.Output)
			}
		})
	}
}

func TestExecute(t *testing.T) {
	tests := []struct {
		Name    string
		Command string
		Secrets map[string]string
		Error   string
	}{
		{
			Name:    "Running echo command",
			Command: "echo",
			Secrets: map[string]string{
				"secretkey": "secretvalue",
			},
		},
		{
			Name:    "Running eco command",
			Command: "eco",
			Secrets: map[string]string{
				"secretkey": "secretvalue",
			},
			Error: "Script Error: exit status 127",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			// Save current function and restore at the end:
			oldThrowError := throwError
			defer func() { throwError = oldThrowError }()

			var got string
			myThrowError := func(err interface{}) {
				got = fmt.Sprint(err)
			}

			throwError = myThrowError
			Execute(tt.Command, tt.Secrets)

			if got != tt.Error {
				t.Fatalf("Error() error:\ngot  %v\nwant %v", got, tt.Error)
			}
		})
	}
}
