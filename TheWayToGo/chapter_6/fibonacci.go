package main

import (
	"fmt"
	"time"
)

func main() {
	result := 0
	start := time.Now()
	for i := 0; i < 40; i++ {
		result = fibonacci(i)
		fmt.Printf("Fibonaac(%d) is:%d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("execute time is %d\n", delta)

}

func fibonacci(i int) (res int) {
	if i <= 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}
