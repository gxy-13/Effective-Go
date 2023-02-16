package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	width, length float32
}

func (re *Rectangle) Area() float32 {
	return re.width * re.length
}

func main() {
	sq1 := new(Square)
	sq1.side = 5
	re1 := new(Rectangle)
	re1.width = 2
	re1.length = 3

	var areaIntf Shaper
	areaIntf = sq1
	fmt.Println(areaIntf.Area())

	areaIntf = re1
	fmt.Println(areaIntf.Area())

	arr := []Shaper{
		&Square{4},
		&Rectangle{5, 7},
	}
	for _, s := range arr {
		fmt.Println(s.Area())
	}
}
