package main

import "fmt"

func longestPalindrome(s string) int {
	k := make(map[byte]int, 0)
	for i, _ := range s {
		k[s[i]]++
	}
	num := 0
	l1 := false
	for _, v1 := range k {
		if v1%2 == 0 {
			num += v1
		} else {
			l1 = true
			num += v1 - 1
		}
	}
	if l1 == true {
		return num + 1
	}
	return num
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}
