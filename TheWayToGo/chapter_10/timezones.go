package main

import "fmt"

type TZ int

const (
	UTC TZ = iota
	GXY
)

var m = map[TZ]string{
	UTC: "Universal Greenwich time",
	GXY: "Guan Xin Yi",
}

func (t TZ) String() string {
	//return m[t]
	if zone, ok := m[t]; ok {
		return zone
	}
	return " "
}

func main() {
	var t TZ = 0
	fmt.Println(t)
}
