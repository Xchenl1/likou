package main

import "fmt"

//func reverseLeftWords(s string, n int) string {
//	s1 := []int32(s)
//	slice := s1[:n]
//	s1 = s1[n:]
//	//取出前面的字符串
//	slice1 := string(slice)
//	s2 := string(s1)
//	result := s2 + slice1
//	return result
//}

func reverseLeftWords(s string, x int) string {
	s1 := []rune(s)
	s2 := s1[:x]
	s1 = s1[x:]
	return string(s1) + string(s2)
}

func main() {
	s := "wehahah"
	fmt.Println(reverseLeftWords(s, 2))
}
