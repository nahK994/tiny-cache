package cache

import (
	"reflect"
	"testing"

	"github.com/nahK994/TinyCache/pkg/cache"
)

func TestCache(t *testing.T) {
	// Initialize a new cache
	c := cache.InitCache()

	t.Run("TestSETAndGET", func(t *testing.T) {
		// Test string value
		c.SET("name", "Shomi")
		val := c.GET("name")
		if val != "Shomi" {
			t.Errorf("Expected 'Shomi', got %v", val)
		}

		// Test integer value
		c.SET("age", 25)
		val = c.GET("age")
		if val != 25 {
			t.Errorf("Expected 25, got %v", val)
		}
	})

	t.Run("TestINCRAndDECR", func(t *testing.T) {
		// Test INCR
		c.SET("counter", 10)
		val := c.INCR("counter")
		if val != 11 {
			t.Errorf("Expected 11, got %v", val)
		}

		// Test DECR
		val = c.DECR("counter")
		if val != 10 {
			t.Errorf("Expected 10, got %v", val)
		}
	})

	t.Run("TestEXIST", func(t *testing.T) {
		c.SET("language", "Go")
		if !c.EXISTS("language") {
			t.Errorf("Expected key 'language' to exist")
		}
		if c.EXISTS("non-existent") {
			t.Errorf("Expected 'non-existent' key to not exist")
		}
	})

	t.Run("TestDEL", func(t *testing.T) {
		c.SET("delete-me", "test")
		c.DEL("delete-me")
		if c.EXISTS("delete-me") {
			t.Errorf("Expected key 'delete-me' to be deleted")
		}
	})

	t.Run("TestLPUSHAndLRANGE", func(t *testing.T) {
		// Test LPUSH
		c.LPUSH("numbers", []string{"one", "two", "three"})
		val := c.LRANGE("numbers", 1, -1)
		expected := []string{"three", "two", "one"}
		if !reflect.DeepEqual(val, expected) {
			t.Errorf("Expected %v, got %v", expected, val)
		}
	})

	t.Run("TestLPOP", func(t *testing.T) {
		c.LPUSH("items", []string{"item1", "item2", "item3"})
		c.LPOP("items")
		val := c.LRANGE("items", 1, -1)
		expected := []string{"item2", "item1"}
		if !reflect.DeepEqual(val, expected) {
			t.Errorf("Expected %v, got %v", expected, val)
		}
	})
}
