package main

import "fmt"

func majorityElement1(nums []int) int {
	k := make(map[int]int, 0)
	for _, v := range nums {
		k[v]++
		if k[v] > len(nums)/2 {
			return v
		}

	}
	return 0
}
func main() {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	fmt.Println(majorityElement1(nums))
}
