package slice

import (
	"math/rand/v2"
	"time"
)

// Cycle returns a function that cycles through the provided items infinitely.
// Each call to the returned function returns the next item in sequence, wrapping
// back to the first item after the last one is returned.
//
// If no items are provided, the returned function will always return the zero
// value for type T.
//
// The returned function is not safe for concurrent use. If you need to call it
// from multiple goroutines, you must synchronize access with a mutex or similar.
//
// Example:
//
//	next := Cycle(1, 2, 3)
//	fmt.Println(next()) // 1
//	fmt.Println(next()) // 2
//	fmt.Println(next()) // 3
//	fmt.Println(next()) // 1
func Cycle[T any](items ...T) func() T {
	if len(items) == 0 {
		var zero T
		return func() T { return zero }
	}

	index := 0
	return func() T {
		item := items[index]
		index = (index + 1) % len(items)
		return item
	}
}

// RandomCycle returns a function that cycles through the provided items with
// randomization. It cycles through all items sequentially, but when it completes
// a full cycle, it randomly picks a new starting point for the next cycle.
//
// The localRand parameter can be used to provide a custom random number generator.
// If nil, a new generator will be created using the current time as the seed.
//
// The returned function is not safe for concurrent use. If you need to call it
// from multiple goroutines, you must synchronize access with a mutex or similar.
//
// Special cases:
//   - If no items are provided, the returned function always returns the zero value for type T.
//   - If only one item is provided, the returned function always returns that item.
//
// Example:
//
//	next := RandomCycle(nil, "a", "b", "c")
//	// Might produce: "b", "c", "a", "c", "a", "b", ...
//	// (cycles through all items, then starts from a random position)
func RandomCycle[T any](localRand *rand.Rand, items ...T) func() T {
	switch sliceLen := len(items); sliceLen {
	case 0:
		var zero T
		return func() T { return zero }
	case 1:
		return func() T { return items[0] }
	default:
		if localRand == nil {
			//nolint:gosec
			localRand = rand.New(
				rand.NewPCG(
					uint64(time.Now().UnixNano()),
					uint64(time.Now().UnixNano()>>32),
				),
			)
		}

		currentIndex := localRand.IntN(sliceLen)
		stopIndex := currentIndex
		return func() T {
			item := items[currentIndex]
			currentIndex++
			if currentIndex == sliceLen {
				currentIndex = 0
			}
			if currentIndex == stopIndex {
				currentIndex = localRand.IntN(sliceLen)
				stopIndex = currentIndex
			}

			return item
		}
	}
}
