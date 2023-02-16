package main

import "fmt"

// Go 不支持面向对象编程语言中那样的构造子方法，但是Go很容易实现构造子工厂方法
// 工厂名字以 new New 开头

type File struct {
	fd   int    // 文件描述符
	name string // 文件名
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

func main() {
	// 等价于 &File{}
	f := NewFile(10, "./test.txt")
	fmt.Println(f)
	// 将结构体名称变成小写，就可以强制使用工厂方法
}
