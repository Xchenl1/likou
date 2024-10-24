package main

// 奇数位于偶数前边
func exchange(nums []int) []int {
	s1 := []int{}
	for _, v := range nums {
		if v%2 == 1 {
			s1 = append(s1, v)
		}
	}
	for _, v := range nums {
		if v%2 == 0 {
			s1 = append(s1, v)
		}
	}
	return s1
}

func main() {
	nums := []int{1, 2, 3, 4}
	exchange(nums)
}
