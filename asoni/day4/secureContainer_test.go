package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/number"
)

func Test_verifyDigits(t *testing.T) {
	testCases := []struct {
		num            int
		expectedAnswer bool
	}{
		{
			num:            123456,
			expectedAnswer: true,
		},
		{
			num:            12345,
			expectedAnswer: false,
		},
		{
			num:            1234,
			expectedAnswer: false,
		},
		{
			num:            123,
			expectedAnswer: false,
		},
		{
			num:            12,
			expectedAnswer: false,
		},
		{
			num:            1,
			expectedAnswer: false,
		},
		{
			num:            1234567,
			expectedAnswer: false,
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			ans := verifyDigits(tc.num)
			assert.Equal(t, tc.expectedAnswer, ans)
		})
	}
}

func Benchmark_verifyDigits(b *testing.B) {
	for n := 0; n < b.N; n++ {
		verifyDigits(278384)
	}
}

func Test_adjacentDigits(t *testing.T) {
	testCases := []struct{
		num	int
		expectedAnswer bool
	}{
		{
			num:	123455,
			expectedAnswer: true,
		}
	}
}