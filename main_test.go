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

func TestMain(t *testing.T) {
	expectedAddResult := 5
	expectedSubtractResult := 5

	// Run the main function
	main()

	// Check the result of Add function
	if result := Add(2, 3); result != expectedAddResult {
		t.Errorf("Add function returned incorrect result. Expected: %d, Got: %d", expectedAddResult, result)
	}

	// Check the result of Subtract function
	if result := Subtract(10, 5); result != expectedSubtractResult {
		t.Errorf("Subtract function returned incorrect result. Expected: %d, Got: %d", expectedSubtractResult, result)
	}
}
