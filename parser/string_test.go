package parser

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseString(t *testing.T) {
	t.Run("ParseString to string", func(t *testing.T) {
		tests := []struct {
			name     string
			input    string
			expected string
		}{
			{"empty string", "", ""},
			{"simple string", "hello", "hello"},
			{"string with spaces", "hello world", "hello world"},
			{"numeric string", "123", "123"},
			{"special characters", "!@#$%^&*()", "!@#$%^&*()"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := ParseString[string](test.input)
				require.NoError(t, err)
				assert.Equal(t, test.expected, result)
			})
		}
	})

	t.Run("ParseString to int", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			expected    int
			expectError bool
		}{
			{"positive int", "42", 42, false},
			{"negative int", "-42", -42, false},
			{"zero", "0", 0, false},
			{"invalid int", "abc", 0, true},
			{"float string", "3.14", 0, true},
			{"empty string", "", 0, true},
			{"overflow string", "99999999999999999999", 0, true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := ParseString[int](test.input)
				if test.expectError {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, test.expected, result)
				}
			})
		}
	})

	t.Run("ParseString to int8", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			expected    int8
			expectError bool
		}{
			{"valid int8", "127", 127, false},
			{"min int8", "-128", -128, false},
			{"overflow int8", "128", 0, true},
			{"underflow int8", "-129", 0, true},
			{"invalid", "abc", 0, true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := ParseString[int8](test.input)
				if test.expectError {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, test.expected, result)
				}
			})
		}
	})

	t.Run("ParseString to int16", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			expected    int16
			expectError bool
		}{
			{"valid int16", "32767", 32767, false},
			{"min int16", "-32768", -32768, false},
			{"overflow int16", "32768", 0, true},
			{"underflow int16", "-32769", 0, true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := ParseString[int16](test.input)
				if test.expectError {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, test.expected, result)
				}
			})
		}
	})

	t.Run("ParseString to int32", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			expected    int32
			expectError bool
		}{
			{"valid int32", "2147483647", 2147483647, false},
			{"min int32", "-2147483648", -2147483648, false},
			{"overflow int32", "2147483648", 0, true},
			{"underflow int32", "-2147483649", 0, true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := ParseString[int32](test.input)
				if test.expectError {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, test.expected, result)
				}
			})
		}
	})

	t.Run("ParseString to int64", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			expected    int64
			expectError bool
		}{
			{"valid int64", "9223372036854775807", 9223372036854775807, false},
			{"min int64", "-9223372036854775808", -9223372036854775808, false},
			{"large number", "123456789012345", 123456789012345, false},
			{"invalid", "not a number", 0, true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := ParseString[int64](test.input)
				if test.expectError {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, test.expected, result)
				}
			})
		}
	})

	t.Run("ParseString to uint", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			expected    uint
			expectError bool
		}{
			{"valid uint", "42", 42, false},
			{"zero", "0", 0, false},
			{"large uint", "4294967295", 4294967295, false},
			{"negative", "-1", 0, true},
			{"invalid", "abc", 0, true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := ParseString[uint](test.input)
				if test.expectError {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, test.expected, result)
				}
			})
		}
	})

	t.Run("ParseString to uint8", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			expected    uint8
			expectError bool
		}{
			{"valid uint8", "255", 255, false},
			{"min uint8", "0", 0, false},
			{"overflow uint8", "256", 0, true},
			{"negative", "-1", 0, true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := ParseString[uint8](test.input)
				if test.expectError {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, test.expected, result)
				}
			})
		}
	})

	t.Run("ParseString to uint16", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			expected    uint16
			expectError bool
		}{
			{"valid uint16", "65535", 65535, false},
			{"min uint16", "0", 0, false},
			{"overflow uint16", "65536", 0, true},
			{"negative", "-1", 0, true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := ParseString[uint16](test.input)
				if test.expectError {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, test.expected, result)
				}
			})
		}
	})

	t.Run("ParseString to uint32", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			expected    uint32
			expectError bool
		}{
			{"valid uint32", "4294967295", 4294967295, false},
			{"min uint32", "0", 0, false},
			{"overflow uint32", "4294967296", 0, true},
			{"negative", "-1", 0, true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := ParseString[uint32](test.input)
				if test.expectError {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, test.expected, result)
				}
			})
		}
	})

	t.Run("ParseString to uint64", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			expected    uint64
			expectError bool
		}{
			{"valid uint64", "18446744073709551615", 18446744073709551615, false},
			{"min uint64", "0", 0, false},
			{"large number", "123456789012345", 123456789012345, false},
			{"negative", "-1", 0, true},
			{"invalid", "not a number", 0, true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := ParseString[uint64](test.input)
				if test.expectError {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, test.expected, result)
				}
			})
		}
	})

	t.Run("ParseString to float64", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			expected    float64
			expectError bool
		}{
			{"integer", "42", 42.0, false},
			{"decimal", "3.14159", 3.14159, false},
			{"negative", "-2.5", -2.5, false},
			{"scientific notation", "1.23e10", 1.23e10, false},
			{"zero", "0", 0.0, false},
			{"invalid", "not a number", 0, true},
			{"empty", "", 0, true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := ParseString[float64](test.input)
				if test.expectError {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.InDelta(t, test.expected, result, 0.0001)
				}
			})
		}
	})

	t.Run("ParseString to bool", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			expected    bool
			expectError bool
		}{
			{"true lowercase", "true", true, false},
			{"True mixed case", "True", true, false},
			{"TRUE uppercase", "TRUE", true, false},
			{"1 as true", "1", true, false},
			{"false lowercase", "false", false, false},
			{"False mixed case", "False", false, false},
			{"FALSE uppercase", "FALSE", false, false},
			{"0 as false", "0", false, false},
			{"invalid", "yes", false, true},
			{"empty", "", false, true},
			{"numeric non-binary", "2", false, true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := ParseString[bool](test.input)
				if test.expectError {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, test.expected, result)
				}
			})
		}
	})

	t.Run("ParseString to time.Duration", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			expected    time.Duration
			expectError bool
		}{
			{"seconds", "10s", 10 * time.Second, false},
			{"minutes", "5m", 5 * time.Minute, false},
			{"hours", "2h", 2 * time.Hour, false},
			{"combined", "1h30m45s", time.Hour + 30*time.Minute + 45*time.Second, false},
			{"milliseconds", "500ms", 500 * time.Millisecond, false},
			{"microseconds", "100us", 100 * time.Microsecond, false},
			{"nanoseconds", "50ns", 50 * time.Nanosecond, false},
			{"negative", "-5s", -5 * time.Second, false},
			{"invalid", "5x", 0, true},
			{"empty", "", 0, true},
			{"no unit", "100", 0, true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := ParseString[time.Duration](test.input)
				if test.expectError {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, test.expected, result)
				}
			})
		}
	})

	t.Run("ParseString to url.URL", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			checkFunc   func(t *testing.T, u url.URL)
			expectError bool
		}{
			{
				name:  "http URL",
				input: "http://example.com",
				checkFunc: func(t *testing.T, u url.URL) {
					t.Helper()
					assert.Equal(t, "http", u.Scheme)
					assert.Equal(t, "example.com", u.Host)
				},
				expectError: false,
			},
			{
				name:  "https URL with path",
				input: "https://example.com/path/to/resource",
				checkFunc: func(t *testing.T, u url.URL) {
					t.Helper()
					assert.Equal(t, "https", u.Scheme)
					assert.Equal(t, "example.com", u.Host)
					assert.Equal(t, "/path/to/resource", u.Path)
				},
				expectError: false,
			},
			{
				name:  "URL with query parameters",
				input: "https://example.com/search?q=test&page=1",
				checkFunc: func(t *testing.T, u url.URL) {
					t.Helper()
					assert.Equal(t, "https", u.Scheme)
					assert.Equal(t, "example.com", u.Host)
					assert.Equal(t, "/search", u.Path)
					assert.Equal(t, "q=test&page=1", u.RawQuery)
				},
				expectError: false,
			},
			{
				name:  "URL with port",
				input: "http://localhost:8080/api",
				checkFunc: func(t *testing.T, u url.URL) {
					t.Helper()
					assert.Equal(t, "http", u.Scheme)
					assert.Equal(t, "localhost:8080", u.Host)
					assert.Equal(t, "/api", u.Path)
				},
				expectError: false,
			},
			{
				name:  "URL with fragment",
				input: "https://example.com/page#section",
				checkFunc: func(t *testing.T, u url.URL) {
					t.Helper()
					assert.Equal(t, "https", u.Scheme)
					assert.Equal(t, "example.com", u.Host)
					assert.Equal(t, "/page", u.Path)
					assert.Equal(t, "section", u.Fragment)
				},
				expectError: false,
			},
			{
				name:  "relative path",
				input: "/path/to/resource",
				checkFunc: func(t *testing.T, u url.URL) {
					t.Helper()
					assert.Empty(t, u.Scheme)
					assert.Empty(t, u.Host)
					assert.Equal(t, "/path/to/resource", u.Path)
				},
				expectError: false,
			},
			{
				name:  "empty string",
				input: "",
				checkFunc: func(t *testing.T, u url.URL) {
					t.Helper()
					assert.Empty(t, u.String())
				},
				expectError: false,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result, err := ParseString[url.URL](test.input)
				if test.expectError {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					if test.checkFunc != nil {
						test.checkFunc(t, result)
					}
				}
			})
		}
	})

	t.Run("Edge cases", func(t *testing.T) {
		t.Run("whitespace handling for numeric types", func(t *testing.T) {
			result, err := ParseString[int]("  42  ")
			require.Error(t, err)
			assert.Equal(t, 0, result)
		})

		t.Run("leading zeros for int", func(t *testing.T) {
			result, err := ParseString[int]("007")
			require.NoError(t, err)
			assert.Equal(t, 7, result)
		})

		t.Run("plus sign for positive numbers", func(t *testing.T) {
			result, err := ParseString[int]("+42")
			require.NoError(t, err)
			assert.Equal(t, 42, result)
		})

		t.Run("case sensitivity for bool", func(t *testing.T) {
			testCases := []string{"t", "T", "f", "F"}
			for _, tc := range testCases {
				result, err := ParseString[bool](tc)
				require.NoError(t, err)
				if tc == "t" || tc == "T" {
					assert.True(t, result)
				} else {
					assert.False(t, result)
				}
			}
		})
	})
}
