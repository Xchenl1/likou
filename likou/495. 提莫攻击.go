package main

import "fmt"

func findPoisonedDuration(timeSeries []int, duration int) int {
	k := 0 //记录最终的中毒时间
	sum := 0
	for _, v := range timeSeries {
		if v > k {
			sum += duration
		} else {
			sum += v + duration - k
		}
		k = v + duration
	}
	fmt.Println(sum)
	return sum
}

//func findPoisonedDuration(timeSeries []int, duration int) (ans int) {
//	expired := 0
//	for _, t := range timeSeries {
//		if t >= expired {
//			ans += duration
//		} else {
//			ans += t + duration - expired
//		}
//		expired = t + duration
//	}
//	return
//}

func main() {
	timeseries := []int{1, 2}
	findPoisonedDuration(timeseries, 2)
}
