package swiss

import (
	"testing"
)

func TestFileExists(t *testing.T) {
	if !FileExists("file_test.go") {
		t.Error("file_test.go should exist")
	}

	if FileExists("file_test.go.not") {
		t.Error("file_test.go.not should not exist")
	}
}
