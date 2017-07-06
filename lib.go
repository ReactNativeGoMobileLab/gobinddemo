package gobinddemo

func Sum(a, b int) int {
	return a + b
}

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	prev := 0
	curr := 1
	for i := 1; i < n; i++ {
		curr, prev = curr+prev, curr
	}
	return curr
}
