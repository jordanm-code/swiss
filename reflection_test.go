package swiss

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  bool
	}{
		// int
		{name: "int empty", input: 0, want: true},
		{name: "int not empty", input: 1, want: false},

		// string
		{name: "string empty", input: "", want: true},
		{name: "string not empty", input: "a", want: false},

		// bool
		{name: "bool empty", input: false, want: true},
		{name: "bool not empty", input: true, want: false},

		// struct
		{name: "struct empty", input: struct{}{}, want: true},
		{name: "struct empty with fields", input: struct{ A string }{A: ""}, want: true},
		{name: "struct not empty with fields", input: struct{ A int }{A: 1}, want: false},

		// slice
		{name: "slice empty", input: []int{}, want: true},
		{name: "slice not empty", input: []int{1}, want: false},

		// map
		{name: "map empty", input: map[int]int{}, want: true},
		{name: "map not empty", input: map[int]int{1: 1}, want: false},

		// misc
		{name: "nil", input: nil, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsEmpty(test.input); got != test.want {
				t.Errorf("IsEmpty(%v) = %v; want %v", test.input, got, test.want)
			}
		})
	}
}
