package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s *stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c *car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(v valuable) {
	fmt.Println(v.getValue())
}

func main() {
	var o valuable = &stockPosition{"GOOG", 401.2, 5}
	showValue(o)
	o = &car{"Benz", "c260l", 500}
	showValue(o)
}
