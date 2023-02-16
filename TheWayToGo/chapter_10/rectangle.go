package main

type Rectangle struct {
	x, y int
}

func (r *Rectangle) Area() int {
	s := r.x * r.y
	return s
}
