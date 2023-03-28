package module2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Task2() {
	var n int
	fmt.Scanln(&n)
	var arr [][]string

	for i := 0; i < n; i++ {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimSuffix(line, "\r")
		input := strings.Split(line, " ")

		// Numeric Conversions
		// intArr := make([]int, len(input))
		// for i, str := range input {
		// 	num, err := strconv.Atoi(str)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	intArr[i] = num
		// }
		// arr = append(arr, intArr)
		arr = append(arr, input)

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
