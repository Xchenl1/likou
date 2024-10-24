package main

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//	func maxDepth(root *TreeNode) int {
//		//采用匿名函数的方式
//		var a func(node *TreeNode)
//		var sum int
//		var sum1 int
//		a = func(node *TreeNode) {
//			if node == nil {
//				return 0
//			}
//			sum1++
//			if sum1 > sum {
//				sum = sum1
//			}
//			left := a(root.Left)
//			right := a(root.Right)
//			sum1--
//			return max(left, right) + 1
//		}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//queue := []*TreeNode{}
	//queue = append(queue, root)
	//ans := 0
	//for len(queue) > 0 {
	//	sz := len(queue)
	//	for sz > 0 {
	//		node := queue[0]
	//		queue = queue[1:]
	//		if node.Left != nil {
	//			queue = append(queue, node.Left)
	//		}
	//		if node.Right != nil {
	//			queue = append(queue, node.Right)
	//		}
	//		sz--
	//	}
	//	ans++
	//}
	//return ans

	var a []*TreeNode
	sum := 0
	a = append(a, root)
	for len(a) > 0 {
		//var b = a[0]
		//var c = a[1:]
		var k = len(a)
		for k > 0 {
			b := a[0]
			a = a[1:]
			if b.Left != nil {
				a = append(a, b.Left)
			}
			if b.Right != nil {
				a = append(a, b.Right)
			}
			k--
		}
		sum++
	}
	return sum
}
func main() {
	//s4 := &TreeNode{1, nil, nil}
	//s3 := &TreeNode{3, nil, nil}
	//s2 := &TreeNode{2, nil, s3}
	//s1 := &TreeNode{1, s4, s2}
	//fmt.Println(maxDepth(s1))
	//s33 := &TreeNode{3, nil, nil}
	//	s22 := &TreeNode{1, nil, nil}
	//	s11 := &TreeNode{1, nil, s22}
	//	fmt.Printf("%v", isSameTree(s1, s11))
}
