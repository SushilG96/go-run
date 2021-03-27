package main

import (
	"fmt"
)

func main() {
	digit := "23"
	fmt.Println(letterCombinations(digit))
}
func letterCombinations(digits string) []string {
	res := []string{}
	d := map[rune]string{
		50: "abc",
		51: "def",
		52: "ghi",
		53: "jkl",
		54: "mno",
		55: "pqrs",
		56: "tuv",
		57: "wxyz",
	}
	fmt.Println(d)
	final := []string{""}
	fmt.Println(final)
	if len(digits) == 0 {
		return (res)
	}

	for _, digit := range digits {
		fmt.Println(digit)
		curr := []string{}
		for _, letter := range d[digit] {
			fmt.Println(string(letter))
			for _, pre := range final {
				curr = append(curr, pre+string(letter))
			}
		}
		final = curr
	}
	fmt.Println(final)
	return res
}
