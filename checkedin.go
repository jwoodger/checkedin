package checkedin

func panicOnFalse(pred bool, message string) {
	if !pred {
		panic(message)
	}
}

// Check wraps around any function call that returns a value and an error, and
// then panics if the error is non-nil. Otherwise, the result is returned.
func Check[T any](result T, err error) T {
	if err != nil {
		panic(err)
	}
	return result
}

// Requires checks a precondition in the code and then panics if the condition
// is not met.
func Requires(condition bool) {
	panicOnFalse(condition, "Precondition error.")
}

// Ensures checks a postcondition in the code and then panics if the condition
// is not met.
func Ensures(condition bool) {
	panicOnFalse(condition, "Postcondition error.")
}
