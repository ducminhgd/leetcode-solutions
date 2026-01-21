package main

import (
	"os"
	"testing"
)

type testCase struct {
	name     string
	skill    []int
	mana     []int
	expected int64
}

var testCases []testCase

func TestMain(m *testing.M) {
	// Setup: Initialize test cases
	testCases = []testCase{
		{
			name:     "Test case 2: skill=[5,4,3,2,1], mana=[1,2,3,4,5]",
			skill:    []int{5, 4, 3, 2, 1},
			mana:     []int{1, 2, 3, 4, 5},
			expected: 35,
		},
		{
			name:     "Test case 3: skill=[1,1,1], mana=[1,1,1]",
			skill:    []int{1, 1, 1},
			mana:     []int{1, 1, 1},
			expected: 3,
		},
		{
			name:     "Test case 4: skill=[1,2,3,4], mana=[1,2]",
			skill:    []int{1, 2, 3, 4},
			mana:     []int{1, 2},
			expected: 21,
		},
	}
	// Run all tests
	exitCode := m.Run()

	// Teardown (if needed)
	// Add any cleanup code here

	os.Exit(exitCode)
}

func TestMinTime(t *testing.T) {
	tests := testCases

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minTime(tt.skill, tt.mana)
			if result != tt.expected {
				t.Errorf("minTime(%v, %v) = %d; expected %d",
					tt.skill, tt.mana, result, tt.expected)
			}
		})
	}
}
