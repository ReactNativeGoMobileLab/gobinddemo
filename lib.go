package gobinddemo

func Sum(a, b int) int {
	return a + b
}

func Fibonacci(n int) []int {
	if n == 0 {
		return []int{}
	}

	v0 := 0
	v1 := 1
	fi := make([]int, n)
	fi[0] = v1
	for i := 1; i < n; i++ {
		fi[i] = v0 + v1
		v0, v1 = v1, fi[i]
	}
	return fi
}
