package main 

import "fmt"

func main(){
	type person struct{
		name string
		age int
		pet string
	}

	diana := person{
		"Diana",
		21,
		"Dog",
	}

	fmt.Println(diana.pet)
}

