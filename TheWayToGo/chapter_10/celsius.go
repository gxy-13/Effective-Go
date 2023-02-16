package main

import (
	"fmt"
	"strconv"
)

type Celsius float64

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 64) + " C"
}

func main() {
	var c Celsius = 20.0
	fmt.Println(c)
}
