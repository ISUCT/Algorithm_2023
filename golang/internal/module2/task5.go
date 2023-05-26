package module2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
)

func partition(arr []int, left, right int) int {
	i := left
	j := right
	piv := arr[(left+right)/2]
	swap := reflect.Swapper(arr)
	for i <= j {
		for arr[i] < piv {
			i++
		}
		for arr[j] > piv {
			j--
		}
		if i >= j {
			break
		}
		swap(i, j)
		i++
		j--
	}
	return j
}

func quickSort(s []int, left, right int) {
	if right > left {
		q := partition(s, left, right)
		quickSort(s, left, q)
		quickSort(s, q+1, right)
	}
}

func convertBytesToInt(buf []byte) (nums []int) {
	num := 0
	negativeFlag := 1
	for i := 0; i < len(buf); i++ {
		if buf[i] == '-' {
			negativeFlag = -1
		}
		if buf[i] >= '0' {
			num = num*10 + int(buf[i]-'0')
		} else if buf[i] == ' ' {
			nums = append(nums, num*negativeFlag)
			num = 0
			negativeFlag = 1
		}
	}
	nums = append(nums, num*negativeFlag)
	return nums
}

func readUserInput(input *os.File, quantity int) []int {
	nums := make([]int, 0, quantity)
	buf := bufio.NewReader(input)
	rawData, err := buf.ReadBytes('\n')
	nums = append(nums, convertBytesToInt(rawData)...)
	for err != nil && err != io.EOF {
		rawData, err = buf.ReadBytes('\n')
		nums = append(nums, convertBytesToInt(rawData)...)
	}
	return nums
}

func countDifferentElems(nums []int) int {
	counter := 1
	curNumber := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] != curNumber {
			curNumber = nums[i]
			counter++
		}
	}
	return counter
}

// Функция для запуска решения задачи
func SolutionCountDifferentElems() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	n := 0
	fmt.Fscanln(input, &n)
	slice := readUserInput(input, n)
	quickSort(slice, 0, len(slice)-1)

	output, _ := os.OpenFile("output.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
	defer output.Close()

	output.WriteString(fmt.Sprintf("%d", countDifferentElems(slice)))
}
