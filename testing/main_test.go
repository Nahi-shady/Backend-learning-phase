package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Calculate(n int) int {
	return n * 2
}

func TestCalculate(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, -2},
		{0, 0},
		{-5, -10},
		{9, 18},
	}

	for _, test := range tests {
		assert.Equal(Calculate(test.input), test.expected)
	}

}
