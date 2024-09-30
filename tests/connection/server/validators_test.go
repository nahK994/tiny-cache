package server

import (
	"testing"
	"time"

	"github.com/nahK994/TinyCache/connection/server"
	"github.com/nahK994/TinyCache/pkg/cache"
	"github.com/nahK994/TinyCache/pkg/errors"
)

var mockCache = cache.Init(60)

func TestAssertKeyExists_KeyExistsAndNotExpired(t *testing.T) {
	key := "existing_key"
	mockCache.SET(key, "value")
	err := server.AssertKeyExists(key)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestAssertKeyExists_KeyExpired(t *testing.T) {
	key := "expired_key"
	data := "value"
	expiryTime := time.Now().Add(-1 * time.Hour) // expired 1 hour ago
	mockCache.SET(key, cache.CacheItem{
		Value: cache.CacheData{
			DataType: cache.String,
			StrData:  &data,
		},
		ExpiryTime: &expiryTime,
	})

	// Test case: key exists but is expired
	err := server.AssertKeyExists(key)
	expectedErr := errors.Err{Type: errors.TypeError}
	if err == nil {
		t.Errorf("Expected error for expired key, but got none")
	} else if err.Error() != expectedErr.Error() {
		t.Errorf("Expected ExpiredKey error, but got: %v", err)
	}
}

func TestAssertKeyExists_KeyDoesNotExist(t *testing.T) {
	key := "non_existing_key"
	err := server.AssertKeyExists(key)
	expectedErr := errors.Err{Type: errors.UndefinedKey}
	if err == nil {
		t.Errorf("Expected error for non-existing key, but got none")
	} else if err.Error() != expectedErr.Error() {
		t.Errorf("Expected UndefinedKey error, but got: %v", err)
	}
}

func TestAssertKeyExists_KeyExistsButNoExpiry(t *testing.T) {
	key := "key_no_expiry"
	mockCache.SET(key, "value") // assuming no expiry by default

	// Test case: key exists, no expiry
	err := server.AssertKeyExists(key)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
