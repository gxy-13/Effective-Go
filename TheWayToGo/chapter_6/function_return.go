package main

import "fmt"

func main() {
	one := Adder()
	fmt.Printf("one adder result is %d\n", one(1))
	two := Adder2(1)
	fmt.Printf("two adder result is %d\n", two(2))
}

func Adder() func(b int) int {
	return func(b int) int {
		return 1 + b
	}
}

func Adder2(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
