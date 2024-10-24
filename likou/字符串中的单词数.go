package main

import (
	"fmt"
	"strings"
)

func countSegments(s string) int {
	if s == "" {
		return 0
	}
	count := 0
	list := strings.Split(s, " ")
	for _, v := range list {
		if v != "" {
			count++
		}
	}
	return count
}

func main() {
	//s := "         "
	s1 := "Hello, my name is John"
	fmt.Println(countSegments(s1))
}
