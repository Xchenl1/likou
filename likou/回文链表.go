package main

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

// 验证回文链表
func isPalindrome2(head *ListNode2) bool {
	if head == nil {
		return true
	}
	//定义map
	a1 := map[int]int{}
	p := head
	i := 1
	for p != nil {
		a1[i] = p.Val
		p = p.Next
		i++
	}
	//先将链表反转
	l := head
	k := &ListNode2{}
	for l != nil {
		a := l.Next
		l.Next = k
		k = l
		l = a
	}
	i = 1
	for k != nil {
		if k.Val == a1[i] {

		} else {
			return false
		}
		k = k.Next
		i++
	}
	return true
}

func main() {
	var s1 ListNode2 = ListNode2{1, nil}
	//var s5 ListNode = ListNode{1, &s1}
	//var s6 ListNode = ListNode{2, &s5}
	s4 := &ListNode2{2, &s1}
	s3 := &ListNode2{2, s4}
	s2 := &ListNode2{1, s3}
	isPalindrome2(s2)
}
