package parser

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// ParseStringSupportedTypes defines the type constraint for types that can be
// parsed from strings using the ParseString family of functions.
type ParseStringSupportedTypes interface {
	string |
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float64 |
		bool | time.Duration | url.URL
}

// ParseString parses a string value into the specified type T.
// It uses the appropriate parsing function based on the target type.
//
// The function supports all types defined in ParseStringSupportedTypes.
// For integers, it parses base-10 numbers with appropriate bit sizes.
// For booleans, it accepts: "1", "t", "T", "TRUE", "true", "True", "0", "f", "F", "FALSE", "false", "False".
// For durations, it accepts strings like "300ms", "1.5h", "2h45m".
// For URLs, it parses according to RFC 3986.
//
// Returns an error if the string cannot be parsed into the target type.
//
// Example:
//
//	num, err := ParseString[int]("42")
//	duration, err := ParseString[time.Duration]("5s")
//	isValid, err := ParseString[bool]("true")
//
//nolint:forcetypeassert
func ParseString[T ParseStringSupportedTypes](rawValue string) (T, error) { //nolint:forcetypeassert
	var value T

	switch any(value).(type) {
	case string:
		value = any(rawValue).(T)
	case int:
		i, err := strconv.Atoi(rawValue)
		if err != nil {
			return value, err
		}
		value = any(i).(T)
	case int8:
		i, err := strconv.ParseInt(rawValue, 10, 8)
		if err != nil {
			return value, err
		}
		value = any(int8(i)).(T)
	case int16:
		i, err := strconv.ParseInt(rawValue, 10, 16)
		if err != nil {
			return value, err
		}
		value = any(int16(i)).(T)
	case int32:
		i, err := strconv.ParseInt(rawValue, 10, 32)
		if err != nil {
			return value, err
		}
		value = any(int32(i)).(T)
	case int64:
		i, err := strconv.ParseInt(rawValue, 10, 64)
		if err != nil {
			return value, err
		}
		value = any(i).(T)
	case uint:
		u, err := strconv.ParseUint(rawValue, 10, 0)
		if err != nil {
			return value, err
		}
		value = any(uint(u)).(T)
	case uint8:
		u, err := strconv.ParseUint(rawValue, 10, 8)
		if err != nil {
			return value, err
		}
		value = any(uint8(u)).(T)
	case uint16:
		u, err := strconv.ParseUint(rawValue, 10, 16)
		if err != nil {
			return value, err
		}
		value = any(uint16(u)).(T)
	case uint32:
		u, err := strconv.ParseUint(rawValue, 10, 32)
		if err != nil {
			return value, err
		}
		value = any(uint32(u)).(T)
	case uint64:
		u, err := strconv.ParseUint(rawValue, 10, 64)
		if err != nil {
			return value, err
		}
		value = any(u).(T)
	case float64:
		f, err := strconv.ParseFloat(rawValue, 64)
		if err != nil {
			return value, err
		}
		value = any(f).(T)
	case bool:
		b, err := strconv.ParseBool(rawValue)
		if err != nil {
			return value, err
		}
		value = any(b).(T)
	case time.Duration:
		d, err := time.ParseDuration(rawValue)
		if err != nil {
			return value, err
		}
		value = any(d).(T)
	case url.URL:
		u, err := url.Parse(rawValue)
		if err != nil {
			return value, err
		}
		value = any(*u).(T)
	default:
		return value, fmt.Errorf("unsupported type: %T", value)
	}

	return value, nil
}

// ParseStringOrZero parses a string value into the specified type T.
// If parsing fails, it returns the zero value for type T instead of an error.
//
// This is useful when you want a simple conversion that falls back to a
// default zero value on error.
//
// Example:
//
//	num := ParseStringOrZero[int]("invalid") // returns 0
//	num := ParseStringOrZero[int]("42")      // returns 42
func ParseStringOrZero[T ParseStringSupportedTypes](rawValue string) T {
	parsedValue, _ := ParseString[T](rawValue)
	return parsedValue
}

// ParseStringWithDefault parses a string value into the specified type T.
// If parsing fails, it returns the provided default value along with the error.
//
// Unlike ParseString, this ensures a usable value is always returned,
// but still provides error information for logging or handling.
//
// Example:
//
//	num, err := ParseStringWithDefault("invalid", 10)
//	// returns: 10, error
//	num, err := ParseStringWithDefault("42", 10)
//	// returns: 42, nil
func ParseStringWithDefault[T ParseStringSupportedTypes](rawValue string, dft T) (T, error) {
	parsedValue, err := ParseString[T](rawValue)
	if err != nil {
		return dft, err
	}
	return parsedValue, nil
}

// ParseStringOrDefault parses a string value into the specified type T.
// If parsing fails, it returns the provided default value and suppresses the error.
//
// This is the most convenient option when you want a fallback value
// and don't need to handle the error explicitly.
//
// Example:
//
//	num := ParseStringOrDefault("invalid", 10) // returns 10
//	num := ParseStringOrDefault("42", 10)      // returns 42
func ParseStringOrDefault[T ParseStringSupportedTypes](rawValue string, dft T) T {
	parsedValue, err := ParseString[T](rawValue)
	if err != nil {
		return dft
	}
	return parsedValue
}
