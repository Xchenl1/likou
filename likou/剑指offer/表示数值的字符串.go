package main

// func isNumber(s string) bool { // 定义函数isNumber，输入为字符串s，输出为bool类型
//
//		s = strings.TrimSpace(s) // 去除字符串s两端的空格，获取新字符串
//		if s == "" {             // 如果新字符串为空，则返回false
//			return false
//		}
//
//		if s[0] == '-' || s[0] == '+' { // 如果新字符串的第一个字符为正负号，则将其去除，并记录字符串的正负性
//			s = s[1:]
//		}
//
//		hasDot := false // 定义布尔变量hasDot，用来标记字符串是否包含小数点
//		hasE := false   // 定义布尔变量hasE，用来标记字符串是否包含指数符号
//		hasNum := false // 定义布尔变量hasNum，用来标记字符串是否包含数字
//
//		for i := 0; i < len(s); i++ { // 遍历字符串s中的每个字符
//			if s[i] >= '0' && s[i] <= '9' { // 如果字符为数字，则标记hasNum为true
//				hasNum = true
//			} else if s[i] == '.' { // 如果字符为小数点
//				if hasDot || hasE { // 如果字符串已经包含小数点或指数符号，则返回false
//					return false
//				}
//				hasDot = true // 否则，标记字符串包含小数点
//			} else if s[i] == 'e' || s[i] == 'E' { // 如果字符为指数符号
//				if hasE || !hasNum { // 如果字符串已经包含指数符号或者前面没有数字，则返回false
//					return false
//				}
//				hasE = true    // 否则，标记字符串包含指数符号
//				hasNum = false // 并将标记hasNum置为false，因为指数符号后面应该跟一个整数
//			} else if s[i] == '-' || s[i] == '+' { // 如果字符为正负号
//				if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' { // 如果不是第一个字符，并且前面的字符不是指数符号，则返回false
//					return false
//				}
//			} else { // 如果字符不是数字、正负号、小数点和指数符号之一，则返回false
//				return false
//			}
//		}
//
//		return hasNum // 如果字符串中包含数字，则返回true，否则返回false
//	}
//func isNumber(s string) bool {
//
//}
//func main() {
//	s := "e"
//	isNumber(s)
//}
