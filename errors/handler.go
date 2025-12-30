package errors

import (
	"errors"
	"fmt"
	"reflect"
)

// ErrorHandler represents a function that handles a specific error type
type ErrorHandler func(error) error

// ErrorMatcher holds the error type/value and its handler
type ErrorMatcher struct {
	ErrorType  any // Can be error value (sentinel) or error type
	Handler    ErrorHandler
	IsSentinel bool // true for sentinel errors, false for custom types
}

// HandleError processes an error against a list of matchers and executes the appropriate handler.
// It returns (true, handlerResult) if a matching handler is found and executed,
// or (false, nil) if no matcher matches the error.
// If err is nil, returns (true, nil).
//
// Deprecated: HandleError is deprecated and will be removed in a future release.
// Use Handle instead, which provides the same functionality.
//
// Example:
//
//	handled, result := HandleError(err,
//	    OnSentinelError(io.EOF, func(e error) error {
//	        return nil // EOF is expected, ignore it
//	    }),
//	    OnCustomError(func(e *CustomError) error {
//	        return fmt.Errorf("custom error: %w", e)
//	    }),
//	)
func HandleError(err error, matchers ...ErrorMatcher) (bool, error) {
	if err == nil {
		return true, nil
	}

	for _, matcher := range matchers {
		if matcher.IsSentinel {
			// Handle sentinel errors with errors.Is
			if sentinelErr, ok := matcher.ErrorType.(error); ok {
				if errors.Is(err, sentinelErr) {
					return true, matcher.Handler(err)
				}
			}
		} else {
			// Handle custom error types with errors.As
			errorType := reflect.TypeOf(matcher.ErrorType)
			errorValue := reflect.New(errorType).Interface()

			if errors.As(err, errorValue) {
				return true, matcher.Handler(err)
			}
		}
	}

	return false, err // No matcher found
}

// Handle processes an error against a list of matchers and executes the appropriate handler.
// It returns (true, handlerResult) if a matching handler is found and executed,
// or (false, nil) if no matcher matches the error.
// If err is nil, returns (true, nil).
//
// Example:
//
//	handled, result := Handle(err,
//	    OnSentinel(io.EOF, func(e error) error {
//	        return nil // EOF is expected, ignore it
//	    }),
//	    OnType(func(e *CustomError) error {
//	        return fmt.Errorf("custom error: %w", e)
//	    }),
//	)
var Handle = HandleError

// HandleErrorOrDie processes an error against a list of matchers and executes the appropriate handler.
// If a matching handler is found, it returns the handler's result.
// If no matcher matches the error, it panics with a descriptive message.
// This function is useful when all expected error types must be handled explicitly.
//
// Deprecated: HandleErrorOrDie is deprecated and will be removed in a future release.
// Use MustHandle instead, which provides the same functionality.
//
// Example:
//
//	result := HandleErrorOrDie(err,
//	    OnSentinelError(context.Canceled, func(e error) error {
//	        return fmt.Errorf("operation canceled")
//	    }),
//	    OnCustomError(func(e *ValidationError) error {
//	        return fmt.Errorf("validation failed: %w", e)
//	    }),
//	) // Panics if err doesn't match any handler
func HandleErrorOrDie(err error, matchers ...ErrorMatcher) error {
	ok, err := HandleError(err, matchers...)
	if !ok {
		panic(fmt.Sprintf("Unhandled error of type %T: %v", err, err))
	}
	return err
}

// MustHandle processes an error against a list of matchers and executes the appropriate handler.
// If a matching handler is found, it returns the handler's result.
// If no matcher matches the error, it panics with a descriptive message.
// This function is useful when all expected error types must be handled explicitly.
//
// Example:
//
//	result := MustHandle(err,
//	    OnSentinel(context.Canceled, func(e error) error {
//	        return fmt.Errorf("operation canceled")
//	    }),
//	    OnType(func(e *ValidationError) error {
//	        return fmt.Errorf("validation failed: %w", e)
//	    }),
//	) // Panics if err doesn't match any handler
var MustHandle = HandleErrorOrDie

// HandleErrorOrDefault processes an error against a list of matchers and executes the appropriate handler.
// If a matching handler is found, it returns the handler's result.
// If no matcher matches the error, it executes the default handler (dft) and returns its result.
// If dft is nil, unmatched errors return nil (effectively suppressing the error).
// This function is useful when you want to handle specific error cases explicitly
// while providing a fallback handler for all other errors.
//
// Deprecated: HandleErrorOrDefault is deprecated and will be removed in a future release.
// Use HandleOr instead, which provides the same functionality.
//
// Example:
//
//	result := HandleErrorOrDefault(err,
//	    func(e error) error {
//	        // Default handler for unmatched errors
//	        return fmt.Errorf("unexpected error: %w", e)
//	    },
//	    OnSentinelError(context.Canceled, func(e error) error {
//	        return fmt.Errorf("operation canceled")
//	    }),
//	    OnCustomError(func(e *ValidationError) error {
//	        return fmt.Errorf("validation failed: %w", e)
//	    }),
//	)
//
//	// Suppress unmatched errors by passing nil as default handler
//	result := HandleErrorOrDefault(err, nil,
//	    OnSentinelError(io.EOF, func(e error) error {
//	        return errors.New("EOF handled")
//	    }),
//	) // Returns nil for unmatched errors
func HandleErrorOrDefault(err error, dft ErrorHandler, matchers ...ErrorMatcher) error {
	ok, err := HandleError(err, matchers...)
	if !ok {
		if dft == nil {
			return nil
		}
		return dft(err)
	}
	return err
}

// HandleOr processes an error against a list of matchers and executes the appropriate handler.
// If a matching handler is found, it returns the handler's result.
// If no matcher matches the error, it executes the default handler (dft) and returns its result.
// If dft is nil, unmatched errors return nil (effectively suppressing the error).
// This function is useful when you want to handle specific error cases explicitly
// while providing a fallback handler for all other errors.
//
// Example:
//
//	result := HandleOr(err,
//	    func(e error) error {
//	        // Default handler for unmatched errors
//	        return fmt.Errorf("unexpected error: %w", e)
//	    },
//	    OnSentinel(context.Canceled, func(e error) error {
//	        return fmt.Errorf("operation canceled")
//	    }),
//	    OnType(func(e *ValidationError) error {
//	        return fmt.Errorf("validation failed: %w", e)
//	    }),
//	)
//
//	// Suppress unmatched errors by passing nil as default handler
//	result := HandleOr(err, nil,
//	    OnSentinel(io.EOF, func(e error) error {
//	        return errors.New("EOF handled")
//	    }),
//	) // Returns nil for unmatched errors
var HandleOr = HandleErrorOrDefault

// OnSentinelError creates an ErrorMatcher for sentinel errors.
// Sentinel errors are predefined error values that are compared using errors.Is.
//
// This is used with HandleError or HandleErrorOrDie to match specific error
// values like io.EOF, context.Canceled, or custom sentinel errors defined with
// errors.New or fmt.Errorf.
//
// The handler function receives the original error and can return a new error
// or nil to suppress it.
//
// Deprecated: OnSentinelError is deprecated and will be removed in a future release.
// Use OnSentinel instead, which provides the same functionality.
//
// Example:
//
//	matcher := OnSentinelError(io.EOF, func(e error) error {
//	    log.Println("reached end of file")
//	    return nil // suppress EOF error
//	})
func OnSentinelError(sentinelErr error, handler ErrorHandler) ErrorMatcher {
	return ErrorMatcher{
		ErrorType:  sentinelErr,
		Handler:    handler,
		IsSentinel: true,
	}
}

// OnSentinel creates an ErrorMatcher for sentinel errors.
// Sentinel errors are predefined error values that are compared using errors.Is.
//
// This is used with Handle or MustHandle to match specific error
// values like io.EOF, context.Canceled, or custom sentinel errors defined with
// errors.New or fmt.Errorf.
//
// The handler function receives the original error and can return a new error
// or nil to suppress it.
//
// Example:
//
//	matcher := OnSentinel(io.EOF, func(e error) error {
//	    log.Println("reached end of file")
//	    return nil // suppress EOF error
//	})
var OnSentinel = OnSentinelError

// OnCustomError creates an ErrorMatcher for custom error types.
// Custom error types are struct types that implement the error interface,
// and are matched using errors.As to unwrap error chains.
//
// The type parameter T specifies the error type to match. The handler function
// receives the unwrapped typed error, allowing you to access type-specific fields
// and methods.
//
// This is particularly useful for handling errors with additional context or data,
// such as validation errors, network errors, or domain-specific errors.
//
// Deprecated: OnCustomError is deprecated and will be removed in a future release.
// Use OnType instead, which provides the same functionality.
//
// Example:
//
//	type ValidationError struct {
//	    Field string
//	    Msg   string
//	}
//	func (e *ValidationError) Error() string {
//	    return fmt.Sprintf("%s: %s", e.Field, e.Msg)
//	}
//
//	matcher := OnCustomError(func(e *ValidationError) error {
//	    log.Printf("validation failed on field %s: %s", e.Field, e.Msg)
//	    return fmt.Errorf("invalid input: %w", e)
//	})
func OnCustomError[T error](handler func(T) error) ErrorMatcher {
	var zero T
	return ErrorMatcher{
		ErrorType: zero,
		Handler: func(err error) error {
			var typedErr T
			if errors.As(err, &typedErr) {
				return handler(typedErr)
			}
			return nil
		},
		IsSentinel: false,
	}
}

// OnType creates an ErrorMatcher for custom error types.
// Custom error types are struct types that implement the error interface,
// and are matched using errors.As to unwrap error chains.
//
// The type parameter T specifies the error type to match. The handler function
// receives the unwrapped typed error, allowing you to access type-specific fields
// and methods.
//
// This is particularly useful for handling errors with additional context or data,
// such as validation errors, network errors, or domain-specific errors.
//
// Example:
//
//	type ValidationError struct {
//	    Field string
//	    Msg   string
//	}
//	func (e *ValidationError) Error() string {
//	    return fmt.Sprintf("%s: %s", e.Field, e.Msg)
//	}
//
//	matcher := OnType(func(e *ValidationError) error {
//	    log.Printf("validation failed on field %s: %s", e.Field, e.Msg)
//	    return fmt.Errorf("invalid input: %w", e)
//	})
func OnType[T error](handler func(T) error) ErrorMatcher {
	return OnCustomError(handler)
}
