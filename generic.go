package swiss

// Map creates a map from a slice of keys.  The value of each key
// is a boolean indicating whether the key is present in the slice.
func Map[S ~[]E, E comparable](s S) map[E]bool {
	m := map[E]bool{}
	for _, v := range s {
		m[v] = true
	}
	return m
}
