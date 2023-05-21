package module3

import "fmt"

func z_func(s string) []int {
	n := len(s)
	z := make([]int, n)
	z[0] = n
	l := 0
	r := 0
	for i := 1; i < n; i++ {
		if i <= r {
			a := r - i + 1
			b := z[i-l]
			if a < b {
				z[i] = a
			} else {
				z[i] = b
			}
		}
		for i+z[i] < n && s[z[i]] == s[z[i]+i] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}

func Task4() {
	var s string
	fmt.Scanln(&s)
	Z := z_func(s)
	min := Z[0]
	for i := 1; i < len(Z); i++ {
		if Z[i] > 0 && Z[i] < min {
			min = Z[i]
		}
	}
	fmt.Println(min)
}
