package errors

import (
	"context"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Custom error types for testing
type CustomError struct {
	Code    int
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("custom error %d: %s", e.Code, e.Message)
}

type ValidationError struct {
	Field string
	Value string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field %s with value %s", e.Field, e.Value)
}

// Sentinel errors for testing
var (
	ErrSentinel1 = errors.New("sentinel error 1")
	ErrSentinel2 = errors.New("sentinel error 2")
)

func TestHandleError(t *testing.T) {
	t.Run("HandleError with nil error", func(t *testing.T) {
		handled, result := HandleError(nil)
		assert.True(t, handled)
		assert.NoError(t, result)
	})

	t.Run("HandleError with sentinel error match", func(t *testing.T) {
		err := io.EOF
		handled, result := HandleError(err,
			OnSentinelError(io.EOF, func(e error) error {
				return errors.New("handled EOF")
			}),
		)
		assert.True(t, handled)
		assert.EqualError(t, result, "handled EOF")
	})

	t.Run("HandleError with wrapped sentinel error", func(t *testing.T) {
		wrappedErr := fmt.Errorf("wrapped: %w", io.EOF)
		handled, result := HandleError(wrappedErr,
			OnSentinelError(io.EOF, func(e error) error {
				return errors.New("handled wrapped EOF")
			}),
		)
		assert.True(t, handled)
		assert.EqualError(t, result, "handled wrapped EOF")
	})

	t.Run("HandleError with custom error type match", func(t *testing.T) {
		err := &CustomError{Code: 404, Message: "not found"}
		handled, result := HandleError(err,
			OnCustomError(func(e *CustomError) error {
				return fmt.Errorf("handled custom error with code %d", e.Code)
			}),
		)
		assert.True(t, handled)
		assert.EqualError(t, result, "handled custom error with code 404")
	})

	t.Run("HandleError with wrapped custom error", func(t *testing.T) {
		customErr := &CustomError{Code: 500, Message: "internal error"}
		wrappedErr := fmt.Errorf("wrapped: %w", customErr)

		handled, result := HandleError(wrappedErr,
			OnCustomError(func(e *CustomError) error {
				return fmt.Errorf("handled wrapped custom error: %s", e.Message)
			}),
		)
		assert.True(t, handled)
		assert.EqualError(t, result, "handled wrapped custom error: internal error")
	})

	t.Run("HandleError with no matching handler", func(t *testing.T) {
		err := errors.New("unhandled error")
		handled, _ := HandleError(err,
			OnSentinelError(io.EOF, func(e error) error {
				return nil
			}),
			OnCustomError(func(e CustomError) error {
				return nil
			}),
		)
		assert.False(t, handled)
	})

	t.Run("HandleError with multiple matchers first match wins", func(t *testing.T) {
		err := io.EOF
		handled, result := HandleError(err,
			OnSentinelError(io.EOF, func(e error) error {
				return errors.New("first handler")
			}),
			OnSentinelError(io.EOF, func(e error) error {
				return errors.New("second handler")
			}),
		)
		assert.True(t, handled)
		assert.EqualError(t, result, "first handler")
	})

	t.Run("HandleError with handler returning nil", func(t *testing.T) {
		err := io.EOF
		handled, result := HandleError(err,
			OnSentinelError(io.EOF, func(e error) error {
				return nil
			}),
		)
		assert.True(t, handled)
		assert.NoError(t, result)
	})

	t.Run("HandleError with multiple error types", func(t *testing.T) {
		customErr := &CustomError{Code: 400, Message: "bad request"}
		validationErr := &ValidationError{Field: "email", Value: "invalid"}

		// Test CustomError handling
		handled1, result1 := HandleError(customErr,
			OnCustomError(func(e *CustomError) error {
				return fmt.Errorf("custom: %d", e.Code)
			}),
			OnCustomError(func(e *ValidationError) error {
				return fmt.Errorf("validation: %s", e.Field)
			}),
		)
		assert.True(t, handled1)
		require.EqualError(t, result1, "custom: 400")

		// Test ValidationError handling
		handled2, result2 := HandleError(validationErr,
			OnCustomError(func(e *CustomError) error {
				return fmt.Errorf("custom: %d", e.Code)
			}),
			OnCustomError(func(e *ValidationError) error {
				return fmt.Errorf("validation: %s", e.Field)
			}),
		)
		assert.True(t, handled2)
		assert.EqualError(t, result2, "validation: email")
	})

	t.Run("HandleError with context errors", func(t *testing.T) {
		// Test context.Canceled
		handled1, result1 := HandleError(context.Canceled,
			OnSentinelError(context.Canceled, func(e error) error {
				return errors.New("operation canceled")
			}),
		)
		assert.True(t, handled1)
		require.EqualError(t, result1, "operation canceled")

		// Test context.DeadlineExceeded
		handled2, result2 := HandleError(context.DeadlineExceeded,
			OnSentinelError(context.DeadlineExceeded, func(e error) error {
				return errors.New("deadline exceeded")
			}),
		)
		assert.True(t, handled2)
		assert.EqualError(t, result2, "deadline exceeded")
	})

	t.Run("HandleError preserves original error in handler", func(t *testing.T) {
		originalErr := &CustomError{Code: 403, Message: "forbidden"}
		var capturedErr error

		handled, _ := HandleError(originalErr,
			OnCustomError(func(e *CustomError) error {
				capturedErr = e
				return nil
			}),
		)

		assert.True(t, handled)
		assert.Equal(t, originalErr, capturedErr)
	})
}

func TestHandleErrorOrDie(t *testing.T) {
	t.Run("HandleErrorOrDie with nil error", func(t *testing.T) {
		result := HandleErrorOrDie(nil)
		assert.NoError(t, result)
	})

	t.Run("HandleErrorOrDie with matched error", func(t *testing.T) {
		err := io.EOF
		result := HandleErrorOrDie(err,
			OnSentinelError(io.EOF, func(e error) error {
				return errors.New("handled EOF in die")
			}),
		)
		assert.EqualError(t, result, "handled EOF in die")
	})

	t.Run("HandleErrorOrDie panics on unmatched error", func(t *testing.T) {
		err := errors.New("unmatched error")

		assert.Panics(t, func() {
			HandleErrorOrDie(err,
				OnSentinelError(io.EOF, func(e error) error {
					return nil
				}),
			)
		})
	})

	t.Run("HandleErrorOrDie with custom error panic", func(t *testing.T) {
		customErr := &CustomError{Code: 500, Message: "server error"}

		assert.Panics(t, func() {
			HandleErrorOrDie(customErr,
				OnCustomError(func(e *ValidationError) error {
					return nil
				}),
			)
		})
	})

	t.Run("HandleErrorOrDie with multiple matchers", func(t *testing.T) {
		validationErr := &ValidationError{Field: "username", Value: ""}

		result := HandleErrorOrDie(validationErr,
			OnSentinelError(io.EOF, func(e error) error {
				return errors.New("EOF handler")
			}),
			OnCustomError(func(e *CustomError) error {
				return errors.New("custom handler")
			}),
			OnCustomError(func(e *ValidationError) error {
				return fmt.Errorf("validation handler: field=%s", e.Field)
			}),
		)

		assert.EqualError(t, result, "validation handler: field=username")
	})
}

func TestOnSentinelError(t *testing.T) {
	t.Run("OnSentinelError creates proper matcher", func(t *testing.T) {
		handler := func(e error) error { return e }
		matcher := OnSentinelError(io.EOF, handler)

		assert.Equal(t, io.EOF, matcher.ErrorType)
		assert.True(t, matcher.IsSentinel)
		assert.NotNil(t, matcher.Handler)
	})

	t.Run("OnSentinelError with custom sentinel", func(t *testing.T) {
		customSentinel := errors.New("custom sentinel")
		callCount := 0

		matcher := OnSentinelError(customSentinel, func(e error) error {
			callCount++
			return errors.New("handled custom sentinel")
		})

		// Test that it matches the sentinel
		handled, result := HandleError(customSentinel, matcher)
		assert.True(t, handled)
		require.EqualError(t, result, "handled custom sentinel")
		assert.Equal(t, 1, callCount)

		// Test that it matches wrapped sentinel
		wrappedErr := fmt.Errorf("wrapped: %w", customSentinel)
		handled, result = HandleError(wrappedErr, matcher)
		assert.True(t, handled)
		require.EqualError(t, result, "handled custom sentinel")
		assert.Equal(t, 2, callCount)
	})
}

func TestOnCustomError(t *testing.T) {
	t.Run("OnCustomError creates proper matcher", func(t *testing.T) {
		matcher := OnCustomError(func(e *CustomError) error {
			return fmt.Errorf("handled: %d", e.Code)
		})

		assert.False(t, matcher.IsSentinel)
		assert.NotNil(t, matcher.Handler)

		// Test the handler works
		err := &CustomError{Code: 200, Message: "ok"}
		result := matcher.Handler(err)
		assert.EqualError(t, result, "handled: 200")
	})

	t.Run("OnCustomError with different error types", func(t *testing.T) {
		// Create matchers for different types
		customMatcher := OnCustomError(func(e *CustomError) error {
			return fmt.Errorf("custom error: code=%d", e.Code)
		})

		validationMatcher := OnCustomError(func(e *ValidationError) error {
			return fmt.Errorf("validation error: field=%s", e.Field)
		})

		// Test with CustomError
		customErr := &CustomError{Code: 404, Message: "not found"}
		handled, result := HandleError(customErr, customMatcher, validationMatcher)
		assert.True(t, handled)
		require.EqualError(t, result, "custom error: code=404")

		// Test with ValidationError
		validationErr := &ValidationError{Field: "age", Value: "-1"}
		handled, result = HandleError(validationErr, customMatcher, validationMatcher)
		assert.True(t, handled)
		assert.EqualError(t, result, "validation error: field=age")
	})

	t.Run("OnCustomError handler receives correct type", func(t *testing.T) {
		var receivedErr *CustomError

		matcher := OnCustomError(func(e *CustomError) error {
			receivedErr = e
			return nil
		})

		originalErr := &CustomError{Code: 301, Message: "redirect"}
		handled, _ := HandleError(originalErr, matcher)

		assert.True(t, handled)
		require.NotNil(t, receivedErr)
		assert.Equal(t, 301, receivedErr.Code)
		assert.Equal(t, "redirect", receivedErr.Message)
	})
}

func TestErrorMatcherEdgeCases(t *testing.T) {
	t.Run("Invalid sentinel error type in matcher", func(t *testing.T) {
		// Create a matcher with invalid ErrorType for sentinel
		matcher := ErrorMatcher{
			ErrorType:  "not an error", // Invalid type
			Handler:    func(e error) error { return e },
			IsSentinel: true,
		}

		err := errors.New("test error")
		handled, _ := HandleError(err, matcher)
		assert.False(t, handled)
	})

	t.Run("Handler that panics", func(t *testing.T) {
		matcher := OnSentinelError(io.EOF, func(e error) error {
			panic("handler panic")
		})

		assert.Panics(t, func() {
			HandleError(io.EOF, matcher)
		})
	})

	t.Run("Complex error chain", func(t *testing.T) {
		// Create a complex error chain
		baseErr := &CustomError{Code: 500, Message: "base"}
		wrapped1 := fmt.Errorf("layer1: %w", baseErr)
		wrapped2 := fmt.Errorf("layer2: %w", wrapped1)
		wrapped3 := fmt.Errorf("layer3: %w", wrapped2)

		handled, result := HandleError(wrapped3,
			OnCustomError(func(e *CustomError) error {
				return fmt.Errorf("found custom error at code %d", e.Code)
			}),
		)

		assert.True(t, handled)
		assert.EqualError(t, result, "found custom error at code 500")
	})
}
