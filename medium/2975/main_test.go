package main

import (
	"os"
	"testing"
)

type testCase struct {
	name     string
	m        int
	n        int
	hFences  []int
	vFences  []int
	expected int
}

var testCases []testCase

// TestMain is the setup phase for all tests in this package
func TestMain(m *testing.M) {
	// Setup: Initialize test cases
	testCases = []testCase{
		{
			name:     "Test case 1: m=4, n=3, hFences=[2,3], vFences=[2]",
			m:        4,
			n:        3,
			hFences:  []int{2, 3},
			vFences:  []int{2},
			expected: 4,
		},
		{
			name:     "Test case 2: m=6, n=7, hFences=[2], vFences=[4]",
			m:        6,
			n:        7,
			hFences:  []int{2},
			vFences:  []int{4},
			expected: -1,
		},
	}

	// Run all tests
	exitCode := m.Run()

	// Teardown (if needed)
	// Add any cleanup code here

	os.Exit(exitCode)
}

func TestAkitKumar(t *testing.T) {
	tests := testCases

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AkitKumar(tt.m, tt.n, tt.hFences, tt.vFences)
			if result != tt.expected {
				t.Errorf("AkitKumar(%d, %d, %v, %v) = %d; expected %d",
					tt.m, tt.n, tt.hFences, tt.vFences, result, tt.expected)
			}
		})
	}
}

func TestPattoBear(t *testing.T) {
	tests := testCases

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PattoBear(tt.m, tt.n, tt.hFences, tt.vFences)
			if result != tt.expected {
				t.Errorf("PattoBear(%d, %d, %v, %v) = %d; expected %d",
					tt.m, tt.n, tt.hFences, tt.vFences, result, tt.expected)
			}
		})
	}
}
