package module2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

const idIndex = 0
const priceIndex = 1

func swapColumns(listOfItems [][]int, indexSortedElem, indexKeyElem int) {
	swapper := reflect.Swapper(listOfItems)
	if (listOfItems[indexSortedElem][priceIndex] == listOfItems[indexKeyElem][priceIndex]) &&
		(listOfItems[indexKeyElem][idIndex] < listOfItems[indexSortedElem][idIndex]) {
		swapper(indexKeyElem, indexSortedElem)
	} else if listOfItems[indexKeyElem][priceIndex] > listOfItems[indexSortedElem][priceIndex] {
		swapper(indexKeyElem, indexSortedElem)
	}
}

func insertionSort(listOfItems [][]int, length int) {
	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if listOfItems[j][priceIndex] < listOfItems[j-1][priceIndex] {
				break
			} else {
				swapColumns(listOfItems, j-1, j)
			}
		}
	}
}

func printListOfItems(listOfItems [][]int, length int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < length; i++ {
		_, err := fmt.Fprintf(writer, "%d %d\n", listOfItems[i][idIndex], listOfItems[i][priceIndex])
		if err != nil {
			log.Fatal(err)
		}
		writer.Flush()
	}
}

func SolutionInsertionSort() {
	reader := bufio.NewReader(os.Stdin)
	n := 0
	fmt.Fscanln(reader, &n)

	listOfItems := make([][]int, n)
	for i := 0; i < n; i++ {
		listOfItems[i] = make([]int, 2)
		fmt.Fscanf(reader, "%d %d\n", &listOfItems[i][idIndex], &listOfItems[i][priceIndex])
	}

	insertionSort(listOfItems, n)
	printListOfItems(listOfItems, n)
}
