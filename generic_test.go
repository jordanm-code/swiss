package swiss

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestPtrTo(t *testing.T) {
	s := "string"
	if p := PtrTo(s); *p != s {
		t.Errorf("Deref(PtrTo(%q)) = %q; want %q", s, Deref(p), s)
	}

	tm := time.Now()
	if p := PtrTo(tm); *p != tm {
		t.Errorf("Deref(PtrTo(%q)) = %q; want %q", tm, Deref(p), tm)
	}

	i := 10
	if p := PtrTo(i); *p != i {
		t.Errorf("Deref(PtrTo(%d)) = %d; want %d", i, Deref(p), i)
	}

	i64 := int64(10)
	if p := PtrTo(i64); *p != i64 {
		t.Errorf("Deref(PtrTo(%d)) = %d; want %d", i64, Deref(p), i64)
	}

	b := true
	if p := PtrTo(b); *p != b {
		t.Errorf("Deref(PtrTo(%t)) = %t; want %t", b, Deref(p), b)
	}
}

func ExamplePtrTo() {
	printString := func(s *string) {
		fmt.Println(*s)
	}

	printString(PtrTo("easy"))
	// Output:
	// easy
}

func TestDeref(t *testing.T) {
	s := "string"
	if p := Deref(&s); p != s {
		t.Errorf("Deref(&%q) = %q; want %q", s, p, s)
	}

	tm := time.Now()
	if p := Deref(&tm); p != tm {
		t.Errorf("Deref(&%v) = %v; want %v", tm, p, tm)
	}

	i := 10
	if p := Deref(&i); p != i {
		t.Errorf("Deref(&%d) = %d; want %d", i, p, i)
	}

	i64 := int64(10)
	if p := Deref(&i64); p != i64 {
		t.Errorf("Deref(&%d) = %d; want %d", i64, p, i64)
	}

	b := true
	if p := Deref(&b); p != b {
		t.Errorf("Deref(&%t) = %t; want %t", b, p, b)
	}

	if p := Deref[string](nil); p != "" {
		t.Errorf("Deref[string](nil) = %q; want %q", p, "")
	}
}

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

func ExampleMap() {
	nums := []int{5, 9, 2, 0, 7, 12, 3, 26}
	m := Map(nums) // O(n)

	find := []int{5, 10}
	for _, n := range find {
		if m[n] { // O(1)
			fmt.Println("Found", n)
		} else {
			fmt.Println("Not found", n)
		}
	}

	// Output:
	// Found 5
	// Not found 10
}

func TestChunk(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chunks := Chunk(s, 3)

	want := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}}
	if !reflect.DeepEqual(chunks, want) {
		t.Errorf("Chunk(%v, 3) = %v; want: %v", s, chunks, want)
	}

	ss := strings.Split("hello world!!", "")
	strChunks := Chunk(ss, 2)
	wantStr := [][]string{{"h", "e"}, {"l", "l"}, {"o", " "}, {"w", "o"}, {"r", "l"}, {"d", "!"}, {"!"}}
	if !reflect.DeepEqual(strChunks, wantStr) {
		t.Errorf("Chunk(%v, 2) = %v; want: %v", ss, strChunks, wantStr)
	}
}

func TestDeduplicate(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 5, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 4, 1, 2, 7, 9, 20, 4, 2, 8}, []int{1, 2, 4, 7, 9, 20, 8}},
	}

	for _, test := range tests {
		if got := Deduplicate(test.input); !reflect.DeepEqual(got, test.expected) {
			t.Errorf("Deduplicate(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func ExampleDeduplicate() {
	names := []string{"Alice", "Bob", "Alice", "Charlie", "Bob", "David"}
	fmt.Println(Deduplicate(names))

	// Output:
	// [Alice Bob Charlie David]
}
