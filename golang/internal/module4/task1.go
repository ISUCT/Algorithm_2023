package module4

import (
	"bufio"
	"fmt"
	"os"
)

func SolutionTask1() {
	stack := InitStack[rune](100000)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = removeNewline(input)

	for _, elem := range input {
		if elem == '(' {
			stack.Push(elem)
		} else {
			if stack.Top() == '(' {
				stack.Pop()
			} else {
				stack.Push(elem)
			}
		}
	}

	fmt.Print(stack.Len())
}
