package module2

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func merge(numbers []int, startIndex, middleIndex, endIndex int) {
	quantityOfElemsOnLeft := middleIndex - startIndex
	quantityOfElemsOnRight := endIndex - middleIndex + 1

	left := make([]int, quantityOfElemsOnLeft+1)
	right := make([]int, quantityOfElemsOnRight+1)

	left[quantityOfElemsOnLeft] = math.MaxInt
	right[quantityOfElemsOnRight] = math.MaxInt

	for i := 0; i < quantityOfElemsOnLeft; i++ {
		left[i] = numbers[startIndex+i]
	}

	for j := 0; j < quantityOfElemsOnRight; j++ {
		right[j] = numbers[middleIndex+j]
	}

	i, j := 0, 0
	k := startIndex
	for ; k <= endIndex; k++ {
		if left[i] <= right[j] {
			numbers[k] = left[i]
			i++
		} else {
			numbers[k] = right[j]
			j++
		}
	}

	fmt.Printf("%d %d %d %d\n", startIndex+1, endIndex+1, numbers[startIndex], numbers[endIndex])
}

func mergeSort(numbers []int, startIndex, endIndex int) {
	if endIndex > startIndex {
		/*
			Пришлось включить серединный элемент в правую часть, а не в левую,
			чтобы ответы на задачу совпадали с яндекс контестом
		*/
		middleIndex := (startIndex+endIndex)/2 + (startIndex+endIndex)%2
		mergeSort(numbers, startIndex, middleIndex-1)
		mergeSort(numbers, middleIndex, endIndex)
		merge(numbers, startIndex, middleIndex, endIndex)
	}
}

func printSortedSliceMergeSort(numbers []int, length int) {
	writer := bufio.NewWriter(os.Stdout)
	i := 0
	for ; i < length-1; i++ {
		fmt.Fprintf(writer, "%d ", numbers[i])
	}
	fmt.Fprintf(writer, "%d", numbers[i])
	writer.Flush()
}

func SolutionMergeSort() {
	reader := bufio.NewReader(os.Stdin)
	n := 0
	fmt.Fscanln(reader, &n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &numbers[i])
	}

	mergeSort(numbers, 0, n-1)
	printSortedSliceMergeSort(numbers, n)
}
