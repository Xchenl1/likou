package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//	func getKthFromEnd(head *ListNode, k int) *ListNode {
//		l := head
//		num := 0
//		for l != nil {
//			num++
//			l = l.Next
//		}
//		num1 := num - k
//		for num1 != 0 {
//			head = head.Next
//			num1--
//		}
//		return head
//	}
//
//	func main() {
//		var s1 ListNode = ListNode{7, nil}
//		s4 := &ListNode{5, &s1}
//		s3 := &ListNode{6, s4}
//		s2 := &ListNode{7, s3}
//		m := getKthFromEnd(s2, 2)
//		for m != nil {
//			fmt.Println(m.Val)
//			m = m.Next
//		}
//	}
func getKthFromEnd(head *ListNode, k int) *ListNode {
	var k11 *ListNode
	L := head
	for i := 1; i <= k; i++ {
		head = head.Next
		k11 = head
	}
	for k11 != nil {
		L = L.Next
		k11 = k11.Next
	}
	return L
}

func main() {
	var s1 = ListNode{7, nil}
	s4 := &ListNode{5, &s1}
	s3 := &ListNode{6, s4}
	s2 := &ListNode{7, s3}
	m := getKthFromEnd(s2, 2)
	for m != nil {
		fmt.Print(m.Val, " ")
		m = m.Next
	}
}
