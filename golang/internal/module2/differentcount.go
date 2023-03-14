package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DifferentCount() {
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
	arr1 := MergeSort(arr, 0, n)
	counter := 1
	for i := 0; i < n-1; i++ {
		if arr1[i] != arr1[i+1] {
			counter++
		}
	}
	fmt.Println(counter)
}
