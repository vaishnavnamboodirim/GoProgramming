package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 10)
	for {
		n, err := r.Read(b)
		fmt.Println( n, err, b)
		fmt.Println( b[:n])
		if err == io.EOF {
			break
		}
	}
}
