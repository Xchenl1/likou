package main

import "fmt"

func findRepeatNumber(nums []int) int {
	map1 := make(map[int]int, 0)
	for _, v := range nums {
		if map1[v] == 1 {
			return v
		}
		map1[v] = 1
	}
	return -1
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))
}
