package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

// 将upPerson的参数修改为不是指针类型，要想达到原来的效果需要添加一个返回值，因为函数是值传递
func upPerson(p Person) Person {
	var newPerson Person
	newPerson.firstName = strings.ToUpper(p.firstName)
	newPerson.lastName = strings.ToUpper(p.lastName)
	return newPerson
}

func main() {
	var p Person
	p.firstName = "Chris"
	p.lastName = "Woodward"
	p = upPerson(p)
	fmt.Printf("The name of the person is %s %s\n", p.firstName, p.lastName)

	// 2—struct as a pointer:
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward" // 这是合法的
	*pers2 = upPerson(*pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	// 3—struct as a literal:
	pers3 := &Person{"Chris", "Woodward"}
	upPerson(*pers3)
}
