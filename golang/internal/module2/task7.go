package module2

import (
	"bufio"
	"fmt"
	"os"
)

type elem struct {
	val  string
	next *elem
	prev *elem
}

type linkedList struct {
	head *elem
	tail *elem
}

func (l *linkedList) insert(element *elem) {
	if l.tail == nil {
		l.head = element
		l.tail = element
	} else {
		element.prev = l.tail
		l.tail.next = element
		l.tail = element
	}
}

func (l *linkedList) print() {
	buf := bufio.NewWriter(os.Stdout)
	defer buf.Flush()
	elem := l.head
	if elem == nil {
		fmt.Printf("empty\n")
	} else {
		for ; elem.next != nil; elem = elem.next {
			fmt.Fprintf(buf, "%s, ", elem.val)
		}
		fmt.Fprintf(buf, "%s\n", elem.val)
	}
}

func (l *linkedList) putElementsToSlice(numbers []string, lastIndexPut int) int {
	for elem := l.head; elem != nil; elem = elem.next {
		numbers[lastIndexPut] = elem.val
		lastIndexPut++
	}
	return lastIndexPut
}

func convertByteToInt(b byte) int {
	return int(b - '0')
}

const quantityOfBuckets = 10

func radixSort(numbers []string, digits int) {
	buckets := make([]linkedList, quantityOfBuckets)
	round := 1
	for d := digits; d > 0; d-- {
		for i := 0; i < quantityOfBuckets; i++ {
			buckets[i] = linkedList{}
		}
		for i := 0; i < len(numbers); i++ {
			bucketNumber := convertByteToInt(numbers[i][d-1])
			buckets[bucketNumber].insert(&elem{val: numbers[i], next: nil, prev: nil})
		}

		fmt.Printf("**********\nPhase %d\n", round)
		round++
		for i := 0; i < quantityOfBuckets; i++ {
			fmt.Printf("Bucket %d: ", i)
			buckets[i].print()
		}

		lastIndexPut := 0
		for i := 0; i < quantityOfBuckets; i++ {
			lastIndexPut = buckets[i].putElementsToSlice(numbers, lastIndexPut)
		}
	}
}

func printArray(numbers []string) {
	buf := bufio.NewWriter(os.Stdout)
	defer buf.Flush()
	i := 0
	for ; i < len(numbers)-1; i++ {
		fmt.Fprintf(buf, "%s, ", numbers[i])
	}
	fmt.Fprintf(buf, "%s\n", numbers[i])
}

func SolutionRadixSort() {
	buf := bufio.NewReader(os.Stdin)
	n := 0
	fmt.Fscanln(buf, &n)

	numbers := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(buf, &numbers[i])
	}
	fmt.Println("Initial array:")
	printArray(numbers)
	radixSort(numbers, len(numbers[0]))
	fmt.Printf("**********\nSorted array:\n")
	printArray(numbers)
}
