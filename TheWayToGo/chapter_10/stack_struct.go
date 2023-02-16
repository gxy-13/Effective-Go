package main

import "fmt"

const LIMIT = 10

type stack struct {
	arr []int
	idx int
}

func NewStack() *stack {
	arr := make([]int, LIMIT)
	return &stack{arr, 0}
}

func (s *stack) push(i int) bool {
	if s.idx == cap(s.arr) {
		return false
	}
	s.arr[s.idx] = i
	s.idx++
	return true
}

func (s *stack) pop() bool {
	if s.idx < 0 {
		return false
	}
	s.idx--
	return true
}

func (s *stack) String() {
	for i := 0; i < s.idx; i++ {
		fmt.Printf("[%d:%d]\n", i, s.arr[i])
	}
}

func main() {
	s := NewStack()
	for i := 0; i < 10; i++ {
		s.push(i)
	}
	s.pop()
	s.String()
}
