package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	type Node struct {
		fName string
		lName string
	}

	fmt.Println("Enter the full file path")
	var inputFileName string
	fmt.Scanf("%s", inputFileName)

	file, err := os.Open(inputFileName)
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", inputFileName, err))
	}

	if err != nil {
		log.Fatal("failed to open")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	nodes := make([]Node, 0, 3)
	var aNode Node

	for i := 0; scanner.Scan(); i++ {
		s := strings.Split(scanner.Text(), " ")
		aNode.fName, aNode.lName = s[0], s[1]
		nodes = append(nodes, aNode)

		fmt.Println(nodes[i])
	}

}
