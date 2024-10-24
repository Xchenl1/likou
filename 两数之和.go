package main

func twoSum(nums []int, target int) []int {
	for i := 0; i <= len(nums)-2; i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	twoSum(nums, 9)
}
