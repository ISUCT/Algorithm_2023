package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Merge(a []int, b []int) []int {
	var i int = 0
	var j int = 0
	merged := make([]int, len(a)+len(b))
	for k := 0; k < len(merged); k++ {
		if j >= len(b) || (i < len(a) && a[i] <= b[j]) {
			merged[k] = a[i]
			i++
		} else {
			merged[k] = b[j]
			j++
		}
	}
	return merged
}

func MergeSort(arr []int, l int, r int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// Находим середину массива
	m := len(arr) / 2
	// Сортируем левую и правую половины независимо
	left := MergeSort(arr[0:m], l, l+m)
	right := MergeSort(arr[m:], l+m, r)
	// Сливаем отсортированные половины
	fmt.Printf("%d %d ", l+1, r)
	result := Merge(left, right)
	if len(result) > 0 {
		fmt.Printf("%d %d\n", result[0], result[len(result)-1])
	}
	return result
}

func MergeTask() {
	var n int
	fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	line = strings.TrimSuffix(line, " ")
	str_arr := strings.Split(line, " ")
	arr := make([]int, n)
	for idx, val := range str_arr {
		arr[idx], err = strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
	}
	result := MergeSort(arr, 0, len(arr))
	for idx, val := range result {
		if idx < len(result)-1 {
			fmt.Printf("%d ", val)
		} else {
			fmt.Printf("%d\n", val)
		}
	}
}
