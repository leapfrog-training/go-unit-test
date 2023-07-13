package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 4 {
		t.Errorf("Expected 4, but got %d", result)
	}
}
