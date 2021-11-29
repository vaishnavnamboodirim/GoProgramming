package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	mapp := make(map[string]string)
	var strKey string 
	var strVal string
	fmt.Scanf("%s %s",&strKey,&strVal)
	mapp[strKey] = strVal 
    data, _ := json.Marshal(mapp)
	fmt.Println(string(data))
	
}