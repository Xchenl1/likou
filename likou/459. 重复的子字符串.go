package main

import (
	"fmt"
)

// 第一种解法
//func repeatedSubstringPattern(s string) bool {
//	//这是一个技巧 如果两个字符串拼接起来 可以找到其中一个字符串(去除第一个和最后一个字符) 那么就是有字符串
//	if len(s) == 0 || len(s) == 1 {
//		return false
//	}
//	s1 := s + s
//	s2 := s1[1 : len(s1)-1]
//	fmt.Println(s2)
//	return strings.Contains(s1, s)
//}

// 第二种解法 主要是kmp算法
// 前缀表（不减一）的代码实现
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	j := 0
	next := make([]int, n)
	next[0] = j
	for i := 1; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	// next[n-1]  最长相同前后缀的长度
	//next[len - 1] = 7，next[len - 1] + 1 = 8，8就是此时字符串asdfasdfasdf的最长相同前后缀的长度。
	//(len - (next[len - 1] + 1)) 也就是： 12(字符串的长度) - 8(最长公共前后缀的长度) = 4， 4正好可以被 12(字符串的长度) 整除，所以说明有重复的子字符串（asdf）
	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}
	return false
}

func main() {
	s := "ababab"
	fmt.Println(repeatedSubstringPattern(s))
}
