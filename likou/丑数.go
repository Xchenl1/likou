package main

import "fmt"

func isUgly(n int) bool {
	//if n <= 0 {
	//	return false
	//}
	//if n == 1 {
	//	return true
	//}
	////循环判断从2开始到n-1
	//for i := 2; i <= n; i++ {
	//	//取余为0说明是余数 但不知道是不是质余数
	//	if n%i == 0 {
	//		//进一步判断 质数只能被1和它本身整除
	//		for i1 := 2; i1 < i; i1++ {
	//			//是质数i1只能是2 若不是不用判断
	//			if i%i1 == 0 {
	//				goto e
	//			}
	//		}
	//		if i == 2 {
	//			continue
	//		} else if i == 3 {
	//			continue
	//		} else if i == 5 {
	//			continue
	//		} else {
	//			return false
	//		}
	//	}
	//e:
	//}
	//return true
	var factors = []int{2, 3, 5}
	for _, f := range factors{
		for n%f == 0{
		n /= f
	}
	}
		return n == 1

}

// 丑数
func main() {
	s := 6
	fmt.Println(isUgly(s))
}
