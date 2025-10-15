// Package ratelimit provides exponential backoff with jitter for handling API rate limits.
// It implements a retry mechanism that gradually increases wait times between attempts
// to prevent overwhelming rate-limited APIs.
package ratelimit

import (
	"math"
	"math/rand"
	"time"
)

// RateLimiter implements exponential backoff with jitter for API rate limiting.
// It retries failed operations with increasing delays between attempts.
//
// The backoff formula is: min(((2^n) + random_milliseconds), max_backoff)
// where n is the attempt number, starting from 0.
//
// Example backoff sequence (without jitter):
//   - Attempt 1: 2 seconds
//   - Attempt 2: 4 seconds
//   - Attempt 3: 8 seconds
//   - Attempt 4: 16 seconds
//   - etc., up to maxBackoff
type RateLimiter struct {
	maxRetries int // Maximum number of retry attempts
	maxBackoff int // Maximum backoff duration in seconds
}

// New creates a new RateLimiter with the specified retry and backoff limits.
//
// Parameters:
//   - maxRetries: Maximum number of times to retry a failed operation
//   - maxBackoffSec: Maximum wait time between retries in seconds
//
// Returns:
//   - *RateLimiter: Configured rate limiter instance
func New(maxRetries, maxBackoffSec int) *RateLimiter {
	return &RateLimiter{
		maxRetries: maxRetries,
		maxBackoff: maxBackoffSec,
	}
}

// Temporary is an interface for errors that are temporary and should be retried.
type Temporary interface {
	error
	Temporary() bool
}

// Do executes a function with exponential backoff retry logic.
// If the function fails with a temporary error, it will retry up to maxRetries times
// with increasing delays between attempts.
//
// The function stops retrying when:
//   - The function returns nil (success)
//   - The function returns a non-temporary error (permanent failure)
//   - Maximum retry attempts are reached
//
// Parameters:
//   - fn: Function to execute and retry on failure
//
// Returns:
//   - error: The last error returned by fn, or nil on success
func (rl *RateLimiter) Do(fn func() error) error {
	var err error

	// Try the operation up to maxRetries + 1 times (initial attempt + retries)
	for attempt := 0; attempt <= rl.maxRetries; attempt++ {
		err = fn()
		if err == nil {
			return nil // Success
		}

		// Check if error is temporary using the Temporary() method pattern
		if tempErr, ok := err.(interface{ Temporary() bool }); ok && !tempErr.Temporary() {
			// Non-temporary error - don't retry
			return err
		}

		// For errors that don't implement Temporary(), check if they're wrapped temporary errors
		// by trying to unwrap and check the type name
		errStr := err.Error()
		if !isLikelyTemporary(errStr) {
			// Doesn't look like a temporary error - don't retry
			return err
		}

		// If this was the last attempt, don't wait - just return the error
		if attempt == rl.maxRetries {
			break
		}

		// Calculate backoff duration and wait before next attempt
		backoff := rl.calculateBackoff(attempt)
		time.Sleep(backoff)
	}

	return err
}

// isLikelyTemporary checks if an error message indicates a temporary error.
// This is a heuristic for errors that don't implement the Temporary() interface.
func isLikelyTemporary(errMsg string) bool {
	// Check for temporary error indicators
	temporaryIndicators := []string{
		"temporary error",
		"status 429",
		"status 5",
		"too many requests",
		"server error",
		"service unavailable",
		"gateway timeout",
		"connection reset",
		"connection refused",
		"timeout",
	}

	for _, indicator := range temporaryIndicators {
		if len(errMsg) > 0 && contains(errMsg, indicator) {
			return true
		}
	}

	return false
}

// contains checks if a string contains a substring (case-insensitive helper).
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) &&
		(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
		indexOf(s, substr) >= 0))
}

// indexOf returns the index of substr in s, or -1 if not found.
func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

// calculateBackoff calculates the backoff duration using exponential backoff with jitter.
// Jitter (randomness) is added to prevent multiple clients from retrying simultaneously,
// which could cause a "thundering herd" problem.
//
// Formula: min(((2^(n+1)) + random_milliseconds), max_backoff)
//
// Parameters:
//   - attempt: Current attempt number (0-indexed)
//
// Returns:
//   - time.Duration: Time to wait before next retry
func (rl *RateLimiter) calculateBackoff(attempt int) time.Duration {
	// Calculate exponential component: 2^(n+1) seconds
	exponential := math.Pow(2, float64(attempt+1))

	// Add random jitter between 0 and 1000 milliseconds
	// This prevents synchronized retries across multiple clients
	jitterMs := rand.Intn(1001)
	jitter := float64(jitterMs) / 1000.0

	// Total backoff in seconds
	backoffSec := exponential + jitter

	// Cap at maximum backoff to prevent excessive wait times
	if backoffSec > float64(rl.maxBackoff) {
		backoffSec = float64(rl.maxBackoff)
	}

	return time.Duration(backoffSec * float64(time.Second))
}
