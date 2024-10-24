package main

func findWords(words []string) []string {
	str := make([]string, 0)
	s := "qwertyuiop"
	k := "asdfghjkl"
	l := "zxcvbnm"
	k1, k2, k3 := make(map[int]bool, 0), make(map[int]bool, 0), make(map[int]bool, 0)
	for _, v := range s {
		k1[int(v)] = true
		k1[int(v-32)] = true
	}
	for _, v := range k {
		k2[int(v)] = true
		k2[int(v-32)] = true
	}
	for _, v := range l {
		k3[int(v)] = true
		k3[int(v-32)] = true
	}
	s1 := 0
	s2 := 0
	s3 := 0
	for _, v := range words {
		for _, v1 := range v {
			if _, ok := k1[int(v1)]; ok {
				s1++
			} else if _, ok := k2[int(v1)]; ok {
				s2++
			} else {
				s3++
			}
			if s1 != 0 && s2 == 0 && s3 == 0 {
			} else if s2 != 0 && s1 == 0 && s3 == 0 {
			} else if s3 != 0 && s1 == 0 && s2 == 0 {
			} else {
				s1, s2, s3 = 0, 0, 0
				goto init
			}
		}
		str = append(str, v)
		s1, s2, s3 = 0, 0, 0
	init:
	}
	return str
}

func main() {
	word := []string{"Hello", "Alaska", "Dad", "Peace"}
	findWords(word)
}
