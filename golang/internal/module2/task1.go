package module2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
)

func printSlice(slice []int, length int, writer io.Writer) {
	i := 0
	for ; i < length-1; i++ {
		fmt.Fprintf(writer, "%d ", slice[i])
	}
	fmt.Fprintf(writer, "%d\n", slice[i])
}

func bubbleSort(numbers []int, length int) (hasChanged bool) {
	swapper := reflect.Swapper(numbers)
	writer := bufio.NewWriter(os.Stdout)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				swapper(j, j+1)
				hasChanged = true
				printSlice(numbers, length, writer)
			}
		}
		writer.Flush()
	}
	return hasChanged
}

func SolutionBubbleSort() {
	n := 0
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &numbers[i])
	}

	hasChanged := bubbleSort(numbers, n)
	if !hasChanged {
		fmt.Print(0)
	}
}
