package main

import (
	"fmt"
)

type EmpName struct {
	f_Name, l_Name string
}

func (e *EmpName) ChangeName(newName string) {
	(*e).f_Name = newName
}

func main() {
	e := EmpName{
		f_Name: "Cristiano",
		l_Name: "Ronaldo",
	}

	ep := &e

	ep.ChangeName("GOAT")
	fmt.Println(e)
}
