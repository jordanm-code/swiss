package swiss

import "testing"

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
