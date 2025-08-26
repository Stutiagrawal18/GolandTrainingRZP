package main

import "testing"

func TestSum(t *testing.T) {

	result := sum(2, 3)
	expected := 5

	// Compare the result with the expected value
	if result != expected {
		// If they don't match, use t.Errorf to report the failure.
		// t.Errorf logs a formatted error and marks the test as failed,
		// but allows the test function to continue running.
		t.Errorf("Sum(2, 3) returned %d, but expected %d", result, expected)
	}
}

// math_test.go
func TestSumTableDriven(t *testing.T) {
	// Define a slice of structs to hold your test cases
	testCases := []struct {
		name     string // A descriptive name for the test case
		a, b     int    // Inputs
		expected int    // The expected output
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -1, -5, -6},
		{"mixed numbers", 10, -3, 7},
		{"zero", 0, 0, 0},
	}

	// Iterate over the test cases
	for _, tc := range testCases {
		// Use t.Run to create a subtest for each case. This gives a nice output.
		t.Run(tc.name, func(t *testing.T) {
			result := sum(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Sum(%d, %d) returned %d, but expected %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}
