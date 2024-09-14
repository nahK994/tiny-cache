package cache

import (
	"testing"

	"github.com/nahK994/TinyCache/pkg/cache"
)

func TestCache(t *testing.T) {
	// Initialize a new cache
	c := cache.InitCache()

	// Test WriteCache and ReadCache for string value
	t.Run("TestWriteAndReadCache", func(t *testing.T) {
		c.WriteCache("name", "Shomi")

		val := c.ReadCache("name")
		if val != "Shomi" {
			t.Errorf("Expected 'Shomi', got %v", val)
		}
	})

	// Test WriteCache and ReadCache for int value
	t.Run("TestWriteAndReadIntCache", func(t *testing.T) {
		c.WriteCache("age", 25)

		val := c.ReadCache("age")
		if val != 25 {
			t.Errorf("Expected 25, got %v", val)
		}
	})

	// Test IsKeyExist for existing and non-existing keys
	t.Run("TestIsKeyExist", func(t *testing.T) {
		if !c.IsKeyExist("age") {
			t.Errorf("Expected key 'age' to exist")
		}

		if c.IsKeyExist("nonexistent") {
			t.Errorf("Expected key 'nonexistent' to not exist")
		}
	})

	// Test INCRCache for an existing int key
	t.Run("TestINCRCache", func(t *testing.T) {
		result := c.INCRCache("age")
		expectedResult := ":26\r\n"
		if result != expectedResult {
			t.Errorf("Expected '%s', got %s", expectedResult, result)
		}

		// Check if the incremented value is correct
		if c.ReadCache("age") != 26 {
			t.Errorf("Expected 'age' to be 26, got %v", c.ReadCache("age"))
		}
	})

	// Test DECRCache for an existing int key
	t.Run("TestDECRCache", func(t *testing.T) {
		result := c.DECRCache("age")
		expectedResult := ":25\r\n"
		if result != expectedResult {
			t.Errorf("Expected '%s', got %s", expectedResult, result)
		}

		// Check if the decremented value is correct
		if c.ReadCache("age") != 25 {
			t.Errorf("Expected 'age' to be 25, got %v", c.ReadCache("age"))
		}
	})

	// Test DeleteCache for removing a key
	t.Run("TestDeleteCache", func(t *testing.T) {
		c.DeleteCache("name")
		if c.IsKeyExist("name") {
			t.Errorf("Expected key 'name' to be deleted")
		}
	})
}
