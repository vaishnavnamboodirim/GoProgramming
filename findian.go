/*Write a program which prompts the user to enter a string.
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise.
 The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter search string")
	var iStirng string
	fmt.Scanln(&iStirng)
	if strings.Contains(iStirng, "a") || strings.Contains(iStirng, "A") && strings.HasPrefix(iStirng, "i") || strings.HasPrefix(iStirng, "I") && strings.HasSuffix(iStirng, "n") || strings.HasSuffix(iStirng, "N") {
		fmt.Println("Found")
	} else {
		fmt.Printf("Not Found")
	}
}
