package main

import "fmt"

func add(a int, b int) (int, int) {
	return a + b, a - b
}

func main() {
	//注意这里不是地址
	fmt.Println(add(1, 2))
}
