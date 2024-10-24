package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	//s1 := make(map[rune]int)
	s2 := make(map[int32]int)
	//k1 := []rune(ransomNote)
	k2 := []int32(magazine)
	//for _, v := range k1 {
	//	s1[v]++
	//}
	for _, v := range k2 {
		s2[v]++
	}
	for _, v := range ransomNote {
		if s2[v] > 0 {
			s2[v]--
		} else if s2[v] <= 0 {
			return false
		} else if _, ok := s2[v]; ok == false {
			return false
		}
	}
	return true
}

func main() {
	ransomNote := "a"
	magazine := "b"
	fmt.Println(canConstruct(ransomNote, magazine))
}
