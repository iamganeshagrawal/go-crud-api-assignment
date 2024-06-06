package datatypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderedMap(t *testing.T) {
	om := NewOrderedMap[string, int]()

	// Test initial length
	assert.Equal(t, 0, om.Len(), "expected length 0, got %d", om.Len())

	// Test Set and Get
	om.Set("a", 1)
	value, ok := om.Get("a")
	assert.True(t, ok, "expected key 'a' to be found")
	assert.Equal(t, 1, value, "expected value 1, got %d", value)

	// Test setting another value
	om.Set("b", 2)
	value, ok = om.Get("b")
	assert.True(t, ok, "expected key 'b' to be found")
	assert.Equal(t, 2, value, "expected value 2, got %d", value)

	// Test length after adding values
	assert.Equal(t, 2, om.Len(), "expected length 2, got %d", om.Len())

	// Test key ordering
	expectedKeys := []string{"a", "b"}
	keys := om.Keys()
	assert.Equal(t, expectedKeys, keys, "expected keys %v, got %v", expectedKeys, keys)

	// Test updating a value
	om.Set("a", 3)
	value, ok = om.Get("a")
	assert.True(t, ok, "expected key 'a' to be found")
	assert.Equal(t, 3, value, "expected value 3, got %d", value)

	// Test deleting a key
	deleted := om.Delete("a")
	assert.True(t, deleted, "expected to delete key 'a', but it was not found")
	assert.Equal(t, 1, om.Len(), "expected length 1 after delete, got %d", om.Len())

	value, ok = om.Get("a")
	assert.False(t, ok, "expected key 'a' to be deleted, but it was found with value %d", value)

	// Test deleting a non-existing key
	deleted = om.Delete("a")
	assert.False(
		t,
		deleted,
		"expected to not delete non-existing key 'a', but it was reported as deleted",
	)

	// Test key ordering after deletion
	expectedKeys = []string{"b"}
	keys = om.Keys()
	assert.Equal(t, expectedKeys, keys, "expected keys %v, got %v", expectedKeys, keys)

	// Test setting a new key after deletion
	om.Set("c", 4)
	assert.Equal(t, 2, om.Len(), "expected length 2 after adding 'c', got %d", om.Len())

	expectedKeys = []string{"b", "c"}
	keys = om.Keys()
	assert.Equal(t, expectedKeys, keys, "expected keys %v, got %v", expectedKeys, keys)
}
