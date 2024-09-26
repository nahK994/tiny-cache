package handlers

func IsKeyExists(key string) bool {
	_ = validateExpiry(key)
	return c.EXISTS(key)
}
