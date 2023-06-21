package module2

import (
	"fmt"
)

func Task6() {
	var n, k int
	fmt.Scan(&n)
	products := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&products[i])
	}

	fmt.Scan(&k)
	orders := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&orders[i])
	}

	counts := make([]int, n)
	for _, v := range orders {
		counts[v-1]++
	}
	for i := 0; i < n; i++ {
		if products[i] < counts[i] {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
