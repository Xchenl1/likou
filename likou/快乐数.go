package main

import (
	"fmt"
	"math"
)

func isHappy(n int) bool {
	//即使在代码中你不需要处理第三种情况（无限循环），你仍然需要理解为什么它永远不会发生，这样你就可以证明为什么你不处理它
	//看了注解发现使用map
	var sum float64
	var n1 map[int]bool
	n1 = make(map[int]bool, 1)
	for n1[n] == false {
		n1[int(sum)] = true
		sum = 0
		for n != 0 {
			sum += math.Pow(float64(n%10), 2)
			n = n / 10
		}
		if sum == 1 {
			return true
		}
		//将sum的值赋给n
		n = int(sum)
	}

	fmt.Println(sum)
	return false
}

func main() {
	fmt.Println(isHappy(3))
}
