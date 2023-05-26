package module4

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var dequeOverflowError = errors.New("Deque overflow")
var dequeUnderflowError = errors.New("Deque underflow")
var triedToAccessEmptyDeque = errors.New("Deque is empty")

type Deque struct {
	head     int
	tail     int
	size     int
	capacity int
	elements []int
}

func InitDeque(numberOfElems int) Deque {
	var emptyDeque Deque
	emptyDeque.elements = make([]int, numberOfElems)
	emptyDeque.capacity = numberOfElems
	return emptyDeque
}

func (d *Deque) IsEmpty() bool {
	return d.Len() <= 0
}

func (d *Deque) isFull() bool {
	return d.size == d.capacity
}

func (d *Deque) Len() int {
	return d.size
}

func (d *Deque) getNextPosition(position int) int {
	return (position + 1) % d.capacity
}

func (d *Deque) getPreviousPosition(position int) int {
	prev := position - 1
	if prev < 0 {
		return d.capacity - 1
	}
	return prev
}

func (d *Deque) PushHead(element int) error {
	if d.isFull() {
		return dequeOverflowError
	}

	d.head = d.getPreviousPosition(d.head)
	d.elements[d.head] = element
	d.size++

	return nil
}

func (d *Deque) PushRear(element int) error {
	if d.isFull() {
		return dequeOverflowError
	}

	d.elements[d.tail] = element
	d.tail = d.getNextPosition(d.tail)
	d.size++

	return nil
}

func (d *Deque) PopHead() (int, error) {
	if d.IsEmpty() {
		return 0, dequeUnderflowError
	}

	element := d.elements[d.head]
	d.head = d.getNextPosition(d.head)
	d.size--

	return element, nil
}

func (d *Deque) PopRear() (int, error) {
	if d.IsEmpty() {
		return 0, dequeUnderflowError
	}

	d.tail = d.getPreviousPosition(d.tail)
	element := d.elements[d.tail]
	d.size--

	return element, nil
}

func (d *Deque) PeekHead() (int, error) {
	if d.IsEmpty() {
		return 0, triedToAccessEmptyDeque
	}

	return d.elements[d.head], nil
}

func (d *Deque) PeekRear() (int, error) {
	if d.IsEmpty() {
		return 0, triedToAccessEmptyDeque
	}

	return d.elements[d.getPreviousPosition(d.tail)], nil
}

func printResult(numbers []int) {
	writer := bufio.NewWriter(os.Stdout)
	for _, elem := range numbers {
		fmt.Fprintln(writer, elem)
	}
	writer.Flush()
}

func SolutionTask4() {
	reader := bufio.NewReader(os.Stdin)
	var lengthOfArray, windowSize int
	fmt.Fscanln(reader, &lengthOfArray, &windowSize)

	numbers := make([]int, lengthOfArray)
	for i := 0; i < lengthOfArray; i++ {
		fmt.Fscan(reader, &numbers[i])
	}

	deque := InitDeque(lengthOfArray)
	mins := make([]int, 0, lengthOfArray-windowSize+1)
	for i := 0; i < lengthOfArray; i++ {
		if val, err := deque.PeekHead(); err == nil && val <= i-windowSize {
			deque.PopHead()
		}

		for index, err := deque.PeekRear(); err == nil && numbers[i] <= numbers[index]; index, err = deque.PeekRear() {
			deque.PopRear()
		}

		deque.PushRear(i)

		if index, _ := deque.PeekHead(); i >= windowSize-1 {
			mins = append(mins, numbers[index])
		}
	}

	printResult(mins)
}
