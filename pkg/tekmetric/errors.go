package tekmetric

// temporaryError represents a temporary error that should be retried.
// This includes rate limit errors (429) and server errors (5xx).
type temporaryError struct {
	statusCode int
	message    string
}

func (e *temporaryError) Error() string {
	return e.message
}

// Temporary returns true indicating this error is temporary and should be retried.
func (e *temporaryError) Temporary() bool {
	return true
}
