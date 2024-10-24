package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	a := map[int]int{}
	s1 := []byte(s)
	t1 := []byte(t)
	//用map解决 但是会出现一个问题 map存数据是方便 但是不好比较
	for k, value := range s1 {
		a[int(value)] = int(t1[k])
		for k1, _ := range s1 {
			if s1[k1] == value && t1[k1] == t1[k] {
				continue
			} else if s1[k1] != value && t1[k1] != t1[k] {
				continue
			} else if s1[k1] == value && t1[k1] != t1[k] {
				return false
			} else if s1[k1] != value && t1[k1] == t1[k] {
				return false
			}
		}
	}
	return true
	//for k1, _ := range a {
	//	if k1 == int(value) && a[k1] == int(t1[k]) {
	//		continue
	//	} else if k1 == int(value) && a[k1] != int(t1[k]) {
	//		return false
	//	} else if k1 != int(value) && a[k1] == a[int(value)] {
	//		return false
	//	}
}

func main() {
	s := "badc"
	k := "baba"
	fmt.Println(isIsomorphic(s, k))
}
