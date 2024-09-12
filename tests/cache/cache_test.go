package cache_test

// func TestCache(t *testing.T) {
// 	cacheInstance := cache.InitCache()

// 	// Test case 1: Empty cache should not have any keys
// 	t.Run("EmptyCache", func(t *testing.T) {
// 		if cacheInstance.IsKeyExist("nonexistent") {
// 			t.Errorf("Expected false, got true for nonexistent key")
// 		}
// 		if cacheInstance.ReadCache("nonexistent") != nil {
// 			t.Errorf("Expected nil, got non-nil for nonexistent key")
// 		}
// 	})

// 	// Test case 2: Insert a key and read it back
// 	t.Run("InsertAndRead", func(t *testing.T) {
// 		cacheInstance.SetCache("name", "Shomi") // Using SetCache method

// 		if !cacheInstance.IsKeyExist("name") {
// 			t.Errorf("Expected true, got false for key 'name'")
// 		}
// 		if val := cacheInstance.ReadCache("name"); val != "Shomi" {
// 			t.Errorf("Expected 'Shomi', got %v", val)
// 		}
// 	})

// 	// Test case 3: Overwrite an existing key
// 	t.Run("OverwriteKey", func(t *testing.T) {
// 		cacheInstance.SetCache("name", "Khan") // Using SetCache to overwrite

// 		if val := cacheInstance.ReadCache("name"); val != "Khan" {
// 			t.Errorf("Expected 'Khan', got %v", val)
// 		}
// 	})

// 	// Test case 4: Test key existence for a different key
// 	t.Run("NonexistentKey", func(t *testing.T) {
// 		if cacheInstance.IsKeyExist("age") {
// 			t.Errorf("Expected false, got true for nonexistent key 'age'")
// 		}
// 	})
// }
