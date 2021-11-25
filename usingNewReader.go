package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	
)


func main(){
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write([]byte("Hello World!"))
	w.Close()

	r,err := gzip.NewReader(&b)
	fmt.Println(r,w,err)
}