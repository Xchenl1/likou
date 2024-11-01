package main

import (
	"fmt"
	"sort"
)

func findErrorNums(nums []int) []int {
	sort.Ints(nums)
	num := make([]int, 0)

	mp1 := make(map[int]int)
	for i := 1; i <= len(nums); i++ {
		mp1[i] = 1
	}

	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
		if mp[v] == 2 {
			num = append(num, v)
		}
	}

	for i, _ := range mp1 {
		if _, ok := mp[i]; !ok {
			num = append(num, i)
			break
		}
	}
	return num
}
func main() {
	num := []int{2, 3, 2}
	fmt.Println(findErrorNums(num))
}
