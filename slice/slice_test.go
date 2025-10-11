package slice

import (
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCycle(t *testing.T) {
	t.Run("cycles through items", func(t *testing.T) {
		next := Cycle(1, 2, 3)

		assert.Equal(t, 1, next())
		assert.Equal(t, 2, next())
		assert.Equal(t, 3, next())
		assert.Equal(t, 1, next()) // wraps back to first
		assert.Equal(t, 2, next())
		assert.Equal(t, 3, next())
		assert.Equal(t, 1, next()) // wraps again
	})

	t.Run("cycles through single item", func(t *testing.T) {
		next := Cycle(42)

		assert.Equal(t, 42, next())
		assert.Equal(t, 42, next())
		assert.Equal(t, 42, next())
	})

	t.Run("returns zero value for empty slice", func(t *testing.T) {
		next := Cycle[int]()

		assert.Equal(t, 0, next())
		assert.Equal(t, 0, next())
		assert.Equal(t, 0, next())
	})

	t.Run("works with string type", func(t *testing.T) {
		next := Cycle("a", "b", "c")

		assert.Equal(t, "a", next())
		assert.Equal(t, "b", next())
		assert.Equal(t, "c", next())
		assert.Equal(t, "a", next())
	})

	t.Run("works with struct type", func(t *testing.T) {
		type Point struct {
			X, Y int
		}
		next := Cycle(
			Point{1, 2},
			Point{3, 4},
		)

		assert.Equal(t, Point{1, 2}, next())
		assert.Equal(t, Point{3, 4}, next())
		assert.Equal(t, Point{1, 2}, next())
	})

	t.Run("works with pointer type", func(t *testing.T) {
		val1, val2 := 10, 20
		next := Cycle(&val1, &val2)

		assert.Equal(t, &val1, next())
		assert.Equal(t, &val2, next())
		assert.Equal(t, &val1, next())
	})

	t.Run("empty string slice returns empty string", func(t *testing.T) {
		next := Cycle[string]()

		assert.Empty(t, next())
		assert.Empty(t, next())
	})

	t.Run("multiple cycles work correctly", func(t *testing.T) {
		next := Cycle("x", "y")

		// First cycle
		assert.Equal(t, "x", next())
		assert.Equal(t, "y", next())
		// Second cycle
		assert.Equal(t, "x", next())
		assert.Equal(t, "y", next())
		// Third cycle
		assert.Equal(t, "x", next())
		assert.Equal(t, "y", next())
	})

	t.Run("each function instance maintains its own state", func(t *testing.T) {
		next1 := Cycle(1, 2, 3)
		next2 := Cycle(1, 2, 3)

		assert.Equal(t, 1, next1())
		assert.Equal(t, 1, next2())
		assert.Equal(t, 2, next1())
		assert.Equal(t, 2, next2())
		assert.Equal(t, 3, next1())
		assert.Equal(t, 1, next1())
		assert.Equal(t, 3, next2())
	})

	t.Run("works with boolean type", func(t *testing.T) {
		next := Cycle(true, false)

		assert.True(t, next())
		assert.False(t, next())
		assert.True(t, next())
		assert.False(t, next())
	})

	t.Run("works with float type", func(t *testing.T) {
		next := Cycle(1.1, 2.2, 3.3)

		assert.InDelta(t, 1.1, next(), 0.001)
		assert.InDelta(t, 2.2, next(), 0.001)
		assert.InDelta(t, 3.3, next(), 0.001)
		assert.InDelta(t, 1.1, next(), 0.001)
	})
}

func TestRandomCycle(t *testing.T) {
	t.Run("returns zero value for empty slice", func(t *testing.T) {
		next := RandomCycle[int](nil)

		assert.Equal(t, 0, next())
		assert.Equal(t, 0, next())
		assert.Equal(t, 0, next())
	})

	t.Run("returns same item for single item slice", func(t *testing.T) {
		next := RandomCycle(nil, 42)

		assert.Equal(t, 42, next())
		assert.Equal(t, 42, next())
		assert.Equal(t, 42, next())
	})

	t.Run("cycles through all items with seeded random", func(t *testing.T) {
		seed := rand.NewPCG(1, 2)
		r := rand.New(seed)
		next := RandomCycle(r, "a", "b", "c")

		// Collect items to verify all are returned
		seen := make(map[string]bool)
		for range 100 {
			item := next()
			seen[item] = true
		}

		// All items should have been seen
		assert.True(t, seen["a"], "should see 'a'")
		assert.True(t, seen["b"], "should see 'b'")
		assert.True(t, seen["c"], "should see 'c'")
	})

	t.Run("works with seeded random generator", func(t *testing.T) {
		// Using same seed should produce same sequence
		seed1 := rand.NewPCG(42, 42)
		r1 := rand.New(seed1)
		next1 := RandomCycle(r1, 1, 2, 3)

		seed2 := rand.NewPCG(42, 42)
		r2 := rand.New(seed2)
		next2 := RandomCycle(r2, 1, 2, 3)

		// First few calls should match
		for range 10 {
			assert.Equal(t, next1(), next2(), "calls with same seed should match")
		}
	})

	t.Run("creates own random generator when nil provided", func(t *testing.T) {
		next := RandomCycle[int](nil, 1, 2, 3)

		// Should not panic and should return valid values
		for range 10 {
			val := next()
			assert.Contains(t, []int{1, 2, 3}, val)
		}
	})

	t.Run("eventually returns all items in cycle", func(t *testing.T) {
		seed := rand.NewPCG(123, 456)
		r := rand.New(seed)
		next := RandomCycle(r, 1, 2, 3, 4, 5)

		// Track items seen in current "window"
		for range 5 {
			seen := make(map[int]int)
			// Collect enough items to ensure we see at least one full cycle
			for range 10 {
				item := next()
				seen[item]++
				assert.Contains(t, []int{1, 2, 3, 4, 5}, item)
			}
			// Should see multiple items
			assert.GreaterOrEqual(t, len(seen), 3, "should see at least 3 different items")
		}
	})

	t.Run("works with string type", func(t *testing.T) {
		seed := rand.NewPCG(1, 2)
		r := rand.New(seed)
		next := RandomCycle(r, "x", "y", "z")

		seen := make(map[string]bool)
		for range 50 {
			item := next()
			seen[item] = true
			assert.Contains(t, []string{"x", "y", "z"}, item)
		}

		assert.True(t, seen["x"])
		assert.True(t, seen["y"])
		assert.True(t, seen["z"])
	})

	t.Run("works with struct type", func(t *testing.T) {
		type Item struct {
			ID int
		}
		seed := rand.NewPCG(1, 2)
		r := rand.New(seed)
		next := RandomCycle(r, Item{1}, Item{2}, Item{3})

		seen := make(map[int]bool)
		for range 50 {
			item := next()
			seen[item.ID] = true
		}

		assert.True(t, seen[1])
		assert.True(t, seen[2])
		assert.True(t, seen[3])
	})

	t.Run("each function instance maintains its own state", func(t *testing.T) {
		seed1 := rand.NewPCG(1, 2)
		r1 := rand.New(seed1)
		next1 := RandomCycle(r1, 1, 2, 3)

		seed2 := rand.NewPCG(3, 5)
		r2 := rand.New(seed2)
		next2 := RandomCycle(r2, 1, 2, 3)

		// Get a few values from each
		vals1 := []int{next1(), next1(), next1()}
		vals2 := []int{next2(), next2(), next2()}

		// They should be different (very high probability with different seeds)
		assert.NotEqual(t, vals1, vals2, "different seeds should produce different sequences")
	})

	t.Run("with two items", func(t *testing.T) {
		seed := rand.NewPCG(99, 100)
		r := rand.New(seed)
		next := RandomCycle(r, "a", "b")

		seen := make(map[string]bool)
		for range 20 {
			item := next()
			seen[item] = true
			assert.Contains(t, []string{"a", "b"}, item)
		}

		assert.True(t, seen["a"])
		assert.True(t, seen["b"])
	})

	t.Run("deterministic with same seed across multiple cycles", func(t *testing.T) {
		// First run
		seed1 := rand.NewPCG(777, 888)
		r1 := rand.New(seed1)
		next1 := RandomCycle(r1, 10, 20, 30)
		sequence1 := make([]int, 20)
		for i := range 20 {
			sequence1[i] = next1()
		}

		// Second run with same seed
		seed2 := rand.NewPCG(777, 888)
		r2 := rand.New(seed2)
		next2 := RandomCycle(r2, 10, 20, 30)
		sequence2 := make([]int, 20)
		for i := range 20 {
			sequence2[i] = next2()
		}

		assert.Equal(t, sequence1, sequence2, "same seed should produce same sequence")
	})

	t.Run("empty string slice returns empty string", func(t *testing.T) {
		next := RandomCycle[string](nil)

		assert.Empty(t, next())
		assert.Empty(t, next())
	})

	t.Run("large number of items", func(t *testing.T) {
		items := make([]int, 100)
		for i := range 100 {
			items[i] = i
		}

		seed := rand.NewPCG(42, 43)
		r := rand.New(seed)
		next := RandomCycle(r, items...)

		seen := make(map[int]bool)
		// Call enough times to likely see all items
		for range 1000 {
			item := next()
			seen[item] = true
			assert.GreaterOrEqual(t, item, 0)
			assert.LessOrEqual(t, item, 99)
		}

		// Should see a good variety of items
		assert.Greater(t, len(seen), 90, "should see most items with 1000 calls")
	})
}
