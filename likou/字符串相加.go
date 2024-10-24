package main

//func addStrings(num1 string, num2 string) string {
//	if num1 == "0" && num2 == "0" {
//		return "0"
//	}
//	s1 := ZhI(num1)
//	s2 := ZhI(num2)
//	s3 := s1 + s2
//	s := ""
//	k := make([]int, 0)
//	for s3 > 0 {
//		k = append(k, s3%10)
//		s3 /= 10
//	}
//	for i := len(k) - 1; i >= 0; i-- {
//		s += string(k[i] + '0')
//	}
//	return s
//}
//func ZhI(nums string) int {
//	var sum1 float64 = 0
//	k1 := len(nums)
//	for _, v := range nums {
//		sum1 += float64(v-'0') * math.Pow(float64(10), float64(k1)-1)
//		k1--
//	}
//	return int(sum1)
//}

//func main() {
//	fmt.Println(addStrings("9333852702227987", "85731737104263"))
//}
