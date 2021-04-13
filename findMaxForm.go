package main

import (
	"fmt"
)

func main() {
	strs := []string{"10", "1", "0", "0001", "111001"}
	m := 5
	n := 3

	fmt.Println(findMaxForm(strs, m, n))
}
func findMaxForm(strs []string, m int, n int) int {
	res := 0
	for _, y := range strs {
		zero := 0
		one := 0
		for _, x := range y {
			if x == 48 {
				zero++
			} else {
				one++
			}
		}
		if (m > zero && n > one) || m >= zero && one == 0 || zero == 0 && n >= one {
			fmt.Println(zero, one)
			res++
		}
	}

	return res
}
