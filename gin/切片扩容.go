package main

import "fmt"

func main() {
	a := make([]int, 25, 25)
	//b := make([]int, 256)
	c := make([]int, 24, 24)
	//如果需要扩容的大小<=2*原切片的容量
	a = append(a, c...)
	fmt.Println("容量：", cap(a), "长度：", len(a))
	//表示解压缩切片
	//a = append(a, c...)
}
