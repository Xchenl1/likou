package main

import "fmt"

func main() {
	var a *int
	fmt.Printf("%T,%v\n", a, a) //初始化为nil 不可以直接赋值
	//*a = 9
	a1 := new(int) //new分配内存  将其初始化为0值
	fmt.Println(a1)
	//*a1 = 9
	fmt.Printf("%T,%v\n", a1, a1)
	fmt.Println(*a1)

	//对于数组
	var l1 *[5]int
	fmt.Printf("%T,%v\n", l1, l1)
	l2 := new([5]int)
	fmt.Printf("%T,%v\n", l2, l2)
	l2[2] = 1
	fmt.Println(*l2)

	//引用类型
	var k1 *[]int
	fmt.Printf("%T,%v\n", k1, k1)
	k2 := new([]int)
	fmt.Println("k2:", k2)
	fmt.Printf("%v", k2)
	*k2 = append(*k2, 1)
	fmt.Printf("%T,%v\n", k2, k2)
	//(*k2)[0] = 1
	fmt.Println(k2)
	k3 := make([]int, 1)
	fmt.Printf("%T,%v\n", k3, k3)
	k3[0] = 1
	fmt.Println(k3)
}
