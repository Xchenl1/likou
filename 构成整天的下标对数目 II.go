package main

//func countCompleteDayPairs(hours []int) int64 {
//	var num int64
//	for i := 0; i < len(hours)-1; i++ {
//		for j := i + 1; j < len(hours); j++ {
//			if (hours[i]+hours[j])%24 == 0 {
//				num++
//			}
//		}
//	}
//	return num
//}

func countCompleteDayPairs(hours []int) int64 {
	var ans int64
	cnt := make(map[int]int)
	for _, hour := range hours {
		// 能与 hour 配对的数
		sh := 24 - hour%24
		// 计算个数
		ans += int64(cnt[sh%24])
		// 累加 cnt
		cnt[hour%24]++
	}
	return ans
}

func main() {
	h := []int{12, 12, 30, 24, 24}
	countCompleteDayPairs(h)
}
