package main

import "fmt"

func minArray(numbers []int) int {
	j := numbers[0]
	for _, v := range numbers {
		if v <= j {
			j = v
		}
	}
	return j
}

func main() {
	nums := []int{1, 3, 5}
	fmt.Println(minArray(nums))
}
