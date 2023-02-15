package main

import (
	"fmt"
	"time"
)

const (
	MaxOutstanding = 10
)

// 设置一个缓冲区容量为10的channel 来限制吞吐量
var sem = make(chan int, MaxOutstanding)

func process(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i < 20; i++ {
		sem <- 1 // 有一个请求就放一个1进缓冲区
		go func(i int) {
			process(i)
			<-sem // 执行结束从缓冲区腾出空间
			// 确保每次goroutine都是唯一的i
		}(i)
	}
	time.Sleep(time.Second * 5)
}
