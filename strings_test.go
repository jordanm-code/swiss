package swiss

import (
	"fmt"
	"slices"
	"testing"
)

func TestIsEmail(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{"", false},
		{"a", false},
		{"a@", false},
		{"a@b", false},
		{"a@b.c", false},
		{"a@b.c.d", false},
		{`"bob is cool"@gmail.com`, false},
		{"a@b.c.do", true},
		{"bob@gmail.com", true},
		{"bob.smith@yahoo.com", true},
	}

	for _, test := range tests {
		if got := IsEmail(test.input); got != test.output {
			t.Errorf("IsEmail(%q) = %t; want %t", test.input, got, test.output)
		}
	}
}

func TestIsUpper(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{"", false},
		{"a", false},
		{"A", true},
		{"a b", false},
		{"A B", true},
		{"HELLO WORLD", true},
		{"hello world", false},
		{"HELLO world", false},
		{"hello WORLD", false},
	}

	for _, test := range tests {
		actual := IsUpper(test.input)
		if actual != test.output {
			t.Errorf("IsUpper(%q) = %t; want %t", test.input, actual, test.output)
		}
	}
}

func TestIsLower(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{"", false},
		{"a", true},
		{"A", false},
		{"a b", true},
		{"A B", false},
		{"hello world", true},
		{"HELLO WORLD", false},
		{"hello WORLD", false},
		{"HELLO world", false},
	}

	for _, test := range tests {
		actual := IsLower(test.input)
		if actual != test.output {
			t.Errorf("IsLower(%q) = %t; want %t", test.input, actual, test.output)
		}
	}
}

func TestSnakeCase(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"", ""},
		{"hello world", "hello_world"},
		{"Hello World", "hello_world"},
		{"helloWorld", "hello_world"},
		{"HelloWorld", "hello_world"},
	}

	for _, test := range tests {
		actual := SnakeCase(test.input)
		if actual != test.output {
			t.Errorf("SnakeCase(%q) = %q; want %q", test.input, actual, test.output)
		}
	}
}

func TestTitleCase(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"", ""},
		{"a", "A"},
		{"A", "A"},
		{"A B", "A B"},
		{"a b", "A B"},
		{"a b c", "A B C"},
		{"hello world", "Hello World"},
		{"HELLO WORLD", "Hello World"},
		{"hello WORLD", "Hello World"},
		{"HELLO world", "Hello World"},
		{"hello world!", "Hello World!"},
	}

	for _, test := range tests {
		actual := TitleCase(test.input)
		if actual != test.output {
			t.Errorf("TitleCase(%q) = %q; want %q", test.input, actual, test.output)
		}
	}
}

func TestCamelCase(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"hello world", "helloWorld"},
		{"HeLlO WoRlD", "helloWorld"},
		{"hello world", "helloWorld"},
		{"HELLO WORLD", "helloWorld"},
		{"hello WORLD", "helloWorld"},
		{"HELLO world", "helloWorld"},
		{"hello world!", "helloWorld!"},
		{"hello_world", "helloWorld"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if CamelCase(tt.input) != tt.output {
				t.Errorf("got %s, want %s", tt.input, tt.output)
			}
		})
	}
}

func ExampleCamelCase() {
	jsonKeys := []string{"first_name", "last_name"}
	for _, key := range jsonKeys {
		fmt.Println(CamelCase(key))
	}

	// Output:
	// firstName
	// lastName
}

func TestPascalCase(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"hello world", "HelloWorld"},
		{"HeLlO WoRlD", "HelloWorld"},
		{"hello world", "HelloWorld"},
		{"HELLO WORLD", "HelloWorld"},
		{"hello WORLD", "HelloWorld"},
		{"HELLO world", "HelloWorld"},
		{"hello world!", "HelloWorld!"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if PascalCase(tt.input) != tt.output {
				t.Errorf("got %s, want %s", tt.input, tt.output)
			}
		})
	}
}

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{"", false},
		{"a", true},
		{"A", true},
		{"AbCdeFg", true},
		{"AbC!deFg", false},
		{"1", false},
		{"a b", false},
		{"A B", false},
		{"hello world", false},
		{"HELLO WORLD", false},
		{"hello WORLD", false},
		{"HELLO world", false},
	}

	for _, test := range tests {
		if got := IsAlpha(test.input); got != test.output {
			t.Errorf("IsAlpha(%q) = %t; want %t", test.input, got, test.output)
		}
	}
}

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{"", false},
		{"a", false},
		{"!", false},
		{"1", true},
		{"1234567890", true},
		{"1 2 3 4 5 6 7 8 9 0", false},
		{"hello WORLD", false},
		{"HELLO world", false},
	}

	for _, test := range tests {
		if got := IsNumeric(test.input); got != test.output {
			t.Errorf("IsNumeric(%q) = %t; want %t", test.input, got, test.output)
		}
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{"", false},
		{"a", true},
		{"A", true},
		{"1", true},
		{"AbCdeFg", true},
		{"AbC!deFg", false},
		{"1", true},
		{"1 2 3 4 5 6 7 8 9 0", false},
		{"hello world", false},
		{"HELLO WORLD", false},
		{"hello WORLD", false},
		{"HELLO world", false},
	}

	for _, test := range tests {
		if got := IsAlphaNumeric(test.input); got != test.output {
			t.Errorf("IsAlphaNumeric(%q) = %t; want %t", test.input, got, test.output)
		}
	}
}

func TestIsHexChar(t *testing.T) {
	tests := []struct {
		input  byte
		output bool
	}{
		{'0', true},
		{'1', true},
		{'2', true},
		{'3', true},
		{'4', true},
		{'5', true},
		{'6', true},
		{'7', true},
		{'8', true},
		{'9', true},
		{'a', true},
		{'b', true},
		{'c', true},
		{'d', true},
		{'e', true},
		{'f', true},
		{'A', true},
		{'B', true},
		{'C', true},
		{'D', true},
		{'E', true},
		{'F', true},
		{'G', false},
		{'!', false},
	}
	for _, tt := range tests {
		if got := IsHexChar(tt.input); got != tt.output {
			t.Errorf("IsHexChar() = %v, want %v", got, tt.output)
		}
	}
}

func TestIsUUID(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output bool
	}{
		{"empty", "", false},
		{"nil", "00000000-0000-0000-0000-000000000000", true},
		{"too many", "00000000-00000-00000-00000-000000000000", false},
		{"too few", "00000000-0000-0000-0000-00000000000", false},
		{"wrong parts", "0000000-00000-0000-0000-00000000000", false},
		{"V1", "61800296-11e2-11ef-9262-0242ac120002", true},
		{"V2", "000003e8-11e2-21ef-8000-325096b39f47", true},
		{"V3", "8160c6f7-7c1e-31b8-a5ec-e9050e4c9a49", true},
		{"V4", "91b38875-7e07-4976-a231-9c28d7293da9", true},
		{"V5", "1c3af639-5b19-52df-8294-f41ae8707b37", true},
		{"V6", "016b21dc-8850-6946-9866-ef9afbd9b57d", true},
		{"V7", "018f76cc-ceb0-7fbe-957c-be10ada3d501", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUUID(tt.input); got != tt.output {
				t.Errorf("IsUUID() = %v, want %v", got, tt.output)
			}
		})
	}
}

func TestSwapCase(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"", ""},
		{"a", "A"},
		{"A", "a"},
		{"a b", "A B"},
		{"A B", "a b"},
		{"hello world", "HELLO WORLD"},
		{"HELLO WORLD", "hello world"},
		{"hello WORLD", "HELLO world"},
		{"HELLO world", "hello WORLD"},
		{"heLLo WoRLd", "HEllO wOrlD"},
	}

	for _, test := range tests {
		if got := SwapCase(test.input); got != test.output {
			t.Errorf("SwapCase(%q) = %q; want %q", test.input, got, test.output)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"hello world", "dlrow olleh"},
		{"حتوي على شريط التمرير على الجانب الأيمن", "نميألا بناجلا ىلع ريرمتلا طيرش ىلع يوتح"},
	}

	for _, test := range tests {
		got := Reverse(test.input)
		if got != test.output {
			t.Errorf("Reverse(%q) = %q; want %q", test.input, got, test.output)
		}
	}
}

func TestIsSnakeCase(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{"", false},
		{"hello", false},
		{"WORLD", false},
		{"Hello_world", false},
		{"hello world_there", false},
		{"a_b", true},
		{"hello_world_there", true},
	}

	for _, test := range tests {
		if got := IsSnakeCase(test.input); got != test.output {
			t.Errorf("IsSnakeCase(%q) = %t; want %t", test.input, got, test.output)
		}
	}
}

func TestSlugify(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"hello world", "hello-world"},
		{"HeLlo WoRld", "hello-world"},
		{" hello world ", "hello-world"},
		{"golang is awesome!!", "golang-is-awesome"},
		{"🎉 lets go 🎉", "lets-go"},
		{" Foo, Bar! baz qux", "foo-bar-baz-qux"},
	}

	for _, test := range tests {
		if got := Slugify(test.input); got != test.output {
			t.Errorf("Slugify(%q) = %q; want %q", test.input, got, test.output)
		}
	}
}

func TestIsURL(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{"", false},
		{"hello", false},
		{"http::://not.valid/??", false},
		{"google.com", false},
		{"/foo/bar", false},
		{"http://example.com", true},
		{"https://example", true},
		{"http://subdomain.example.com", true},
	}

	for _, test := range tests {
		if got := IsURL(test.input); got != test.output {
			t.Errorf("IsURL(%q) = %t; want %t", test.input, got, test.output)
		}
	}
}

func TestExtractURLs(t *testing.T) {
	tests := []struct {
		input  string
		output []string
	}{
		{"", nil},
		{"hello", nil},
		{"hello world", nil},
		{"hello http://example.com", []string{"http://example.com"}},
		{"hello https://example.com", []string{"https://example.com"}},
		{"hello http://example.com world", []string{"http://example.com"}},
		{"hello http://example.com world http://example2.com", []string{"http://example.com", "http://example2.com"}},
		{"hello http://example.com world https://example.com", []string{"http://example.com", "https://example.com"}},
		{"hello http://example.com world https://", []string{"http://example.com"}},
	}

	for _, test := range tests {
		if got := ExtractURLs(test.input); slices.Compare(got, test.output) != 0 {
			t.Errorf("ExtractURLs(%q) = %q; want %q", test.input, got, test.output)
		}
	}
}
