package main

import (
	"os"
	"testing"
)

type testCase struct {
	name      string
	mat       [][]int
	threshold int
	expected  int
}

var testCases []testCase

func TestMain(m *testing.M) {
	// Setup: Initialize test cases
	testCases = []testCase{
		{
			name:      "Test case 1",
			mat:       [][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}},
			threshold: 4,
			expected:  2,
		},
		{
			name:      "Test case 2",
			mat:       [][]int{{1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}},
			threshold: 6,
			expected:  3,
		},
	}

	// Run all tests
	exitCode := m.Run()

	// Teardown (if needed)
	// Add any cleanup code here

	os.Exit(exitCode)
}

func TestMaxSideLength(t *testing.T) {
	tests := testCases

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSideLength(tt.mat, tt.threshold)
			if result != tt.expected {
				t.Errorf("maxSideLength(%v, %d) = %d; expected %d",
					tt.mat, tt.threshold, result, tt.expected)
			}
		})
	}
}
