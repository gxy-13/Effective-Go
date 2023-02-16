package main

type employee struct {
	salary float64
}

func (e *employee) giveRaise(x float64) {
	e.salary = (1 + x) * e.salary
}

func main() {

}
