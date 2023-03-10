package main

import (
	"fmt"
)

type T struct {
	a int
	b float32
	c string
}

func (t *T) String() string {
	//return strconv.Itoa(t.a) + "/" + strconv.FormatFloat(float64(t.b), 'f', 6, 64) + "/" + "\"" + t.c + "\""
	return fmt.Sprintf("%d / %f / %q", t.a, t.b, t.c)
}

func main() {
	t := T{7, -2.35, "abc\tdef"}
	fmt.Println(t.String())
}
