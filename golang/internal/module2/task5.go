package module2

import (
	"fmt"
)

func Task5() {
	var n int
	fmt.Scan(&n)

	numbers := make(map[int]bool)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		numbers[num] = true
	}

	fmt.Println(len(numbers))
}
