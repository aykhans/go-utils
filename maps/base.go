package maps

import "maps"

// InitMap initializes a map pointer if it is nil.
// If the map is already initialized, this function does nothing.
//
// This is useful for ensuring a map is ready to use before adding entries,
// preventing nil pointer dereferences.
//
// Example:
//
//	var m map[string]int
//	InitMap(&m)
//	m["key"] = 42 // safe to use
func InitMap[K comparable, V any, T ~map[K]V](m *T) {
	if *m == nil {
		*m = make(T)
	}
}

// UpdateMap merges entries from the new map into the old map.
// If the old map is nil, it will be initialized first.
// Existing keys in the old map will be overwritten with values from the new map.
//
// This function modifies the old map in place by copying all key-value pairs
// from the new map into it.
//
// Example:
//
//	old := map[string]int{"a": 1, "b": 2}
//	new := map[string]int{"b": 3, "c": 4}
//	UpdateMap(&old, new)
//	// old is now: {"a": 1, "b": 3, "c": 4}
func UpdateMap[K comparable, V any, T ~map[K]V](oldMap *T, newMap T) {
	InitMap(oldMap)
	maps.Copy(*oldMap, newMap)
}
