package module3

import "fmt"

func cyclical_shift(s, t string) int {
	n := len(t)
	tt := t + t
	for i := 0; i < n; i++ {
		if tt[i:i+n] == s {
			return i
		}
	}
	return -1
}

func Task2() {
	var s, t string
	fmt.Scanln(&s)
	fmt.Scanln(&t)
	fmt.Println(cyclical_shift(s, t))
}
