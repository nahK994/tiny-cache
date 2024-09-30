package server

import (
	"time"

	"github.com/nahK994/TinyCache/pkg/errors"
)

func validateExpiry(key string) error {
	item := c.GET(key)
	if item.ExpiryTime != nil && time.Now().After(*item.ExpiryTime) {
		c.DEL(key)
		return errors.Err{Type: errors.ExpiredKey}
	}
	return nil
}

func AssertKeyExists(key string) error {
	if err := validateExpiry(key); err != nil {
		return err
	}

	if !c.EXISTS(key) {
		return errors.Err{Type: errors.UndefinedKey}
	}
	return nil
}
