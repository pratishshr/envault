package file

import (
	"reflect"
	"testing"
)

func TestExists(t *testing.T) {
	tests := []struct {
		Name   string
		Path   string
		Output bool
	}{
		{
			Name:   "Path Exists",
			Path:   "./file.go",
			Output: true,
		},
		{
			Name:   "Path Does not Exists",
			Path:   "/bad/path",
			Output: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			exists := Exists(tt.Path)
			if !reflect.DeepEqual(exists, tt.Output) {
				t.Fatalf("Exists() error:\ngot  %v\nwant %v", exists, tt.Output)
			}
		})
	}
}
