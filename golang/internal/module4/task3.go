package module4

import (
	"bufio"
	"fmt"
	"os"
)

func printSlice[E Element](slice []E) {
	writer := bufio.NewWriter(os.Stdout)

	for i := 0; i < len(slice); i++ {
		fmt.Fprintf(writer, "%v ", slice[i])
	}

	writer.Flush()
}

func Task3Solution() {
	reader := bufio.NewReader(os.Stdin)

	quantityOfElements := 0
	fmt.Fscanf(reader, "%d\n", &quantityOfElements)

	numbers := make([]int, quantityOfElements)
	for i := 0; i < quantityOfElements; i++ {
		fmt.Fscan(reader, &numbers[i])
	}

	increasingMonotonicStack := InitStack[int](100000)
	nextSmallerElement := make([]int, quantityOfElements)

	for j := quantityOfElements - 1; j >= 0; j-- {
		for !increasingMonotonicStack.IsEmpty() && numbers[j] <= numbers[increasingMonotonicStack.Top()] {
			increasingMonotonicStack.Pop()
		}

		if increasingMonotonicStack.IsEmpty() {
			nextSmallerElement[j] = -1
		} else {
			nextSmallerElement[j] = increasingMonotonicStack.Top()
		}

		increasingMonotonicStack.Push(j)
	}

	printSlice(nextSmallerElement)
}
