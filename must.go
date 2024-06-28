// Package must provides a wrapper for calls to a function returning (T, error)
package must

// Must wraps a call to a function returning (T, error)
// and panics if the error is non-nil.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
