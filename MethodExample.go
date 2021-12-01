package main

import (
	"fmt"
)

type Employee struct{
	firstName, lastName string
}

func (e Employee) fullname() string{
	return e.firstName + " " + e.lastName
} 

func main(){
	e := Employee{
		firstName: "Cristiano",
		lastName: "Ronaldo",
	}

	fmt.Println(e.fullname())
}

