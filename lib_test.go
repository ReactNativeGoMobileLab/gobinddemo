package gobinddemo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, Sum(1, 2), 3)
}

func TestFibonacci(t *testing.T) {
	assert.Equal(t, Fibonacci(0), 0)
	assert.Equal(t, Fibonacci(1), 1)
	assert.Equal(t, Fibonacci(2), 1)
	assert.Equal(t, Fibonacci(3), 2)
	assert.Equal(t, Fibonacci(4), 3)
	assert.Equal(t, Fibonacci(10), 55)
}
