package main

import "fmt"

func main() {
	fmt.Print("你好：")
	fmt.Printf("nihao\n")
	fmt.Println("asda")
	fmt.Print("我的年龄是" + "22" + "岁" + "\n")
	fmt.Printf("我的年龄是%d岁\n", 22)
	//%o:8进制 %b:二进制 %d:10进制 %T:类型 %v:值
	fmt.Printf("%b %d %o %O %x %X\n", 10, 10, 10, 10, 10, 10)
	fmt.Printf("%.2f\n", 2.198)
}
