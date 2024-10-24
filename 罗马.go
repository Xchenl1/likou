package main

import "fmt"

var LuoToNum = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000, "IV": 4, "IX": 9,
	"XL": 40, "XC": 90, "CD": 400, "CM": 900}

func romanToInt(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'I' || s[i] == 'X' || s[i] == 'C' {
			// 判断后一位是否可以拼成key
			j := i
			j++
			if j < len(s) {
				s1 := string(s[i]) + string(s[j])
				if v, ok := LuoToNum[s1]; ok {
					i = j
					num += v
					continue
				}
			}
		}
		// 累加
		num += LuoToNum[string(s[i])]
	}
	return num
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
