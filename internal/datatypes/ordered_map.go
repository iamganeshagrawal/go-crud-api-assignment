package datatypes

// OrderedMap is a data structure that maintains the order of keys
// while providing O(1) access to values.
//
// This implementation is non-thread-safe.
//
// The OrderedMap is parameterized by the key type K and the value type T.
// It is implemented using a slice of keys and a map of keys to values.
// The slice maintains the order of keys, while the map provides O(1) access to values.
type OrderedMap[K comparable, T any] struct {
	keys   []K
	values map[K]T
}

func NewOrderedMap[K comparable, T any]() *OrderedMap[K, T] {
	return &OrderedMap[K, T]{
		keys:   make([]K, 0),
		values: make(map[K]T),
	}
}

func (om *OrderedMap[K, T]) Set(key K, value T) {
	if _, ok := om.values[key]; !ok {
		om.keys = append(om.keys, key)
	}
	om.values[key] = value
}

func (om *OrderedMap[K, T]) Get(key K) (T, bool) {
	value, ok := om.values[key]
	return value, ok
}

func (om *OrderedMap[K, T]) Keys() []K {
	keys := make([]K, len(om.keys))
	copy(keys, om.keys)
	return keys
}

func (om *OrderedMap[K, T]) Delete(key K) bool {
	if _, ok := om.values[key]; !ok {
		return false
	}

	delete(om.values, key)
	for i, k := range om.keys {
		if k == key {
			om.keys = append(om.keys[:i], om.keys[i+1:]...)
			break
		}
	}

	return true
}

func (om *OrderedMap[K, T]) Len() int {
	return len(om.keys)
}
