package main

// 左移 低位补零  右移高位补零
func hammingWeight(num uint32) int {
	nums := 0
	for num != 0 {
		if num&1 != 0 {
			nums++
		}
		num = num >> 1
	}
	//fmt.Println(nums)
	return nums
}

func main() {
	var s uint32 = 15
	hammingWeight(s)
}
