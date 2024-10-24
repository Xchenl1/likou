package test1

import (
	"fmt"
	"testing"
)

func Add11(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	result := Add11(2, 2)
	fmt.Println(result)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; expected 5", result)
	}
}

func Test(t *testing.T) {
	
}
