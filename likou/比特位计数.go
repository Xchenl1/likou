package main

import "fmt"

func countBits(n int) []int {
	var nums []int
	for i := 0; i <= n; i++ {
		var k int = 0
		l := i
		for {
			if l%2 == 1 {
				k++
			}
			l = l / 2
			if l == 1 {
				k++
				nums = append(nums, k)
				break
			} else if l == 0 {
				nums = append(nums, k)
				break
			} else {
				continue
			}
		}
	}
	return nums
}

func main() {
	fmt.Println(countBits(5))
}
