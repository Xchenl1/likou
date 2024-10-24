package main

import "fmt"

//	func main() {
//		defer func() {
//			fmt.Println("deferred function 1")
//			foo()
//			fmt.Println("deferred function 2")
//			baz()
//		}()
//		fmt.Println("main function")
//	}
//
//	func foo() {
//		fmt.Println("foo start")
//		defer func() {
//			fmt.Println("deferred function in foo")
//			bar()
//		}()
//		fmt.Println("foo end")
//	}
//
//	func bar() {
//		fmt.Println("bar")
//	}
//
//	func baz() {
//		fmt.Println("baz start")
//		defer func() {
//			fmt.Println("deferred function in baz")
//			qux()
//		}()
//		fmt.Println("baz end")
//	}
//
//	func qux() {
//		fmt.Println("qux")
//	}
func deferPrint() (i int) {
	defer func() {
		fmt.Println(i)
		i = 4
	}()
	return 2
}
func main() {
	fmt.Println(deferPrint())
}
