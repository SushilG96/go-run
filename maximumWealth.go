package main

func maximumWealth(accounts [][]int) int {
	m := -1
	for _, x := range accounts {
		tmp := sum(x)
		if tmp > m {
			m = tmp
		}
	}
	return m
}

func sum(w []int) int {
	s := 0
	for _, i := range w {
		s = s + i
	}
	return s
}
