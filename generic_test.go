package swiss

import "testing"

// TestMap tests the Map function.
func TestMap(t *testing.T) {
	i := Map([]int{1, 2, 3, 3})
	if !i[1] || !i[2] || !i[3] || len(i) != 3 {
		t.Errorf("Map failed (int): %v", i)
	}

	s := Map([]string{"a", "b", "c", "c"})
	if !s["a"] || !s["b"] || !s["c"] || len(s) != 3 {
		t.Errorf("Map failed (string): %v", s)
	}
}
