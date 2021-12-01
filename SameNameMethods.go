package main

import (
	"fmt"
)

type CallNameENG struct{
	fName, lName string
	
}

type CallNameESP struct{
	fName, lName string
}

func (e CallNameENG) CallName() string{
	return e.fName + " " + e.lName 
} 

func (e CallNameESP) CallName() string{
	return e.fName + " " + e.lName 
} 

func main(){
	eng := CallNameENG{
		fName: "Jammie",
		lName: "Vardy",
	}

	esp := CallNameESP{
		fName: "Lionel",
		lName: "Messi",
	}

	fmt.Println(eng.CallName())
	fmt.Println(esp.CallName())
}
