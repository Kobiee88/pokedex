package main

import "testing"

func TestAddGet(t *testing.T) {
	initCache()
	testKey := "testKey"
	testValue := []byte("testValue")

	cache.Add(testKey, testValue)
	retrievedValue, exists := cache.Get(testKey)

	if !exists {
		t.Errorf("Expected key %s to exist in cache", testKey)
	}
	if string(retrievedValue) != string(testValue) {
		t.Errorf("Expected value %s, got %s", testValue, retrievedValue)
	}
}
