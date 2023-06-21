package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func merge(left, right []int, pos_left int, pos_right int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	// Описывает осуществление слияния, по одному на каждой строке. По условию задачи.
	fmt.Println(pos_left+1, pos_right+len(right), result[0], result[len(result)-1])
	return result
}

func mergeSort(arr []int, pos int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid], pos)
	right := mergeSort(arr[mid:], pos+mid)
	return merge(left, right, pos, pos+mid)
}

func Task3() {
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

	intArr := make([]int, n)
	for i := 0; i < n; i++ {
		intArr[i], _ = strconv.Atoi(arr[i])
	}

	sortedArr := mergeSort(intArr, 0)
	for _, v := range sortedArr {
		fmt.Print(v, " ")
	}
}
