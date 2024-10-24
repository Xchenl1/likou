package main

//	func twoSum(nums []int, target int) []int {
//		for i := 0; i < len(nums)-1; i++ {
//			for j := i + 1; j < len(nums); j++ {
//				if nums[i]+nums[j] == target {
//					return []int{i, j}
//				}
//			}
//		}
//		return []int{}
//	}
//

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, v := range nums {
		if j, ok := mp[target-v]; ok {
			return []int{i, j}
		}
		mp[v] = i
	}
	return []int{}
}
func main() {

}
