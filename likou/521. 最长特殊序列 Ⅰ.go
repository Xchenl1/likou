package main

import (
	"fmt"
)

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	} else if len(a) > len(b) {
		return len(a)
	} else {
		return len(b)
	}
}
func main() {
	a := "aba"
	b := "cdca"
	fmt.Println(findLUSlength(a, b))
}
