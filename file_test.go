package swiss

import (
	"fmt"
)

func ExampleFileExists() {
	fmt.Printf("file_test.go exists: %v\n", FileExists("file_test.go"))
	fmt.Printf("file_test.go.not exists: %v\n", FileExists("file_test.go.not"))

	// Output:
	// file_test.go exists: true
	// file_test.go.not exists: false
}
