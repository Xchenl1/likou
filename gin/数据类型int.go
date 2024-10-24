package main

import "fmt"

var (
	a int8
	//范围是-127——128
	b1 uint8
	//范围0——255
	c int32 = 123123
	//int32 范围-2147483648--2147483647

)

func main() {
	fmt.Println(a, b1, c)
}
