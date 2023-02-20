package main

import (
	"fmt"
	"time"
)

// 利用缓存优化斐波那契 131958 < 567671875

const (
	LIMIT = 42
)

var fi [LIMIT]uint64

func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < 40; i++ {
		result = fibonacciMemoization(i)
		fmt.Printf("Fibonacci(%d) result is %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("execute time is %d\n", delta)
}

func fibonacciMemoization(n int) (res uint64) {
	if fi[n] != 0 {
		res = fi[n]
		return res
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacciMemoization(n-1) + fibonacciMemoization(n-2)
	}
	fi[n] = res
	return
}
