package main

import (
	"strconv"
)

var nums1 = map[int]string{10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f"}

func toHex(num int) string {
	if num < 0 {
		num += 4294967296
	}
	s1 := ""
	k := make([]int, 0)
	if num >= 0 {
		for num > 0 {
			k = append(k, num%16)
			num /= 16
		}

		for i := len(k) - 1; i >= 0; i-- {
			if i1, v := nums1[k[i]]; v {
				s1 += i1
			} else {
				s1 += strconv.Itoa(k[i])
			}
		}
	}
	return s1
}
func main() {
	toHex(26)
}
