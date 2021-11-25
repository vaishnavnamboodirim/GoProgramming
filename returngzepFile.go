package main

import (
	"fmt"
	"os"
	
)

func main(){
	emptyFile, _ := os.Create("emptyFile.txt")
	func createFile(emptyFile) (*gzip.Reader, func(), error) {
		r, err := os.Open(emptyFile)
		if err != nil {
		return nil, nil, err
		}
		gr, err := gzip.NewReader(r)
		if err != nil {
			return nil, nil, err
			}
			return gr, func() {
			gr.Close()
			r.Close()
			}, nil
			}
    
    emptyFile.Close()
}