package module1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Summ() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	arr := strings.Split(line, " ")

	a, _ := strconv.Atoi(arr[0])
	b, _ := strconv.Atoi(arr[1])
	summ := a + b
	fmt.Println(summ)
}

func BubbleSort() {
	var n int
	fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	str_arr := strings.Split(line, " ")
	arr := make([]int, n)
	for idx, val := range str_arr {
		arr[idx], err = strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
	}
	num_swaps := 0
	for i := 0; i <= n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
				num_swaps += 1
				fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
			}
		}
	}
	if num_swaps == 0 {
		fmt.Println("0")
	}
}
