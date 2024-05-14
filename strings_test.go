package swiss

import (
	"testing"
)

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
