package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNilOrZero(t *testing.T) {
	t.Run("nil pointer returns true", func(t *testing.T) {
		var intPtr *int
		assert.True(t, IsNilOrZero(intPtr))

		var strPtr *string
		assert.True(t, IsNilOrZero(strPtr))

		var boolPtr *bool
		assert.True(t, IsNilOrZero(boolPtr))

		var floatPtr *float64
		assert.True(t, IsNilOrZero(floatPtr))
	})

	t.Run("pointer to zero value returns true", func(t *testing.T) {
		t.Run("int zero", func(t *testing.T) {
			val := 0
			assert.True(t, IsNilOrZero(&val))
		})

		t.Run("string zero", func(t *testing.T) {
			val := ""
			assert.True(t, IsNilOrZero(&val))
		})

		t.Run("bool zero", func(t *testing.T) {
			val := false
			assert.True(t, IsNilOrZero(&val))
		})

		t.Run("float64 zero", func(t *testing.T) {
			val := 0.0
			assert.True(t, IsNilOrZero(&val))
		})

		t.Run("uint zero", func(t *testing.T) {
			val := uint(0)
			assert.True(t, IsNilOrZero(&val))
		})
	})

	t.Run("pointer to non-zero value returns false", func(t *testing.T) {
		t.Run("int non-zero", func(t *testing.T) {
			val := 42
			assert.False(t, IsNilOrZero(&val))

			negVal := -1
			assert.False(t, IsNilOrZero(&negVal))
		})

		t.Run("string non-zero", func(t *testing.T) {
			val := "hello"
			assert.False(t, IsNilOrZero(&val))

			spaceVal := " "
			assert.False(t, IsNilOrZero(&spaceVal))
		})

		t.Run("bool non-zero", func(t *testing.T) {
			val := true
			assert.False(t, IsNilOrZero(&val))
		})

		t.Run("float64 non-zero", func(t *testing.T) {
			val := 3.14
			assert.False(t, IsNilOrZero(&val))

			negVal := -2.5
			assert.False(t, IsNilOrZero(&negVal))
		})

		t.Run("uint non-zero", func(t *testing.T) {
			val := uint(100)
			assert.False(t, IsNilOrZero(&val))
		})
	})

	t.Run("works with all integer types", func(t *testing.T) {
		t.Run("int8", func(t *testing.T) {
			var nilPtr *int8
			assert.True(t, IsNilOrZero(nilPtr))

			zero := int8(0)
			assert.True(t, IsNilOrZero(&zero))

			nonZero := int8(127)
			assert.False(t, IsNilOrZero(&nonZero))
		})

		t.Run("int16", func(t *testing.T) {
			var nilPtr *int16
			assert.True(t, IsNilOrZero(nilPtr))

			zero := int16(0)
			assert.True(t, IsNilOrZero(&zero))

			nonZero := int16(1000)
			assert.False(t, IsNilOrZero(&nonZero))
		})

		t.Run("int32", func(t *testing.T) {
			var nilPtr *int32
			assert.True(t, IsNilOrZero(nilPtr))

			zero := int32(0)
			assert.True(t, IsNilOrZero(&zero))

			nonZero := int32(100000)
			assert.False(t, IsNilOrZero(&nonZero))
		})

		t.Run("int64", func(t *testing.T) {
			var nilPtr *int64
			assert.True(t, IsNilOrZero(nilPtr))

			zero := int64(0)
			assert.True(t, IsNilOrZero(&zero))

			nonZero := int64(9223372036854775807)
			assert.False(t, IsNilOrZero(&nonZero))
		})

		t.Run("uint8", func(t *testing.T) {
			var nilPtr *uint8
			assert.True(t, IsNilOrZero(nilPtr))

			zero := uint8(0)
			assert.True(t, IsNilOrZero(&zero))

			nonZero := uint8(255)
			assert.False(t, IsNilOrZero(&nonZero))
		})

		t.Run("uint16", func(t *testing.T) {
			var nilPtr *uint16
			assert.True(t, IsNilOrZero(nilPtr))

			zero := uint16(0)
			assert.True(t, IsNilOrZero(&zero))

			nonZero := uint16(65535)
			assert.False(t, IsNilOrZero(&nonZero))
		})

		t.Run("uint32", func(t *testing.T) {
			var nilPtr *uint32
			assert.True(t, IsNilOrZero(nilPtr))

			zero := uint32(0)
			assert.True(t, IsNilOrZero(&zero))

			nonZero := uint32(4294967295)
			assert.False(t, IsNilOrZero(&nonZero))
		})

		t.Run("uint64", func(t *testing.T) {
			var nilPtr *uint64
			assert.True(t, IsNilOrZero(nilPtr))

			zero := uint64(0)
			assert.True(t, IsNilOrZero(&zero))

			nonZero := uint64(18446744073709551615)
			assert.False(t, IsNilOrZero(&nonZero))
		})
	})

	t.Run("works with float types", func(t *testing.T) {
		t.Run("float32", func(t *testing.T) {
			var nilPtr *float32
			assert.True(t, IsNilOrZero(nilPtr))

			zero := float32(0.0)
			assert.True(t, IsNilOrZero(&zero))

			nonZero := float32(1.5)
			assert.False(t, IsNilOrZero(&nonZero))
		})

		t.Run("float64", func(t *testing.T) {
			var nilPtr *float64
			assert.True(t, IsNilOrZero(nilPtr))

			zero := float64(0.0)
			assert.True(t, IsNilOrZero(&zero))

			nonZero := float64(123.456)
			assert.False(t, IsNilOrZero(&nonZero))
		})
	})

	t.Run("works with struct types", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		t.Run("nil struct pointer", func(t *testing.T) {
			var nilPtr *Person
			assert.True(t, IsNilOrZero(nilPtr))
		})

		t.Run("pointer to zero struct", func(t *testing.T) {
			zero := Person{}
			assert.True(t, IsNilOrZero(&zero))
		})

		t.Run("pointer to non-zero struct", func(t *testing.T) {
			nonZero := Person{Name: "John", Age: 30}
			assert.False(t, IsNilOrZero(&nonZero))

			partiallyFilled := Person{Name: "Jane"}
			assert.False(t, IsNilOrZero(&partiallyFilled))

			onlyAge := Person{Age: 25}
			assert.False(t, IsNilOrZero(&onlyAge))
		})
	})

	t.Run("works with array types", func(t *testing.T) {
		t.Run("nil array pointer", func(t *testing.T) {
			var nilPtr *[3]int
			assert.True(t, IsNilOrZero(nilPtr))
		})

		t.Run("pointer to zero array", func(t *testing.T) {
			zero := [3]int{0, 0, 0}
			assert.True(t, IsNilOrZero(&zero))
		})

		t.Run("pointer to non-zero array", func(t *testing.T) {
			nonZero := [3]int{1, 2, 3}
			assert.False(t, IsNilOrZero(&nonZero))

			partiallyFilled := [3]int{1, 0, 0}
			assert.False(t, IsNilOrZero(&partiallyFilled))
		})
	})

	t.Run("edge cases", func(t *testing.T) {
		t.Run("pointer to negative number", func(t *testing.T) {
			val := -1
			assert.False(t, IsNilOrZero(&val))
		})

		t.Run("pointer to whitespace string", func(t *testing.T) {
			val := " "
			assert.False(t, IsNilOrZero(&val))

			tab := "\t"
			assert.False(t, IsNilOrZero(&tab))

			newline := "\n"
			assert.False(t, IsNilOrZero(&newline))
		})

		t.Run("pointer to byte (uint8) zero", func(t *testing.T) {
			val := byte(0)
			assert.True(t, IsNilOrZero(&val))

			nonZero := byte('A')
			assert.False(t, IsNilOrZero(&nonZero))
		})

		t.Run("pointer to rune (int32) zero", func(t *testing.T) {
			val := rune(0)
			assert.True(t, IsNilOrZero(&val))

			nonZero := rune('Z')
			assert.False(t, IsNilOrZero(&nonZero))
		})
	})

	t.Run("consistency with ToPtr", func(t *testing.T) {
		t.Run("ToPtr of zero value is nil or zero", func(t *testing.T) {
			intPtr := ToPtr(0)
			assert.True(t, IsNilOrZero(intPtr))

			strPtr := ToPtr("")
			assert.True(t, IsNilOrZero(strPtr))

			boolPtr := ToPtr(false)
			assert.True(t, IsNilOrZero(boolPtr))
		})

		t.Run("ToPtr of non-zero value is not nil or zero", func(t *testing.T) {
			intPtr := ToPtr(42)
			assert.False(t, IsNilOrZero(intPtr))

			strPtr := ToPtr("hello")
			assert.False(t, IsNilOrZero(strPtr))

			boolPtr := ToPtr(true)
			assert.False(t, IsNilOrZero(boolPtr))
		})
	})

	t.Run("real-world usage scenarios", func(t *testing.T) {
		t.Run("optional configuration value", func(t *testing.T) {
			type Config struct {
				Port    *int
				Timeout *int
			}

			// No port configured (nil)
			config1 := Config{}
			assert.True(t, IsNilOrZero(config1.Port))

			// Port explicitly set to 0
			zero := 0
			config2 := Config{Port: &zero}
			assert.True(t, IsNilOrZero(config2.Port))

			// Port configured to 8080
			port := 8080
			config3 := Config{Port: &port}
			assert.False(t, IsNilOrZero(config3.Port))
		})

		t.Run("optional string field", func(t *testing.T) {
			type User struct {
				Name     string
				Nickname *string
			}

			// No nickname provided
			user1 := User{Name: "John"}
			assert.True(t, IsNilOrZero(user1.Nickname))

			// Nickname explicitly empty
			emptyNick := ""
			user2 := User{Name: "Jane", Nickname: &emptyNick}
			assert.True(t, IsNilOrZero(user2.Nickname))

			// Nickname provided
			nick := "Johnny"
			user3 := User{Name: "John", Nickname: &nick}
			assert.False(t, IsNilOrZero(user3.Nickname))
		})
	})
}
