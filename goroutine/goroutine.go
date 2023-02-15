package main

import "fmt"

func main() {
	// 创建一个无缓冲的channel
	c := make(chan int)
	// 用goroutine开启一个排序，完成时会发出信号
	go func() {
		arr := []int{1, 2, 7, 6, 5}
		quickSort(arr, 0, 4)
		c <- 1
	}()
	// 在从c中获取到值之前一直阻塞
	i := <-c
	fmt.Println(i, "排序完成")
}

func quickSort(a []int, l, r int) {
	if l >= r {
		return
	}
	x := a[l+(r-l)/2]
	i, j := l-1, r+1
	for i < j {
		for {
			i++
			if a[i] >= x {
				break
			}
		}
		for {
			j--
			if a[j] <= x {
				break
			}
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	quickSort(a, l, j)
	quickSort(a, j+1, r)
}
