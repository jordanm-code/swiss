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
