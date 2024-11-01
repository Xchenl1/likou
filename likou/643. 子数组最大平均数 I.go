package main

//// 时间复杂度超了
//func findMaxAverage(nums []int, k int) float64 {
//	average := float64(-1 << 64)
//	var num int
//	fmt.Println(num)
//	for i := 0; i <= len(nums)-k; i++ {
//		for j := 0; j < k; j++ {
//			num += nums[i+j]
//		}
//		avg := float64(num) / float64(k)
//		if avg > average {
//			average = avg
//		}
//		num = 0
//	}
//	return average
//}

// 时间复杂度超了 该为滑动窗口
func findMaxAverage(nums []int, k int) float64 {
	average := float64(-1 << 64)
	sum := 0
	average1 := float64(-1 << 64)
	// 先统计前元素 然后依次减去最左边元素，添加最右边元素
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	average = float64(sum) / float64(k)
	for j := k; j < len(nums); j++ {
		sum += nums[j] - nums[j-k]
		average1 = float64(sum) / float64(k)
		if average1 > average {
			average = average1
		}
	}
	return average
}

func main() {
	num := []int{12231}
	findMaxAverage(num, 1)
}
