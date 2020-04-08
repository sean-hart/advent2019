package main

import (
	"testing"
)

func TestRunComputer(t *testing.T) {
	for _, tc := range testCasesRunComputer {
		actual := RunComputer(tc.input)
		if !SliceEqual(actual, tc.expected) {
			t.Fatalf("FAIL: %v Expected: %v Actual: %v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkRunComputer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCasesRunComputer {
			RunComputer(tc.input)
		}
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func SliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
