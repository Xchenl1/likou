package main

import (
	"fmt"
	"sort"
)

// 我的方法是先从小到大排序 然后看看第一位是不是0 是0的话一次依次查找 找到不是0的数字错误
//
//	func minNumber(nums []int) string {
//		for i := 0; i < len(nums); i++ {
//			for j := 0; j < len(nums)-i-1; j++ {
//				if nums[j] > nums[j+1] {
//					k := nums[j+1]
//					nums[j+1] = nums[j]
//					nums[j] = k
//				}
//			}
//		}
//		s := ""
//		return s
//	}
func minNumber(nums []int) string {
	//使用sort包下的slice函数 非常巧妙
	sort.Slice(nums, func(i, j int) bool { //比如10 2 102<210 所以10<2 所以10在2前边
		x := fmt.Sprintf("%d%d", nums[i], nums[j])
		y := fmt.Sprintf("%d%d", nums[j], nums[i])
		return x < y
	})
	s := ""
	for _, v := range nums {
		s = fmt.Sprintf("%s%d", s, v)
	}
	return s
}

//// 位数
//func WeiShu(num int) int {
//	i := 0
//	for num != 0 {
//		num = num / 10
//		i++
//	}
//	return i
//}

func main() {
	nums := []int{10, 2}
	fmt.Println(minNumber(nums))
}
