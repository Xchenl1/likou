package main

import "fmt"

func main() {
	//break 只退出当前循环 而break lable 可以退出for所有循环 只能用于for循环
start:
	for i := 0; i < 5; i++ {
		for i > 0 {
			fmt.Println(i)
			i--
			break start //退  出两个for循环
		}
	}
	//goto 但是不能跨函数进行跳跃
	for i := 0; i < 5; i++ {
		for i > 0 {
			i--
			goto start1 //退  出两个for循环
		}
	}
start1:
}
