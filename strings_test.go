package swiss

import (
	"testing"
)

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
		name  string
		given string
		want  string
		err   error
	}{
		{"ok, all lowercase", "hello world", "helloWorld", nil},
		{"ok, mixed case", "HeLlO WoRlD", "helloWorld", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CamelCase(tt.given)
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}
