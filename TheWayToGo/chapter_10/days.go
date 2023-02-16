package main

import (
	"fmt"
	"strconv"
)

const (
	_ Day = iota
	MO
	TU
	WE
	TH
	FI
	SA
	SU
)

var week = []string{"MO", "TU", "WE", "TH", "FI", "SA", "SU"}

type Day int

func (d Day) String() string {
	return "the day " + strconv.Itoa(int(d)) + "is " + week[d-1]
}

func main() {
	var d Day = 1
	fmt.Println(d)
	var day = SU
	fmt.Println(day)
}
