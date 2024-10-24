package main

import (
	"fmt"
)

func main() {
	var arr []int
	var arr2 [5]int
	fmt.Println(arr[0])
	var arr1 = [3]int{1, 2, 4}
	fmt.Println(arr1)
	arr3 := [...]int{1, 23, 4, 5}
	fmt.Println(arr3)
	//数组的索引 1.遍历数组
	for i, value := range arr1 {
		fmt.Println(i, value)
	}
	//
	arr = append(arr, 9)
	arr = arr[:1] //[0,1)
	fmt.Println(arr2)
}
