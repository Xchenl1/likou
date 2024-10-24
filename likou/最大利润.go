package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	//sum := 0
	//for i := 0; i < len(prices); i++ {
	//	for j := i; j < len(prices); j++ {
	//		if prices[i]-prices[j] > 0 {
	//			continue
	//		} else if prices[i]-prices[j] <= 0 && prices[j]-prices[i] > sum {
	//			sum = prices[j] - prices[i]
	//		}
	//	}
	//}
	//return sum
	//minValue := math.MaxInt64
	//fmt.Println(minValue)
	//数学思维是这样：我可以先记录数组第一个元素的大小
	//依次遍历 如果后边数据减去这个数据>0 那么说明有利润
	//但是会出现这种情况 后边的进价会更低 那么出现的利润一定会比前边进价的利润高
	var max int
	min := 10000
	for i, _ := range prices {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}

func main() {
	a := []int{7, 1, 6, 13, 6, 5, 4}
	fmt.Println(maxProfit(a))
}
