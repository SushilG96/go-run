package main

import (
	"fmt"
)

func main() {
	A := []int{1, 2, 0}
	fmt.Println(isIdealPermutation(A))
}
func isIdealPermutation(A []int) bool {
	for i := 0; i < len(A); i++ {
		if abs(A[i], i) > 1 {
			return false
		}
	}
	return true
}
func abs(a int, i int) int {
	cal := a - i
	if cal > 0 {
		return cal
	}
	return -cal
}
