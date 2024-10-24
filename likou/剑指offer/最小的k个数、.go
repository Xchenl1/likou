package main

import "fmt"

// 改造冒泡
func getLeastNumbers(arr []int, k int) []int {
	j1 := k
	for i := 0; i < len(arr); i++ {
		if k == 0 {
			return arr[len(arr)-j1:]
		}
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				k1 := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = k1
			}
		}
		k--
	}
	return arr
}

func main() {
	nums := []int{0, 0, 2, 3, 2, 1, 1, 2, 0, 4}
	fmt.Println(getLeastNumbers(nums, 10))
}
