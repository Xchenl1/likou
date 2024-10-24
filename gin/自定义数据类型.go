package main

import "fmt"

type msgType uint16

func main() {
	var sunccess msgType = 1000
	var errormsg msgType = 2000
	var num uint16
	sunccess = msgType(num)
	fmt.Println(sunccess, errormsg, num)
	//byte是uint8的别名
	var a1 byte = 0
	var o1 uint8 = 0
	fmt.Println(a1 == o1)
	//rune是int32的别名
	var i1 rune = 11
	var i2 int32 = 11
	fmt.Println(i1 == i2)
	//
}
