package main

func missingNumber(nums []int) int {
	l := len(nums)
	s := 0
	for _, i := range nums {
		s += i
	}
	return (l*(l+1))/2 - s
}
