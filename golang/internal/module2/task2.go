package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Task2() {
	scanner := bufio.NewScanner(os.Stdin)
	var n int
	fmt.Scanln(&n)
	arr := make([][]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		input2 := make([]int, 2)
		input2[0], _ = strconv.Atoi(input[0])
		input2[1], _ = strconv.Atoi(input[1])
		arr[i] = input2

		j := i - 1
		x := arr[i]
		for j >= 0 && (arr[j][1] < x[1] || arr[j][1] == x[1] && arr[j][0] > x[0]) {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = x
	}
	for _, v := range arr {
		fmt.Println(v[0], v[1])
	}
}
