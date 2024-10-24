package main

// 两个字符串每个字符出现的次数相同
func isAnagram(s string, t string) bool {
	s1 := []byte(s)
	s2 := []byte(t)
	s11 := map[byte]int{}
	for _, v := range s1 {
		s11[v] += 1
	}
	s22 := map[byte]int{}
	for _, v := range s2 {
		s22[v] += 1
	}
	if len(s1) > len(s2) {
		for _, v := range s1 {
			if s11[v] != s22[v] {
				return false
			}
		}
	} else {
		for _, v := range s2 {
			if s11[v] != s22[v] {
				return false
			}
		}
	}
	return true
}

func main() {
	s := "ab"
	t := "a"
	isAnagram(s, t)
}
