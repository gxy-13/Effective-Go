package main

import "fmt"

type Address struct {
	addr string
}

type VCard struct {
	name  string
	birth string
	addr  *Address
}

func main() {
	v := &VCard{
		"Tom",
		"1999-1-1",
		&Address{
			"名称大道",
		},
	}
	fmt.Printf("%v\n", v)
}
