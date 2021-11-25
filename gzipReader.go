package main

import (
	"fmt"
	"os"
	
)

func main(){
	emptyFile, err := os.Create("emptyFile.txt")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(emptyFile)
    emptyFile.Close()
}