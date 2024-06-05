package swiss

import (
	"fmt"
)

func ExampleFileExists() {
	if FileExists("file_test.go") {
		fmt.Println("file_test.go exists")
	}

	if !FileExists("file_test.go.not") {
		fmt.Println("file_test.go.not not exists")
	}

	// Output:
	// file_test.go exists
	// file_test.go.not not exists
}
