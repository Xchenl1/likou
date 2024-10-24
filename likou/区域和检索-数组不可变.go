package main

import "fmt"

type NumArray struct {
	Nums []int
}

func Constructor1(nums []int) NumArray {
	return NumArray{
		Nums: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	var sum int = 0
	for left <= right {
		sum += this.Nums[left]
		left++
	}
	return sum
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor1(nums)
	fmt.Println(obj.SumRange(2, 2))
}
