package main

import "fmt"

type base struct{}

func (base) Magic() {
	fmt.Println("base magic")
}

func (self base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func main() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}
