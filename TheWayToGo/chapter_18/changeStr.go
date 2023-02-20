package main

import "fmt"

func main() {
	str := "hello"
	c := []byte(str)
	c[0] = 'a'
	s2 := string(c)
	fmt.Println(s2)
}
