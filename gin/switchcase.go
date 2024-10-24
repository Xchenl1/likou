package main

import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	//默认带break
	switch a {
	case 1:
		fmt.Println("1")
		//fallthrough穿透继续往下输出
	case 2:
		fmt.Println("2")
	}
}
