package module2

import (
	"fmt"
	"os"
)

func Stock() {
	var n, m int
	fmt.Fscan(os.Stdin, &n)
	kind := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(os.Stdin, &kind[i])
	}

	fmt.Fscan(os.Stdin, &m)
	for i := 0; i < m; i++ {
		var b int
		fmt.Fscan(os.Stdin, &b)
		kind[b-1]--
	}

	for i := 0; i < n; i++ {
		if kind[i] < 0 {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
