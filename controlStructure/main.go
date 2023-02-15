package main

import "fmt"

func main() {
	TestLabel([]int{8, 14, 13, 12, 11})
}

func TestLabel(src []int) {
	sum := 0
Loop:
	for n := 0; n < len(src); n++ {
		switch {
		case src[n] > 10:
			{
				if src[n] < 15 {
					sum++
					break
				}
			}
		case src[n] < 10:
			{
				if src[n] == 8 {
					break Loop
				}
			}
		}
	}
	sum = 5
	fmt.Println(sum)
}
