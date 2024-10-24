package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	k := make(map[int]int, 0)
	for _, v := range nums {
		k[v]++
	}
	k1 := make([]int, 0)
	i := 1
	for i <= len(nums) {
		if _, ok := k[i]; ok {
		} else {
			k1 = append(k1, i)
		}
		i++
	}
	return k1
}
func main() {
	nums := []int{1, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
