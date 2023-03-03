package fibo

var fiboList = []int{0, 1, 1}

func fib(n int) int {
	if len(fiboList) > n {
		return fiboList[n]
	}

	fiboList = append(fiboList, fib(n-1)+fib(n-2))
	return fiboList[n]
}
