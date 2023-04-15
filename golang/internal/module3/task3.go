package module3

import (
	"fmt"
	"strings"
)

func Task3() {
	var s string
	fmt.Scan(&s)

	n := len(s)
	k := 1

	for i := 2; i <= n; i++ {
		if n%i == 0 {
			if strings.Repeat(s[:n/i], i) == s {
				k = i
			}
		}
	}
	fmt.Println(k)
}
