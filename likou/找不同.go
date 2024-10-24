package main

import "fmt"

func findTheDifference(s string, t string) byte {
	s1 := make(map[int32]int)
	//将map相同的数存入key中 然后去统计
	for _, v := range s {
		s1[v]++
	}
	for i, v := range t {
		if _, ok := s1[v]; ok && s1[v] != 0 {
			s1[v]--
		} else if ok == false || s1[v] == 0 {
			return t[i]
		}
	}
	return 0
}

func main() {
	s := "a"
	y := "aa"
	fmt.Println(findTheDifference(s, y))
}
