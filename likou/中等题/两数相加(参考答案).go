package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

func main() {
	l111 := ListNode{4, nil}
	l222 := ListNode{6, &l111}
	l333 := ListNode{5, &l222}

	l1 := ListNode{1, nil}
	l2 := ListNode{0, &l1}
	l3 := ListNode{0, &l2}
	l4 := ListNode{0, &l3}
	l5 := ListNode{0, &l4}
	l6 := ListNode{0, &l5}
	l7 := ListNode{0, &l6}
	l8 := ListNode{0, &l7}
	l9 := ListNode{0, &l8}
	l10 := ListNode{0, &l9}
	l11 := ListNode{0, &l10}
	l12 := ListNode{0, &l11}
	l13 := ListNode{0, &l12}
	l14 := ListNode{0, &l13}
	l15 := ListNode{0, &l14}
	l16 := ListNode{0, &l15}
	l17 := ListNode{0, &l16}
	l18 := ListNode{0, &l17}
	l19 := ListNode{0, &l18}
	l20 := ListNode{0, &l19}
	l21 := ListNode{0, &l20}
	l22 := ListNode{0, &l21}
	l23 := ListNode{0, &l22}
	l24 := ListNode{0, &l23}
	l25 := ListNode{0, &l24}
	l26 := ListNode{0, &l25}
	l27 := ListNode{0, &l26}
	l28 := ListNode{0, &l27}
	l29 := ListNode{0, &l28}
	l30 := ListNode{0, &l29}
	l31 := ListNode{0, &l30}
	l32 := ListNode{1, &l31}
	//Add(Fanzhuan(Shchu(&l3)), Fanzhuan(Shchu(&l3)))
	i := addTwoNumbers(&l333, &l32)
	for i != nil {
		fmt.Print(i.Val)
		i = i.Next
	}
}
