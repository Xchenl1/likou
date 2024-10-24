package main

import "fmt"

func main() {
	//
	var c0 uint = 65
	//btye 可以说是 uint的别名
	var c1 byte = 65
	var c2 = 'a'
	var c3 rune = '中'
	fmt.Printf("%v,%c,%T \n", c0, c0, c0)
	fmt.Printf("%v,%c,%T \n", c1, c1, c1)
	fmt.Printf("%v,%c,%T \n", c2, c2, c2)
	fmt.Printf("%v,%c,%T \n", c3, c3, c3)
}
