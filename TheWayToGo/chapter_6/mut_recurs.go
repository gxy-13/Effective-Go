package main

import "fmt"

// 相互调用的递归函数，多个函数之间相互调用形成闭环
func main() {
	fmt.Printf("%d is even: is %t\n", 16, even(16))
}

func even(n int) bool {
	if n == 0 {
		return true
	}
	return odd(revSign(n) - 1)
}

func odd(n int) bool {
	if n == 0 {
		return false
	}
	return even(revSign(n) - 1)
}

func revSign(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
