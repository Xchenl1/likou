package main

import "fmt"

//type ListNode struct {
//	Val int
//	Next *ListNode
//}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := Fanzhuan(Shchu(l1))
	s2 := Fanzhuan(Shchu(l2))
	s3 := Add(s1, s2)
	if len(s3) == 1 {
		return &ListNode{int(s3[0]), nil}
	} else if len(s3) > 1 {
		k2 := &ListNode{int(s3[len(s3)-1]), nil}
		k1 := &ListNode{}
		for i := len(s3) - 2; i >= 0; i-- {
			k1 = &ListNode{int(s3[i]), k2}
			k2 = k1
		}
		return k2
	}
	return nil
}

func Cifang(n int64) int64 {
	var i int64 = 1
	var sum int64 = 1
	for i < n {
		sum *= 10
		i++
	}
	return sum
}

// Add 两数组相加
func Add(s1 []int, s2 []int) []int64 {
	//k := math.Pow(10, float64(len(s1))-1)
	//问题出在这
	k := Cifang(int64(len(s1)))
	//不是int64为容易出错
	var sum int64 = 0
	for _, v := range s1 {
		sum += int64(v) * k
		k = k / 10
	}
	//k1 := math.Pow(10, float64(len(s2))-1)
	k1 := Cifang(int64(len(s2)))
	var sum1 int64 = 0
	for _, v := range s2 {
		sum1 += int64(v) * k1
		k1 = k1 / 10
	}
	sum3 := sum1 + sum
	s3 := []int64{}
	if sum3 == 0 {
		s3 = append(s3, sum3)
		return s3
	}
	for sum3 != 0 {
		s3 = append(s3, sum3%10)
		sum3 /= 10
	}
	fmt.Println(s3)
	return s3
}

// Shchu 获取val值
func Shchu(l *ListNode) []int {
	s1 := make([]int, 0)
	if l.Next == nil {
		s1 = append(s1, l.Val)
		return s1
	}
	for l != nil {
		s1 = append(s1, l.Val)
		l = l.Next
	}
	fmt.Println(s1)
	return s1
}

// Fanzhuan 反转数组
func Fanzhuan(s1 []int) []int {
	for i := 0; i < len(s1)/2; i++ {
		k := s1[i]
		s1[i] = s1[len(s1)-i-1]
		s1[len(s1)-i-1] = k
	}
	fmt.Println(s1)
	return s1
}

//func main() {
//	l111 := ListNode{4, nil}
//	l222 := ListNode{6, &l111}
//	l333 := ListNode{5, &l222}
//
//	l1 := ListNode{1, nil}
//	l2 := ListNode{0, &l1}
//	l3 := ListNode{0, &l2}
//	l4 := ListNode{0, &l3}
//	l5 := ListNode{0, &l4}
//	l6 := ListNode{0, &l5}
//	l7 := ListNode{0, &l6}
//	l8 := ListNode{0, &l7}
//	l9 := ListNode{0, &l8}
//	l10 := ListNode{0, &l9}
//	l11 := ListNode{0, &l10}
//	l12 := ListNode{0, &l11}
//	l13 := ListNode{0, &l12}
//	l14 := ListNode{0, &l13}
//	l15 := ListNode{0, &l14}
//	l16 := ListNode{0, &l15}
//	l17 := ListNode{0, &l16}
//	l18 := ListNode{0, &l17}
//	l19 := ListNode{0, &l18}
//	l20 := ListNode{0, &l19}
//	l21 := ListNode{0, &l20}
//	l22 := ListNode{0, &l21}
//	l23 := ListNode{0, &l22}
//	l24 := ListNode{0, &l23}
//	l25 := ListNode{0, &l24}
//	l26 := ListNode{0, &l25}
//	l27 := ListNode{0, &l26}
//	l28 := ListNode{0, &l27}
//	l29 := ListNode{0, &l28}
//	l30 := ListNode{0, &l29}
//	l31 := ListNode{0, &l30}
//	l32 := ListNode{1, &l31}
//	//Add(Fanzhuan(Shchu(&l3)), Fanzhuan(Shchu(&l3)))
//	i := addTwoNumbers1(&l333, &l32)
//	for i != nil {
//		fmt.Print(i.Val)
//		i = i.Next
//	}
//}
