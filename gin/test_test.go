package main

import "testing"

func Add11(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	result := Add11(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; expected 5", result)
	}
}
