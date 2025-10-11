package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitMap(t *testing.T) {
	t.Run("initializes nil map", func(t *testing.T) {
		var m map[string]int
		require.Nil(t, m)

		InitMap(&m)

		assert.NotNil(t, m)
		assert.Empty(t, m)
		assert.Empty(t, m)
	})

	t.Run("does not reinitialize existing map", func(t *testing.T) {
		m := map[string]int{"key": 42}
		originalLen := len(m)

		InitMap(&m)

		assert.Len(t, m, originalLen)
		assert.Equal(t, 42, m["key"])
	})

	t.Run("allows adding entries after initialization", func(t *testing.T) {
		var m map[string]int
		InitMap(&m)

		m["test"] = 100
		assert.Equal(t, 100, m["test"])
		assert.Len(t, m, 1)
	})

	t.Run("works with empty initialized map", func(t *testing.T) {
		m := make(map[string]int)
		InitMap(&m)

		assert.NotNil(t, m)
		assert.Empty(t, m)
	})

	t.Run("works with different key types", func(t *testing.T) {
		t.Run("int keys", func(t *testing.T) {
			var m map[int]string
			InitMap(&m)

			assert.NotNil(t, m)
			m[1] = "one"
			assert.Equal(t, "one", m[1])
		})

		t.Run("struct keys", func(t *testing.T) {
			type Key struct {
				ID   int
				Name string
			}
			var m map[Key]bool
			InitMap(&m)

			assert.NotNil(t, m)
			k := Key{ID: 1, Name: "test"}
			m[k] = true
			assert.True(t, m[k])
		})
	})

	t.Run("works with different value types", func(t *testing.T) {
		t.Run("string values", func(t *testing.T) {
			var m map[string]string
			InitMap(&m)

			assert.NotNil(t, m)
			m["key"] = "value"
			assert.Equal(t, "value", m["key"])
		})

		t.Run("struct values", func(t *testing.T) {
			type Value struct {
				Count int
				Name  string
			}
			var m map[string]Value
			InitMap(&m)

			assert.NotNil(t, m)
			m["test"] = Value{Count: 5, Name: "foo"}
			assert.Equal(t, 5, m["test"].Count)
		})

		t.Run("slice values", func(t *testing.T) {
			var m map[string][]int
			InitMap(&m)

			assert.NotNil(t, m)
			m["numbers"] = []int{1, 2, 3}
			assert.Equal(t, []int{1, 2, 3}, m["numbers"])
		})

		t.Run("pointer values", func(t *testing.T) {
			var m map[string]*int
			InitMap(&m)

			assert.NotNil(t, m)
			val := 42
			m["ptr"] = &val
			assert.Equal(t, 42, *m["ptr"])
		})
	})

	t.Run("works with custom map types", func(t *testing.T) {
		type CustomMap map[string]int
		var m CustomMap
		InitMap(&m)

		assert.NotNil(t, m)
		m["custom"] = 99
		assert.Equal(t, 99, m["custom"])
	})
}

func TestUpdateMap(t *testing.T) {
	t.Run("merges new entries into existing map", func(t *testing.T) {
		oldMap := map[string]int{"a": 1, "b": 2}
		newMap := map[string]int{"c": 3, "d": 4}

		UpdateMap(&oldMap, newMap)

		assert.Len(t, oldMap, 4)
		assert.Equal(t, 1, oldMap["a"])
		assert.Equal(t, 2, oldMap["b"])
		assert.Equal(t, 3, oldMap["c"])
		assert.Equal(t, 4, oldMap["d"])
	})

	t.Run("overwrites existing keys", func(t *testing.T) {
		oldMap := map[string]int{"a": 1, "b": 2}
		newMap := map[string]int{"b": 3, "c": 4}

		UpdateMap(&oldMap, newMap)

		assert.Len(t, oldMap, 3)
		assert.Equal(t, 1, oldMap["a"])
		assert.Equal(t, 3, oldMap["b"], "existing key 'b' should be overwritten")
		assert.Equal(t, 4, oldMap["c"])
	})

	t.Run("initializes nil map before merging", func(t *testing.T) {
		var oldMap map[string]int
		newMap := map[string]int{"a": 1, "b": 2}

		require.Nil(t, oldMap)
		UpdateMap(&oldMap, newMap)

		assert.NotNil(t, oldMap)
		assert.Len(t, oldMap, 2)
		assert.Equal(t, 1, oldMap["a"])
		assert.Equal(t, 2, oldMap["b"])
	})

	t.Run("handles empty new map", func(t *testing.T) {
		oldMap := map[string]int{"a": 1}
		newMap := map[string]int{}

		UpdateMap(&oldMap, newMap)

		assert.Len(t, oldMap, 1)
		assert.Equal(t, 1, oldMap["a"])
	})

	t.Run("handles empty old map", func(t *testing.T) {
		oldMap := map[string]int{}
		newMap := map[string]int{"a": 1, "b": 2}

		UpdateMap(&oldMap, newMap)

		assert.Len(t, oldMap, 2)
		assert.Equal(t, 1, oldMap["a"])
		assert.Equal(t, 2, oldMap["b"])
	})

	t.Run("handles both empty maps", func(t *testing.T) {
		oldMap := map[string]int{}
		newMap := map[string]int{}

		UpdateMap(&oldMap, newMap)

		assert.NotNil(t, oldMap)
		assert.Empty(t, oldMap)
	})

	t.Run("works with different types", func(t *testing.T) {
		t.Run("string values", func(t *testing.T) {
			oldMap := map[string]string{"key1": "value1"}
			newMap := map[string]string{"key2": "value2"}

			UpdateMap(&oldMap, newMap)

			assert.Len(t, oldMap, 2)
			assert.Equal(t, "value1", oldMap["key1"])
			assert.Equal(t, "value2", oldMap["key2"])
		})

		t.Run("struct values", func(t *testing.T) {
			type Person struct {
				Name string
				Age  int
			}
			oldMap := map[string]Person{
				"john": {Name: "John", Age: 30},
			}
			newMap := map[string]Person{
				"jane": {Name: "Jane", Age: 25},
			}

			UpdateMap(&oldMap, newMap)

			assert.Len(t, oldMap, 2)
			assert.Equal(t, "John", oldMap["john"].Name)
			assert.Equal(t, "Jane", oldMap["jane"].Name)
		})

		t.Run("int keys", func(t *testing.T) {
			oldMap := map[int]string{1: "one"}
			newMap := map[int]string{2: "two"}

			UpdateMap(&oldMap, newMap)

			assert.Len(t, oldMap, 2)
			assert.Equal(t, "one", oldMap[1])
			assert.Equal(t, "two", oldMap[2])
		})
	})

	t.Run("modifies map in place", func(t *testing.T) {
		oldMap := map[string]int{"a": 1}
		originalMap := oldMap
		newMap := map[string]int{"b": 2}

		UpdateMap(&oldMap, newMap)

		// Verify it's the same map (modified in place)
		assert.Equal(t, 2, oldMap["b"])
		// Note: originalMap is pointing to the same underlying map
		// so it should also have the new value
		assert.Equal(t, 2, originalMap["b"])
	})

	t.Run("does not modify new map", func(t *testing.T) {
		oldMap := map[string]int{"a": 1}
		newMap := map[string]int{"b": 2, "c": 3}
		newLen := len(newMap)

		UpdateMap(&oldMap, newMap)

		assert.Len(t, newMap, newLen, "new map should not be modified")
		assert.Equal(t, 2, newMap["b"])
		assert.Equal(t, 3, newMap["c"])
		_, hasA := newMap["a"]
		assert.False(t, hasA, "new map should not contain keys from old map")
	})

	t.Run("works with custom map types", func(t *testing.T) {
		type CustomMap map[string]int
		oldMap := CustomMap{"a": 1}
		newMap := CustomMap{"b": 2}

		UpdateMap(&oldMap, newMap)

		assert.Len(t, oldMap, 2)
		assert.Equal(t, 1, oldMap["a"])
		assert.Equal(t, 2, oldMap["b"])
	})

	t.Run("overwrites with zero values", func(t *testing.T) {
		oldMap := map[string]int{"a": 10, "b": 20}
		newMap := map[string]int{"a": 0}

		UpdateMap(&oldMap, newMap)

		assert.Len(t, oldMap, 2)
		assert.Equal(t, 0, oldMap["a"], "should overwrite with zero value")
		assert.Equal(t, 20, oldMap["b"])
	})

	t.Run("complex merge scenario", func(t *testing.T) {
		oldMap := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}
		newMap := map[string]int{
			"b": 20, // overwrite
			"c": 30, // overwrite
			"d": 40, // new
			"e": 50, // new
		}

		UpdateMap(&oldMap, newMap)

		expected := map[string]int{
			"a": 1,  // preserved
			"b": 20, // overwritten
			"c": 30, // overwritten
			"d": 40, // new
			"e": 50, // new
		}

		assert.Equal(t, expected, oldMap)
	})
}
