package swiss

// PtrTo returns a pointer to the value passed in.
func PtrTo[T any](v T) *T {
	return &v
}

// Deref returns the value that the pointer points to.
func Deref[T any](v *T) (r T) {
	if v == nil {
		return
	}
	r = *v
	return
}

// Map creates a map from a slice of keys.  The value of each key
// is a boolean indicating whether the key is present in the slice.
func Map[S ~[]E, E comparable](s S) map[E]bool {
	m := map[E]bool{}
	for _, v := range s {
		m[v] = true
	}
	return m
}
