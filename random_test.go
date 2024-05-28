package swiss

import "testing"

func TestRandomString(t *testing.T) {
	tests := []struct {
		length int
	}{
		{0},
		{5},
		{100},
		{1000},
	}
	for _, test := range tests {
		RandomSeed()
		got := RandomString(test.length)
		if len(got) != test.length {
			t.Errorf("RandomString(%d) = %d", test.length, len(got))
		}
	}
}
