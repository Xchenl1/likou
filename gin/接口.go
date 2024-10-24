package main

import "fmt"

type Anminal interface {
	Sing()
	Dump(time int)
	Rap() string
}
type Chicken struct {
	Name string
}

func (c Chicken) Sing() {
	fmt.Println("调用了sing方法")
}
func (c Chicken) Dump(time int) {
	fmt.Println("调用了dump方法")
}
func (c Chicken) Rap() string {
	fmt.Println("调用了rap方法")
	return "rap"
}

func Sing(animnal Anminal) {
	value, ok := animnal.(Chicken)
	fmt.Println(value, ok)
	animnal.Sing()
}
func main() {

	chicken := Chicken{"caixukun"}
	chicken.Rap()
	chicken.Sing()
	chicken.Dump(12)
	Sing(chicken)
}
