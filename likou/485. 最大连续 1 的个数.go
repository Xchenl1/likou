package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	num := 0
	max := 0
	for _, v := range nums {
		if v == 1 {
			num++
			//这一步非常重要
			if num > max {
				max = num
			}
		} else if v == 0 {
			num = 0
		}
	}
	return max
}

func main() {
	nums := []int{1, 0, 1, 1, 0, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}
