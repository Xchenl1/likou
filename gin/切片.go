package main

import (
	"fmt"
)

func main() {
	//切片是一种特殊的数据类型
	var sc = [5]int{1, 2, 3, 4, 5}
	var slice = sc[0:2] //左闭右开
	//如果改变切片的地址 那么原来的数组也会改变
	fmt.Println(slice)
	//切片类型
	var s []int
	fmt.Println(s == nil)
	s = make([]int, 3)
	fmt.Println(s)
	//var a []int = make([]int, 6)
	s1 := []int{1, 2}
	fmt.Println(s1)
	//追加切片
	var s2 []int = []int{1, 2, 4}
	s2 = append(s2, 5)
	fmt.Println(s2)
	//赋值数组 如果不够则会 一次赋值最前面的数据
	var s3 = make([]int, 3)
	copy(s3, s2)
	fmt.Println(s3)
	str := "asd"
	fmt.Printf("%c", []byte(str))
	for i, v := range str {
		fmt.Printf("%d %c %c", i, str[i], v)
	}
}
