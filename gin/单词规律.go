package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	//按照空格将字符串分解成数组
	s1 := strings.Split(s, " ")
	word2ch := map[byte]string{}
	ch2word := map[string]byte{}
	if len(pattern) != len(s1) {
		return false
	}
	for i, v := range s1 {
		l := pattern[i]
		if word2ch[l] != v && word2ch[l] != "" || ch2word[v] != l && ch2word[v] > 0 {
			return false
		}
		word2ch[l] = v
		ch2word[v] = l
	}
	return true
}

func main() {
	s := "dog dog dog dog"
	pattern := "abba"
	fmt.Println(wordPattern(pattern, s))
}
