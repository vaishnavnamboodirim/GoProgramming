package main

import "fmt"

func main(){
	evenVals := []int{2,4,6,8,10,12}

	for i,v := range evenVals{
		fmt.Println(i,v)
	}

	strVals := map[string]bool {"Fred":false, "McTominey":true,"Van De Beek":true}
    for k := range strVals{
		fmt.Println(k)
	}
}

