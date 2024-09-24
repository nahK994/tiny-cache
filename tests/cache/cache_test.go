package cache

import (
	"reflect"
	"testing"

	"github.com/nahK994/TinyCache/pkg/config"
)

func TestCache(t *testing.T) {
	// Initialize a new cache
	c := config.App.Cache

	t.Run("TestSETAndGET", func(t *testing.T) {
		// Test string value
		c.SET("name", "Shomi")
		item := c.GET("name")
		if item.Val != "Shomi" {
			t.Errorf("Expected 'Shomi', got %v", item.Val)
		}

		// Test integer value
		c.SET("age", 25)
		item = c.GET("age")
		if item.Val != 25 {
			t.Errorf("Expected 25, got %v", item.Val)
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
		val := c.LRANGE("numbers", 0, 2)
		expected := []string{"three", "two", "one"}
		if !reflect.DeepEqual(val, expected) {
			t.Errorf("Expected %v, got %v", expected, val)
		}
	})

	t.Run("TestLPUSH_LPOP", func(t *testing.T) {
		c.LPUSH("items1", []string{"item1", "item2", "item3"})
		c.LPOP("items1")
		val := c.LRANGE("items1", 0, 1)
		expected := []string{"item2", "item1"}
		if !reflect.DeepEqual(val, expected) {
			t.Errorf("Expected %v, got %v", expected, val)
		}
	})

	t.Run("TestRPUSH_RPOP", func(t *testing.T) {
		c.RPUSH("items2", []string{"item1", "item2", "item3"})
		c.RPOP("items2")
		val := c.LRANGE("items2", 0, -1)
		expected := []string{"item1", "item2"}
		if !reflect.DeepEqual(val, expected) {
			t.Errorf("Expected %v, got %v", expected, val)
		}
	})

	t.Run("TestFLUSHALL", func(t *testing.T) {
		c.SET("key", "value")
		c.FLUSHALL()
		if c.EXISTS("key") {
			t.Errorf("'key' exists after FLUSHALL")
		}
	})
}
