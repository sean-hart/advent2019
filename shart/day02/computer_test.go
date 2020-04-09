package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunComputer(t *testing.T) {
	for _, tc := range testCasesRunComputer {
		newInput := make([]int, len(tc.input))
		copy(newInput, tc.input)
		actual, _ := RunComputer(newInput, 0)
		assert.Equal(t, actual, tc.expected)
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkRunComputer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCasesRunComputer {
			newInput := make([]int, len(tc.input))
			copy(newInput, tc.input)
			RunComputer(newInput, 0)
		}
	}
}
