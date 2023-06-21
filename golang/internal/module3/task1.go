package module3

import "fmt"

func get_hash(s string, n int, p float64, x float64) float64 {
	res := 0.0
	for i := 0; i < n; i++ {
		res = (res*x + float64(s[i]))
		res = float64(int64(res) % int64(p))
	}
	return res
}

func rabin_karp(s, t string, p float64, x float64) []int {
	lent := len(t)
	lens := len(s)
	ht := get_hash(t, lent, p, x)
	hs := get_hash(s, lent, p, x)
	xt := 1.0
	for i := 0; i < lent; i++ {
		xt *= x
		xt = float64(int64(xt) % int64(p))
	}
	var res []int
	for i := 0; i <= lens-lent; i++ {
		if ht == hs {
			if s[i:i+lent] == t {
				res = append(res, i)
			}
		}
		if i+lent < lens {
			hs = (hs*x - float64(int(s[i]))*xt + float64(int(s[i+lent])))
			hs = float64(int64(hs) % int64(p))
			hs = float64(int64(hs+p) % int64(p))
		}
	}
	return res
}

func Task1() {
	var s, t string
	fmt.Scanln(&s)
	fmt.Scanln(&t)
	p := 1e9 + 7
	x := 26
	res := rabin_karp(s, t, p, float64(x))
	for _, v := range res {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
