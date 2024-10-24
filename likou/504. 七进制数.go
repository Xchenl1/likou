package main

import (
	"fmt"
	"math"
	"strconv"
)

func convertToBase7(num int) string {
	s1 := ""
	if num < 0 {
		num = int(math.Abs(float64(num))) // 取绝对值
		for num > 0 {
			s1 = strconv.Itoa(num%7) + s1
			num /= 7
		}
		s1 = "-" + s1
	} else if num == 0 {
		return "1"
	} else {
		for num > 0 {
			s1 = strconv.Itoa(num%7) + s1
			num /= 7
		}
	}
	fmt.Println(s1)
	return s1
}
func main() {
	nums := -7
	convertToBase7(nums)
}
