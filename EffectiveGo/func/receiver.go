package main

import (
	"fmt"
)

func main() {
	var s = &ByteSlice{'a', 'b', 'c'}
	s.AppendMethod([]byte{'e', 'f', 'g'})
	fmt.Println(*s)
}

type ByteSlice []byte

func (b *ByteSlice) AppendMethod(data []byte) {
	slice := *b
	l := len(slice)
	if l+len(data) > cap(slice) {
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	*b = slice
}
