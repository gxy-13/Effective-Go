package main

import "fmt"

type Base struct {
	id int
}

func (b *Base) SetId(id int) {
	b.id = id
}

func (b *Base) Id() int {
	return b.id
}

type person struct {
	Base
	FistName, LastName string
}

type Employee struct {
	person
	salary float64
}

func main() {
	e := new(Employee)
	e.SetId(10)
	fmt.Println(e.Id())
}
