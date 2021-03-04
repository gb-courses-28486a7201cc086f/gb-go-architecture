package cache

import "testing"

var (
	testLRUData map[string]int = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}
	testCacheSize = 3
)

func TestLRUCache(t *testing.T) {
	cache, err := NewLRUCache(testCacheSize)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Run("fill cache", func(t *testing.T) {
		testKeys := []string{"one", "two", "three"}
		expectedKeys := []string{"three", "two", "one"}

		for _, key := range testKeys {
			cache.Set(key, testLRUData[key])
		}

		gotKeys := cache.Keys()

		if len(gotKeys) != len(expectedKeys) {
			t.Errorf("too may keys, got %d, expected %d", len(gotKeys), len(expectedKeys))
		}

		for i := range gotKeys {
			if gotKeys[i] != expectedKeys[i] {
				t.Errorf("invalid keys order, got %s, expected %s", gotKeys[i], expectedKeys[i])
			}
		}
	})

	t.Run("get key", func(t *testing.T) {
		testKey := "one"
		expectedValue := testLRUData[testKey]
		// get operation should hit key
		expectedKeys := []string{"one", "three", "two"}

		value, err := cache.Get(testKey)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if value != expectedValue {
			t.Errorf("invalid value got %d, expected %d", value, expectedValue)
		}

		gotKeys := cache.Keys()

		if len(gotKeys) != len(expectedKeys) {
			t.Errorf("too may keys, got %d, expected %d", len(gotKeys), len(expectedKeys))
		}

		for i := range gotKeys {
			if gotKeys[i] != expectedKeys[i] {
				t.Errorf("invalid keys order, got %s, expected %s", gotKeys[i], expectedKeys[i])
			}
		}
	})

	t.Run("update key", func(t *testing.T) {
		testKey := "two"
		testValue := 20
		// update operation should hit key
		expectedKeys := []string{"two", "one", "three"}

		cache.Set(testKey, testValue)
		updValue, err := cache.Get(testKey)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if updValue != testValue {
			t.Errorf("invalid value got %d, expected %d", updValue, testValue)
		}

		gotKeys := cache.Keys()

		if len(gotKeys) != len(expectedKeys) {
			t.Errorf("too may keys, got %d, expected %d", len(gotKeys), len(expectedKeys))
		}

		for i := range gotKeys {
			if gotKeys[i] != expectedKeys[i] {
				t.Errorf("invalid keys order, got %s, expected %s", gotKeys[i], expectedKeys[i])
			}
		}
	})

	t.Run("add extra key", func(t *testing.T) {
		testKey := "four"
		testValue := testLRUData[testKey]
		// add extra key should drop the older one
		expectedKeys := []string{"four", "two", "one"}
		olderKey := "three"

		cache.Set(testKey, testValue)
		value, err := cache.Get(testKey)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if value != testValue {
			t.Errorf("invalid value got %d, expected %d", value, testValue)
		}

		// older value should not exists
		_, err = cache.Get(olderKey)
		if err != ErrNotExists {
			t.Fatalf("unexpected error: got %v, expected %v", err, ErrNotExists)
		}

		gotKeys := cache.Keys()

		if len(gotKeys) != len(expectedKeys) {
			t.Errorf("too may keys, got %d, expected %d", len(gotKeys), len(expectedKeys))
		}

		for i := range gotKeys {
			if gotKeys[i] != expectedKeys[i] {
				t.Errorf("invalid keys order, got %s, expected %s", gotKeys[i], expectedKeys[i])
			}
		}
	})

	t.Run("delete key", func(t *testing.T) {
		testKey := "four"
		expectedKeys := []string{"two", "one"}

		err := cache.Delete(testKey)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// deleted value should not exists
		_, err = cache.Get(testKey)
		if err != ErrNotExists {
			t.Fatalf("unexpected error: got %v, expected %v", err, ErrNotExists)
		}

		gotKeys := cache.Keys()

		if len(gotKeys) != len(expectedKeys) {
			t.Errorf("too may keys, got %d, expected %d", len(gotKeys), len(expectedKeys))
		}

		for i := range gotKeys {
			if gotKeys[i] != expectedKeys[i] {
				t.Errorf("invalid keys order, got %s, expected %s", gotKeys[i], expectedKeys[i])
			}
		}
	})

	t.Run("delete not existing key", func(t *testing.T) {
		testKey := "notexistsing"

		err := cache.Delete(testKey)
		if err != ErrNotExists {
			t.Fatalf("unexpected error: got %v, expected %v", err, ErrNotExists)
		}
	})
}
