package leetFibonacciNumber

func fib(n int) int {
	dpc := make([]int, 2)
	dpc[0] = 0
	dpc[1] = 1
	for i := 2; i <= n; i++ {
		dpc[i%2] += dpc[(i-1)%2]
	}
	return dpc[n%2]
}
