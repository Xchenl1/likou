package main

import "fmt"

type ll int

func CS() interface{} {
	return 1
}
func main() {
	s := CS()
	switch s.(type) {
	case interface{}:
		fmt.Println("接口")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("无")
	}
	//不能进行赋值或者取值操作没有初始化 m为空
	//使用make对其进行初始化 可以操作
	var m map[string]int
	//m[""] = 1
	fmt.Println(m[""])
	m1 := make(map[string]int)
	m1["1"] = 1
	delete(m1, "1")
	//
	//m2 := new(map[string]int)
}
