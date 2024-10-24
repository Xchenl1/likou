package main

import "fmt"

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
var nums = 1

func guess(num int) int {
	if num > nums {
		return -1
	} else if num == nums {
		return 0
	} else if num < nums {
		return 1
	}
	return 0
}

//func guessNumber(n int) int {
//	//这种方法会超出内存
//	//var guessnumber1 func(n int) int
//	//key := 0
//	//guessnumber1 = func(n int) int {
//	//	if guess(n) == -1 {
//	//		guessnumber1(n - 1)
//	//	}
//	//	if guess(n) == 1 {
//	//		guessnumber1(n + 1)
//	//	}
//	//	if guess(n) == 0 {
//	//		key = n
//	//		return key
//	//	}
//	//	return 0
//	//}
//	//guessnumber1(n)
//	//return key
//}

// 还是不行
//
//	func guessNumber(n int) int {
//		key := 0
//		for {
//			if guess(n) == -1 {
//				n = n - 1
//			}
//			if guess(n) == 0 {
//				key = n
//				break
//			}
//			if guess(n) == 1 {
//				n = n + 1
//			}
//		}
//		return key
//	}
//
// 变相考察二分查找
func guessNumber(n int) int {
	left := 1
	right := n
	key := 0
	for left <= right {
		mid := (left + right) / 2
		if guess(mid) == -1 {
			right = mid - 1
		} else if guess(mid) == 0 {
			key = mid
			break
		} else if guess(mid) == 1 {
			left = mid + 1
		}
	}
	return key
}
func main() {
	fmt.Println(guessNumber(1))
}
