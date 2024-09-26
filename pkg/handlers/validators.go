package handlers

import (
	"time"

	"github.com/nahK994/TinyCache/pkg/errors"
)

func CheckEmptyList(key string) error {
	if !IsKeyExists(key) {
		return errors.Err{Type: errors.EmptyList}
	}
	return nil
}

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

func AssertListType(key string) error {
	if _, ok := c.GET(key).Val.([]string); !ok {
		return errors.Err{Type: errors.TypeError}
	}
	return nil
}

func AssertIntType(key string) error {
	if _, ok := c.GET(key).Val.(int); !ok {
		return errors.Err{Type: errors.TypeError}
	}
	return nil
}
