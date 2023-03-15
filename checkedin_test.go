package checkedin_test

import (
	"errors"
	"testing"

	ci "github.com/jwoodger/checkedin"
)

func shouldPanic(t *testing.T, fn func()) {
	t.Helper()

	// Recover after any panics which occur.
	defer func() {
		_ = recover()
	}()

	fn()

	// If execution gets to this point then there was no panic.
	t.Error("Function should have panicked.")
}

func TestCheckShouldPanic(t *testing.T) {
	shouldPanic(t, func() {
		_ = ci.Check(1, errors.New(""))
	})
}

func TestCheckDoesNotPanic(t *testing.T) {
	_ = ci.Check(1, nil)
}

func TestRequiresShouldPanic(t *testing.T) {
	shouldPanic(t, func() {
		ci.Requires(2+2 == 5)
	})
}

func TestRequiresDoesNotPanic(t *testing.T) {
	ci.Ensures(1+1 == 2)
}

func TestEnsuresShouldPanic(t *testing.T) {
	shouldPanic(t, func() {
		ci.Requires(2+2 == 5)
	})
}

func TestEnsuresDoesNotPanic(t *testing.T) {
	ci.Ensures(1+1 == 2)
}
