package main

import (
	"fmt"
	"strconv"
)

//func translateNum(num int) int {
//	if num < 10 {
//		return 1
//	}
//	if num >= 10 && num <= 25 {
//		return 2
//	}
//	str := strconv.Itoa(num)
//	dp := make([]int, len(str))
//	dp[0] = 1                               //边界
//	if str[:2] >= "10" && str[:2] <= "25" { //边界
//		dp[1] = 2
//	} else {
//		dp[1] = 1
//	}
//	for i := 2; i < len(str); i++ {
//		newnum := str[i-1 : i+1]
//		if newnum >= "10" && newnum <= "25" {
//			dp[i] = dp[i-1] + dp[i-2] //递归关系式
//		} else {
//			dp[i] = dp[i-1]
//		}
//	}
//	return dp[len(str)-1]
//}

func translateNum(num int) int {
	//思路：跟青蛙跳台阶问题出差不多  如果从num[0]开始 有1种可能 ; num[1]开始 如果这两位数>=10并且<=26 说明有两种可能
	//然后一依次类推 向后移动
	if num < 10 {
		return 1
	}
	if num >= 10 && num <= 25 {
		return 2
	} else if num >= 26 && num <= 99 {
		return 1
	}
	num1 := strconv.Itoa(num)
	dp := make([]int, len(num1))
	dp[0] = 1
	for i := 1; i < len(num1); i++ {
		if num1[i-1:i+1] >= "10" && num1[i-1:i+1] <= "25" {
			if i == 1 {
				dp[i] = 2
			} else if i >= 2 {
				dp[i] = dp[i-1] + dp[i-2]
			}
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(num1)-1]
}
func main() {
	num := 12258
	fmt.Println(translateNum(num))
}
