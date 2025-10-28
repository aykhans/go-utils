package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumLen(t *testing.T) {
	t.Run("NumLen with zero", func(t *testing.T) {
		assert.Equal(t, 1, NumLen(0))
		assert.Equal(t, int8(1), NumLen(int8(0)))
		assert.Equal(t, int16(1), NumLen(int16(0)))
		assert.Equal(t, int32(1), NumLen(int32(0)))
		assert.Equal(t, int64(1), NumLen(int64(0)))
		assert.Equal(t, uint(1), NumLen(uint(0)))
		assert.Equal(t, uint8(1), NumLen(uint8(0)))
		assert.Equal(t, uint16(1), NumLen(uint16(0)))
		assert.Equal(t, uint32(1), NumLen(uint32(0)))
		assert.Equal(t, uint64(1), NumLen(uint64(0)))
	})

	t.Run("NumLen with single digit positive numbers", func(t *testing.T) {
		tests := []struct {
			name     string
			input    int
			expected int
		}{
			{"one", 1, 1},
			{"five", 5, 1},
			{"nine", 9, 1},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := NumLen(test.input)
				assert.Equal(t, test.expected, result)
			})
		}
	})

	t.Run("NumLen with single digit negative numbers", func(t *testing.T) {
		tests := []struct {
			name     string
			input    int
			expected int
		}{
			{"negative one", -1, 1},
			{"negative five", -5, 1},
			{"negative nine", -9, 1},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := NumLen(test.input)
				assert.Equal(t, test.expected, result)
			})
		}
	})

	t.Run("NumLen with multi-digit positive numbers", func(t *testing.T) {
		tests := []struct {
			name     string
			input    int
			expected int
		}{
			{"two digits", 42, 2},
			{"three digits", 123, 3},
			{"four digits", 9999, 4},
			{"five digits", 12345, 5},
			{"six digits", 100000, 6},
			{"seven digits", 9876543, 7},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := NumLen(test.input)
				assert.Equal(t, test.expected, result)
			})
		}
	})

	t.Run("NumLen with multi-digit negative numbers", func(t *testing.T) {
		tests := []struct {
			name     string
			input    int
			expected int
		}{
			{"negative two digits", -42, 2},
			{"negative three digits", -123, 3},
			{"negative four digits", -9999, 4},
			{"negative five digits", -12345, 5},
			{"negative six digits", -100000, 6},
			{"negative seven digits", -9876543, 7},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := NumLen(test.input)
				assert.Equal(t, test.expected, result)
			})
		}
	})

	t.Run("NumLen with int8", func(t *testing.T) {
		tests := []struct {
			name     string
			input    int8
			expected int8
		}{
			{"zero", 0, 1},
			{"positive single", 9, 1},
			{"positive double", 99, 2},
			{"positive triple", 127, 3},
			{"negative single", -9, 1},
			{"negative double", -99, 2},
			{"negative triple", -128, 3},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := NumLen(test.input)
				assert.Equal(t, test.expected, result)
			})
		}
	})

	t.Run("NumLen with int16", func(t *testing.T) {
		tests := []struct {
			name     string
			input    int16
			expected int16
		}{
			{"zero", 0, 1},
			{"positive single", 5, 1},
			{"positive double", 42, 2},
			{"positive triple", 999, 3},
			{"positive max", 32767, 5},
			{"negative single", -5, 1},
			{"negative double", -42, 2},
			{"negative triple", -999, 3},
			{"negative min", -32768, 5},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := NumLen(test.input)
				assert.Equal(t, test.expected, result)
			})
		}
	})

	t.Run("NumLen with int32", func(t *testing.T) {
		tests := []struct {
			name     string
			input    int32
			expected int32
		}{
			{"zero", 0, 1},
			{"positive single", 7, 1},
			{"positive five", 12345, 5},
			{"positive max", 2147483647, 10},
			{"negative single", -7, 1},
			{"negative five", -12345, 5},
			{"negative min", -2147483648, 10},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := NumLen(test.input)
				assert.Equal(t, test.expected, result)
			})
		}
	})

	t.Run("NumLen with int64", func(t *testing.T) {
		tests := []struct {
			name     string
			input    int64
			expected int64
		}{
			{"zero", 0, 1},
			{"positive single", 3, 1},
			{"positive ten", 1234567890, 10},
			{"positive max", 9223372036854775807, 19},
			{"negative single", -3, 1},
			{"negative ten", -1234567890, 10},
			{"negative min", -9223372036854775808, 19},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := NumLen(test.input)
				assert.Equal(t, test.expected, result)
			})
		}
	})

	t.Run("NumLen with uint", func(t *testing.T) {
		tests := []struct {
			name     string
			input    uint
			expected uint
		}{
			{"zero", 0, 1},
			{"single digit", 5, 1},
			{"double digit", 42, 2},
			{"triple digit", 999, 3},
			{"large number", 123456789, 9},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := NumLen(test.input)
				assert.Equal(t, test.expected, result)
			})
		}
	})

	t.Run("NumLen with uint8", func(t *testing.T) {
		tests := []struct {
			name     string
			input    uint8
			expected uint8
		}{
			{"zero", 0, 1},
			{"single digit", 9, 1},
			{"double digit", 99, 2},
			{"max value", 255, 3},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := NumLen(test.input)
				assert.Equal(t, test.expected, result)
			})
		}
	})

	t.Run("NumLen with uint16", func(t *testing.T) {
		tests := []struct {
			name     string
			input    uint16
			expected uint16
		}{
			{"zero", 0, 1},
			{"single digit", 8, 1},
			{"double digit", 50, 2},
			{"triple digit", 500, 3},
			{"max value", 65535, 5},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := NumLen(test.input)
				assert.Equal(t, test.expected, result)
			})
		}
	})

	t.Run("NumLen with uint32", func(t *testing.T) {
		tests := []struct {
			name     string
			input    uint32
			expected uint32
		}{
			{"zero", 0, 1},
			{"single digit", 6, 1},
			{"five digits", 54321, 5},
			{"max value", 4294967295, 10},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := NumLen(test.input)
				assert.Equal(t, test.expected, result)
			})
		}
	})

	t.Run("NumLen with uint64", func(t *testing.T) {
		tests := []struct {
			name     string
			input    uint64
			expected uint64
		}{
			{"zero", 0, 1},
			{"single digit", 4, 1},
			{"ten digits", 9876543210, 10},
			{"max value", 18446744073709551615, 20},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := NumLen(test.input)
				assert.Equal(t, test.expected, result)
			})
		}
	})

	t.Run("NumLen with boundary values", func(t *testing.T) {
		t.Run("powers of 10", func(t *testing.T) {
			assert.Equal(t, 1, NumLen(1))
			assert.Equal(t, 2, NumLen(10))
			assert.Equal(t, 3, NumLen(100))
			assert.Equal(t, 4, NumLen(1000))
			assert.Equal(t, 5, NumLen(10000))
			assert.Equal(t, 6, NumLen(100000))
			assert.Equal(t, 7, NumLen(1000000))
		})

		t.Run("one less than powers of 10", func(t *testing.T) {
			assert.Equal(t, 1, NumLen(9))
			assert.Equal(t, 2, NumLen(99))
			assert.Equal(t, 3, NumLen(999))
			assert.Equal(t, 4, NumLen(9999))
			assert.Equal(t, 5, NumLen(99999))
			assert.Equal(t, 6, NumLen(999999))
		})

		t.Run("negative powers of 10", func(t *testing.T) {
			assert.Equal(t, 1, NumLen(-1))
			assert.Equal(t, 2, NumLen(-10))
			assert.Equal(t, 3, NumLen(-100))
			assert.Equal(t, 4, NumLen(-1000))
			assert.Equal(t, 5, NumLen(-10000))
		})
	})

	t.Run("NumLen consistency across types", func(t *testing.T) {
		// Same value, different types should return same digit count
		assert.Equal(t, 2, NumLen(42))
		assert.Equal(t, int8(2), NumLen(int8(42)))
		assert.Equal(t, int16(2), NumLen(int16(42)))
		assert.Equal(t, int32(2), NumLen(int32(42)))
		assert.Equal(t, int64(2), NumLen(int64(42)))
		assert.Equal(t, uint(2), NumLen(uint(42)))
		assert.Equal(t, uint8(2), NumLen(uint8(42)))
		assert.Equal(t, uint16(2), NumLen(uint16(42)))
		assert.Equal(t, uint32(2), NumLen(uint32(42)))
		assert.Equal(t, uint64(2), NumLen(uint64(42)))
	})
}
