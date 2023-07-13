package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(10, 5)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}
