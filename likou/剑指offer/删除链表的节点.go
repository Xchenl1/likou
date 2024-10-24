package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//	func deleteNode(head *ListNode, val int) *ListNode {
//		l1 := &ListNode{}
//		l := head
//		if head.Val == val {
//			head = head.Next
//		}
//		//找到前一个节点  让前一个节点指向下一个节点
//		for l != nil {
//			if l.Val == val {
//				l1.Next = l.Next
//			}
//			l1 = l
//			l = l.Next
//		}
//		return head
//	}
//
//	func main() {
//		var s1 ListNode = ListNode{7, nil}
//		s4 := &ListNode{5, &s1}
//		s3 := &ListNode{6, s4}
//		s2 := &ListNode{7, s3}
//		m := deleteNode(s2, 7)
//		for m != nil {
//			fmt.Println(m.Val)
//			m = m.Next
//		}
//	}
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	k1, k := head, head
	for head != nil {
		if head.Val == val {
			k.Next = head.Next
			break
		}
		k = head
		head = head.Next
	}
	return k1
}

func main() {
	var s1 ListNode = ListNode{7, nil}
	s4 := &ListNode{5, &s1}
	s3 := &ListNode{6, s4}
	s2 := &ListNode{7, s3}
	m := deleteNode(s2, 7)
	for m != nil {
		fmt.Println(m.Val)
		m = m.Next
	}
}
