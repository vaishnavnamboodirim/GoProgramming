package main

import "fmt"

func main() {
intSet := map[int]string{}
vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
for _, v := range vals {
intSet[v] = "Eww David!"
}
fmt.Println(len(vals), len(intSet))
fmt.Println(intSet)
}
