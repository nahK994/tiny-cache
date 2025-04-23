package cache

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/nahK994/TinyCache/pkg/cache"
)

func TestCache(t *testing.T) {
	// Initialize a new cache
	c := cache.NewCache(10, 100)

	t.Run("TestSETAndGET", func(t *testing.T) {
		// Test string value
		c.SET("name", "Shomi")
		item, _ := c.GET("name")
		strItem := string(item.Value)
		if strItem != "Shomi" {
			t.Errorf("Expected 'Shomi', got %v", strItem)
		}

		c.SET("age", "25")
		item, _ = c.GET("age")
		intItem, _ := strconv.Atoi(string(item.Value))
		if intItem != 25 {
			t.Errorf("Expected 25, got %v", intItem)
		}

		_, isExists := c.GET("new_key")
		if isExists {
			t.Errorf("new_key not exists")
		}
	})

	t.Run("TestINCRAndDECR", func(t *testing.T) {
		c.SET("counter", 10)
		val := c.INCR("counter")
		if val != 11 {
			t.Errorf("Expected 11, got %v", val)
		}

		val = c.DECR("counter")
		if val != 10 {
			t.Errorf("Expected 10, got %v", val)
		}
	})

	t.Run("TestEXISTS", func(t *testing.T) {
		c.SET("language", "Go")
		if !c.EXISTS("language") {
			t.Errorf("Expected key 'language' to exist")
		}
		if c.EXISTS("non-existent") {
			t.Errorf("Expected 'non-existent' key not to exist")
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
		c.LPUSH("numbers", []string{"one", "two", "three"})
		val := c.LRANGE("numbers", -5, 6)
		expected := []string{"three", "two", "one"}
		if !reflect.DeepEqual(val, expected) {
			t.Errorf("Expected %v, got %v", expected, val)
		}

		c.LPUSH("numbers", []string{"four", "five", "six"})
		val = c.LRANGE("numbers", 0, 2)
		expected = []string{"six", "five", "four"}
		if !reflect.DeepEqual(val, expected) {
			t.Errorf("Expected %v, got %v", expected, val)
		}
	})

	t.Run("TestRPUSHAndLRANGE", func(t *testing.T) {
		c.RPUSH("numbers1", []string{"one", "two", "three"})
		val := c.LRANGE("numbers1", -5, 6)
		expected := []string{"one", "two", "three"}
		if !reflect.DeepEqual(val, expected) {
			t.Errorf("Expected %v, got %v", expected, val)
		}

		c.RPUSH("numbers1", []string{"four", "five", "six"})
		val = c.LRANGE("numbers1", 3, 5)
		expected = []string{"four", "five", "six"}
		if !reflect.DeepEqual(val, expected) {
			t.Errorf("Expected %v, got %v", expected, val)
		}
	})

	t.Run("TestLPUSH_LPOP", func(t *testing.T) {
		c.LPUSH("items1", []string{"item1", "item2", "item3"})
		c.LPOP("items1")
		val := c.LRANGE("items1", -3, -1)
		expected := []string{"item2", "item1"}
		if !reflect.DeepEqual(val, expected) {
			t.Errorf("Expected %v, got %v", expected, val)
		}
	})

	t.Run("TestRPUSH_RPOP", func(t *testing.T) {
		c.RPUSH("items2", []string{"item1", "item2", "item3"})
		c.RPOP("items2")
		val := c.LRANGE("items2", 0, 1)
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

	t.Run("TestEXPIRE", func(t *testing.T) {
		c.SET("key", "value")
		c.EXPIRE("key", 5)
		data, _ := c.GET("key")

		time.Sleep(2 * time.Second)
		if time.Now().After(*data.ExpiryTime) {
			t.Errorf("'key' isn't supposed to be expired so soon")
		}

		time.Sleep(6 * time.Second)
		if !time.Now().After(*data.ExpiryTime) {
			t.Errorf("'key' is supposed to be expired by now")
		}
	})

	t.Run("TestIncrementFrequency", func(t *testing.T) {
		c.SET("count", "0")
		c.IncrementFrequency("count")
		c.IncrementFrequency("count")

		item, _ := c.GET("count")
		if (*item).Frequency != 2 {
			t.Errorf("Expected frequency 2, got %d", item.Frequency)
		}
	})
}
