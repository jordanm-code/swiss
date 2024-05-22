package swiss

import (
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
