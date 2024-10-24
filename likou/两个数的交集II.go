package main

func intersect(nums1 []int, nums2 []int) []int {
	m1 := map[int]int{}
	for _, v := range nums1 {
		m1[v]++
	}
	m2 := []int{}
	for _, v := range nums2 {
		if m1[v] > 0 {
			m2 = append(m2, v)
			m1[v]--
		}
	}
	return m2
}

func main() {

}
