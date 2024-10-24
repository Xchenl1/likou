package main

import "fmt"

//func Analyse(src []string, dest []string) (same []string, dif []string).

func Analyse(src []string, dest []string) (same []string, dif []string) {
	m1 := make(map[string]struct{})
	srcLen, destLen := len(src), len(dest)
	if srcLen == 0 {
		return same, dest
	}
	if destLen == 0 {
		return same, src
	}

	for _, value := range dest {
		m1[value] = struct{}{}
	}

	for _, value := range src {
		if _, ok := m1[value]; ok {
			same = append(same, value)
			delete(m1, value)
		} else {
			dif = append(dif, value)
		}
	}

	for value, _ := range m1 {
		dif = append(dif, value)
	}
	return
}

// 代码块
func main() {
	src := []string{"a", "b", "c", "d", "e"}
	dest := []string{"a", "c", "e", "f", "g", "1"}
	same, dif := Analyse(src, dest)
	fmt.Println(same, dif)
}
