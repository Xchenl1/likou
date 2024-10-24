package main

import "fmt"

func main() {
	//四种for循环
	//第一种
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//第二种 go语言没有while循环语句 可以使用for循环代替
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
}
