package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func merge2(left, right []int) ([]int, int) {
	k := 0
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
			k = k + len(left) - i
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result, k
}

func mergeSort2(arr []int, _ int) ([]int, int) {
	if len(arr) <= 1 {
		return arr, 0
	}
	mid := len(arr) / 2
	left, kl := mergeSort2(arr[:mid], 0)
	right, kr := mergeSort2(arr[mid:], 0)
	a, k := merge2(left, right)
	return a, k + kl + kr
}

func Task4() {
	var n int
	fmt.Scanln(&n)

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	arr := strings.Split(line, " ")

	intArr := make([]int, len(arr))
	for i, str := range arr {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		intArr[i] = num
	}

	_, count := mergeSort2(intArr, 0)
	fmt.Println(count)
}
