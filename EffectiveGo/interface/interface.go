package main

import "fmt"

type Stringer interface {
	String() string
}

type student struct {
	name string
}

func (s *student) String() string {
	return s.name
}

func Paint(val interface{}) interface{} {
	// 类型选择
	switch str := val.(type) {
	case string:
		return str
	case Stringer:
		return str.String()
	}
	return nil
}

func main() {
	stu := new(student)
	stu.name = "hello"
	fmt.Println(Paint(stu))
	fmt.Println("Hello,Go!")
}
