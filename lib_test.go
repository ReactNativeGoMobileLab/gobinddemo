package gobinddemo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, Sum(1, 2), 3)
}

func TestFibonacci(t *testing.T) {
	assert.Equal(t, Fibonacci(0), []int{})
	assert.Equal(t, Fibonacci(1), []int{1})
	assert.Equal(t, Fibonacci(2), []int{1, 1})
	assert.Equal(t, Fibonacci(3), []int{1, 1, 2})
	assert.Equal(t, Fibonacci(4), []int{1, 1, 2, 3})
	assert.Equal(t, Fibonacci(10), []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55})
}
