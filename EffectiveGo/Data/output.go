package main

import "fmt"

func main() {
	var s MyString = "Hello,Go!"
	fmt.Println(s)
}

type MyString string

func (m MyString) String() string {
	//  调用Sprintf 来构造 String 会无限递归 String 方法
	//	return fmt.Sprintf("MyString=%s", m)
	// 可以转化为基本的字符串类型， 它没有String方法
	return fmt.Sprintf("MyString=%s", string(m))
}
