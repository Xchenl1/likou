package main

import (
	"fmt"
	"strings"
)

//	func detectCapitalUse(word string) bool {
//		if len(word) == 1 {
//			return true
//		}
//		if word[0] >= 'a' && word[0] <= 'z' {
//			for _, v := range word {
//				if v >= 'A' && v <= 'Z' {
//					return false
//				}
//			}
//		} else if word[0] >= 'A' && word[0] <= 'Z' {
//			word = word[1:]
//			if word[0] >= 'a' && word[0] <= 'z' {
//				for _, v := range word {
//					if v >= 'A' && v <= 'Z' {
//						return false
//					}
//				}
//			} else {
//				word = word[1:]
//				if word[0] >= 'A' && word[0] <= 'Z' {
//					for _, v := range word {
//						if v >= 'a' && v <= 'z' {
//							return false
//						}
//					}
//				}
//			}
//		}
//		return true
//	}
func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}
	//fmt.Println(strings.ToUpper(word))
	//fmt.Println(strings.ToLower(word))
	if strings.ToUpper(word) == word || strings.ToLower(word) == word {
		return true
	}
	if strings.ToUpper(string(word[0])) == string(word[0]) && strings.ToLower(word[1:]) == word[1:] {
		return true
	}
	return false
}
func main() {
	word := "FFFf"
	fmt.Println(detectCapitalUse(word))
}
