package main

import (
	"fmt"
)

//	func TestS(t *testing.T) {
//		s1 := make([]int, 2, 3)
//		s2 := make([]int, 2, 3)
//		f1(s1)
//		f2(s2)
//		fmt.Println("m:", s1, s2)
//	}
func f1(s []int) {
	s[0] = 1
	s = append(s, []int{1, 2, 3, 4}...) //触发扩容
	s[0] = 0                            //扩容之后操作元素原切片s1不受影响，因为底层数组已经不是原切片数组
	fmt.Println("f1:", s)
}

func f2(s []int) {
	s[0] = 1
	s = append(s, 1) //没有扩容
	s[0] = 0         //没有扩容,原切片s2受影响
	fmt.Println("f2:", s)
}
func main() {
	s1 := make([]int, 2, 3)
	s2 := make([]int, 2, 3)
	fmt.Println(s1, s2)
	f1(s1)
	f2(s2)
	fmt.Println("m:", s1, s2)
}
