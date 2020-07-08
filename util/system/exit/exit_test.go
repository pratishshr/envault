package exit

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	// Save current function and restore at the end:
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	osExit = myExit
	Error(errors.New("example error message"))
	if exp := 1; got != exp {
		t.Fatalf("Error() error:\ngot  %v\nwant %v", got, exp)
	}
}
