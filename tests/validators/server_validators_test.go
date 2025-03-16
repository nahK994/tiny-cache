package validators

import (
	"testing"
	"time"

	"github.com/nahK994/TinyCache/pkg/cache"
	"github.com/nahK994/TinyCache/pkg/config"
	"github.com/nahK994/TinyCache/pkg/errors"
	"github.com/nahK994/TinyCache/pkg/validators"
)

var c *cache.Cache = config.App.Cache

func TestValidator(t *testing.T) {
	c.SET("key", "value")
	c.EXPIRE("key", 5)
	time.Sleep(2 * time.Second)
	err := validators.ValidateExpiry("key")
	if err != nil {
		t.Errorf("'key' isn't supposed be expired so soon")
	}

	time.Sleep(6 * time.Second)
	err = validators.ValidateExpiry("key")
	if err == nil {
		t.Errorf("Expected %v", errors.Err{Type: errors.ExpiredKey})
	}
}
