package main

import "fmt"

func main() {
	a := 1
	b := 2
	defer calc(a, calc(a, b, "0"), "1") //1 3 1
	a = 0
	defer calc(a, calc(a, b, "3"), "2")
}
func calc(x, y int, s string) int {
	fmt.Println(s)
	fmt.Println(x, y, x+y)
	return x + y
}

//结果：
//0
//1 2 3
//3
//0 2 2
//2
//0 2 2
//1
//1 3 4
