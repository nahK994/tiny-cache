package server

import (
	"testing"
	"time"

	"github.com/nahK994/TinyCache/connection/server"
	"github.com/nahK994/TinyCache/pkg/config"
	"github.com/nahK994/TinyCache/pkg/errors"
)

var mockCache = config.App.Cache

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
	server.HandleCommand("*3\r\n$3\r\nSET\r\n$11\r\nexpired_key\r\n$5\r\nvalue\r\n")
	server.HandleCommand("*3\r\n$6\r\nEXPIRE\r\n$11\r\nexpired_key\r\n$1\r\n1\r\n")
	time.Sleep(1 * time.Second)
	err := server.AssertKeyExists(key)
	if err == nil {
		t.Errorf("Expected error for expired key, but got none")
	} else if err.Error() != (errors.Err{Type: errors.ExpiredKey}).Error() {
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
