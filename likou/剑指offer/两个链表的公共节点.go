package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 链表必须有原结构
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	s1 := map[*ListNode]bool{}
	for headA != nil {
		s1[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if s1[headB] == true {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func main() {
	var s1 ListNode = ListNode{8, nil}
	s4 := &ListNode{7, &s1}
	s3 := &ListNode{6, s4}
	s2 := &ListNode{5, s3}
	//var s11 ListNode = ListNode{1, nil}
	s22 := &ListNode{1, s4}
	m := getIntersectionNode(s2, s22)
	for m != nil {
		fmt.Println(m.Val)
		m = m.Next
	}
}
