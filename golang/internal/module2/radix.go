package module2

import (
	"fmt"
)

func Radix() {
	var n int
	fmt.Scanln(&n)
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr[i])
	}
	fmt.Println(arr)
}
