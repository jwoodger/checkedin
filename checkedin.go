package checkedin

// Check wraps around any function call that returns a value and an error, and
// then panics if the error is non-nil. Otherwise, the result is returned.
func Check[T any](result T, err error) T {
	if err != nil {
		panic(err)
	}
	return result
}
