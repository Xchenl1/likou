package main

import (
	"fmt"
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	s1 := ""
	for i, _ := range s {
		//fmt.Println(string(s[i]))
		if s[i] == '-' {
			continue
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			//转换大写
			s1 += strings.ToUpper(string(s[i]))
		} else if s[i] >= '0' && s[i] <= '9' {
			s1 += string(s[i])
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			s1 += string(s[i])
		}
	}
	//fmt.Println(s1)
	sum := 0
	k2 := ""
	for i := len(s1) - 1; i >= 0; i-- {
		if sum < k {
			k2 = string(s1[i]) + k2
			sum++
		} else if sum >= k {
			k2 = string(s1[i]) + "-" + k2
			sum = 1
		}
	}
	fmt.Println(k2)
	return k2
}

func main() {
	S := "2-5g-3-J"
	licenseKeyFormatting(S, 2)
}
