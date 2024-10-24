package main

import "fmt"

func a1(name *string) {
	//name是指针 *name 表示name所在地址 要想改变name的值 必须要*name
	*name = "123"
	//name是地址 *name是指向该地址的数据
	fmt.Println(name)
	fmt.Println(*name)
}

func main() {
	//fmt.Println()
	var name string
	fmt.Scanln(&name)
	a1(&name)
}
