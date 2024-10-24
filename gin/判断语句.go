package main

import (
	"fmt"
)

func main() {
	fmt.Printf(" ")
	a := 1
	if a > 0 {
		fmt.Println(a)
	}
	var a1 int
	fmt.Scanln(&a1)
	//if else
	if a > 0 {
		fmt.Println(a)
	} else {
		fmt.Println(a + 1)
	}
	//if else if
	if a > 0 {
		fmt.Println(1)
	} else if a == 0 {
		fmt.Println(0)
	}

	//if简短语句
	if i := 0; i < 10 {
		fmt.Println(i)
		i++
	}
}
