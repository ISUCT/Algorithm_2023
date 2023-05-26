package module4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SolutionTask2() {
	stack := InitStack[string](100)

	reader := bufio.NewReader(os.Stdin)

	trash, _ := reader.ReadString('\n')
	_ = trash

	input, _ := reader.ReadString('\n')
	input = removeNewline(input)
	trains := strings.Split(input, " ")

	neededTrain := 1
	for _, elem := range trains {
		if elem != strconv.Itoa(neededTrain) {
			stack.Push(elem)
		} else {
			neededTrain += 1
			for ; stack.Top() == strconv.Itoa(neededTrain); neededTrain++ {
				stack.Pop()
			}
		}
	}

	if stack.Len() > 0 {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}
}
