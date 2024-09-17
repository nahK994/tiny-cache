package cache

import (
	"testing"

	"github.com/nahK994/TinyCache/pkg/cache"
)

func TestCache(t *testing.T) {
	// Initialize a new cache
	c := cache.InitCache()

	// Test SET and GET for string value
	t.Run("TestWriteAndReadCache", func(t *testing.T) {
		c.SET("name", "Shomi")

		val := c.GET("name")
		if val != "Shomi" {
			t.Errorf("Expected 'Shomi', got %v", val)
		}
	})

	// Test SET and GET for int value
	t.Run("TestWriteAndReadIntCache", func(t *testing.T) {
		c.SET("age", 25)

		val := c.GET("age")
		if val != 25 {
			t.Errorf("Expected 25, got %v", val)
		}
	})

	// Test EXIST for existing and non-existing keys
	t.Run("TestIsKeyExist", func(t *testing.T) {
		if !c.EXIST("age") {
			t.Errorf("Expected key 'age' to exist")
		}

		if c.EXIST("nonexistent") {
			t.Errorf("Expected key 'nonexistent' to not exist")
		}
	})

	// Test INCR for an existing int key
	t.Run("TestINCRCache", func(t *testing.T) {
		result := c.INCR("age")
		expectedResult := 26
		if result != expectedResult {
			t.Errorf("Expected '%d', got %d", expectedResult, result)
		}

		// Check if the incremented value is correct
		if c.GET("age") != 26 {
			t.Errorf("Expected 'age' to be 26, got %v", c.GET("age"))
		}
	})

	// Test DECR for an existing int key
	t.Run("TestDECRCache", func(t *testing.T) {
		result := c.DECR("age")
		expectedResult := 25
		if result != expectedResult {
			t.Errorf("Expected '%d', got %d", expectedResult, result)
		}

		// Check if the decremented value is correct
		if c.GET("age") != 25 {
			t.Errorf("Expected 'age' to be 25, got %v", c.GET("age"))
		}
	})

	// Test DEL for removing a key
	t.Run("TestDeleteCache", func(t *testing.T) {
		c.DEL("name")
		if c.EXIST("name") {
			t.Errorf("Expected key 'name' to be deleted")
		}
	})
}
