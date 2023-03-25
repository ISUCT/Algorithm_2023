package module2

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func mergeInvCount(numbers []int, startIndex, middleIndex, endIndex int) (invCount int64) {
	quantityOfElemsOnLeft := middleIndex - startIndex + 1
	quantityOfElemsOnRight := endIndex - middleIndex

	left := make([]int, quantityOfElemsOnLeft+1)
	right := make([]int, quantityOfElemsOnRight+1)

	left[quantityOfElemsOnLeft] = math.MaxInt
	right[quantityOfElemsOnRight] = math.MaxInt

	for i := 0; i < quantityOfElemsOnLeft; i++ {
		left[i] = numbers[startIndex+i]
	}

	for j := 0; j < quantityOfElemsOnRight; j++ {
		right[j] = numbers[middleIndex+1+j]
	}

	i, j := 0, 0
	k := startIndex
	for ; k <= endIndex; k++ {
		if left[i] <= right[j] {
			numbers[k] = left[i]
			i++
		} else {
			numbers[k] = right[j]
			invCount += int64(middleIndex) - int64(startIndex) - int64(i) + 1
			j++
		}
	}

	return invCount
}

func mergeSortInvCount(numbers []int, startIndex, endIndex int) (invCount int64) {
	if endIndex > startIndex {
		middleIndex := (startIndex + endIndex) / 2
		invCount = mergeSortInvCount(numbers, startIndex, middleIndex)
		invCount += mergeSortInvCount(numbers, middleIndex+1, endIndex)
		invCount += mergeInvCount(numbers, startIndex, middleIndex, endIndex)
	}
	return invCount
}

func SolutionInversionCount() {
	reader := bufio.NewReader(os.Stdin)
	n := 0
	fmt.Fscanln(reader, &n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &numbers[i])
	}

	invCount := mergeSortInvCount(numbers, 0, n-1)
	fmt.Print(invCount)
}
