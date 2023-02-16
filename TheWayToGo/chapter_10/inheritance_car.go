package main

import "fmt"

type Engine interface {
	stop()
	start()
}

type Car struct {
	Engine
	wheelCount int
}

func (c *Car) GoToWorkIn() {
	c.start()
	c.stop()
}

func (c *Car) numberOfWheel() int {
	return c.wheelCount
}

func (c *Car) start() {
	fmt.Println("start....")
}

func (c *Car) stop() {
	fmt.Println("stop....")
}

type Mercedes struct {
	Car
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("hi Merkel")
}

func main() {
	car := Car{wheelCount: 4}
	car.GoToWorkIn()
	m := Mercedes{
		Car{nil, 4},
	}
	fmt.Println(m.numberOfWheel())
	m.GoToWorkIn()
	m.sayHiToMerkel()
}
