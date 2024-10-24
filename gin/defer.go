package main

import "fmt"

func add1(s string) string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic("有误！！！")
	return s
}
func main() {
	//defer函数会立即求值 相当于栈 先进后出
	fmt.Printf("asd\n")
	go add1("1234")
	fmt.Println(123)
	defer fmt.Println("12312313")
	//捕获异常机制
	//init函数会先于main函数执行 常用于初始化
}
