package main

import (
	"fmt"
	"strings"
)

//	func reverseWords(s string) string {
//		var str string
//		//s = strings.TrimSpace(s)
//		i := len(s) - 1
//		var j int
//		for i >= 0 {
//			for i >= 0 && s[i] == ' ' {
//				i--
//			}
//			j = i
//			for i >= 0 && s[i] != ' ' {
//				i--
//			}
//			str += s[i+1:j+1] + " "
//		}
//		return strings.TrimSpace(str)
//	}
//

// 双指针法
func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	i, j := len(s)-1, 0
	s1 := ""
	for i >= 0 {
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i
		for i >= 0 && s[i] != ' ' {
			i--
		}
		s1 += s[i+1:j+1] + " "
	}
	return strings.TrimSpace(s1)
}

func main() {
	s1 := "a good   example"
	//s2 := "the sky is blue"
	//fmt.Println(reverseWords(s2))
	fmt.Println(reverseWords(s1))
}
