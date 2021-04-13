package main

import (
	"fmt"
)

func main() {
	n := 8
	k := 5
	fmt.Println(constructArray(n, k))
}
func constructArray(n int, k int) []int {
	l := []int{}
	for x := 1; x < n-k; x++ {
		l = append(l, x)
	}
	fmt.Println(l)
	for i := 0; i < k+1; i++ {
		if i%2 == 0 {
			l = append(l, n-k+i/2)
			fmt.Println(l, n-k+i/2)
		} else {
			l = append(l, n-i/2)

		}
	}
	return l
}
