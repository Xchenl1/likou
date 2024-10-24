package main

import "fmt"

// 将奇数放在偶数前面
func exchange(nums []int) []int {
	s1 := make([]int, 0)
	s2 := make([]int, 0)
	for _, v := range nums {
		//奇数放在前面
		if v%2 != 0 {
			s1 = append(s1, v)
		}
		if v%2 == 0 {
			s2 = append(s2, v)
		}
	}
	for _, v := range s2 {
		s1 = append(s1, v)
	}
	return s1
}

func main() {
	s := []int{1, 2, 3, 4}
	fmt.Println(exchange(s))
}
