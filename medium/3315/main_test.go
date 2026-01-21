package main

import (
	"os"
	"slices"
	"testing"
)

type testCase struct {
	name     string
	nums     []int
	expected []int
}

var testCases []testCase

func TestMain(m *testing.M) {
	// Setup: Initialize test cases
	testCases = []testCase{
		{
			name:     "Test case 1: nums=[2,3,5,7]",
			nums:     []int{2, 3, 5, 7},
			expected: []int{-1, 1, 4, 3},
		},
		{
			name:     "Test case 2: nums=[11,13,31]",
			nums:     []int{11, 13, 31},
			expected: []int{9, 12, 15},
		},
	}

	// Run all tests
	exitCode := m.Run()

	// Teardown (if needed)
	// Add any cleanup code here

	os.Exit(exitCode)
}

func TestMinBitwiseArray(t *testing.T) {
	tests := testCases

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minBitwiseArray(tt.nums)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("minBitwiseArray(%v) = %v; expected %v",
					tt.nums, result, tt.expected)
			}
		})
	}
}
