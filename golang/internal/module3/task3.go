package module3

import (
	"bufio"
	"fmt"
	"os"
)

func Z_function(s string) []int {
	lengthOfString := len(s)
	z := make([]int, lengthOfString)

	for i := 1; i < lengthOfString; i++ {
		for i+z[i] < lengthOfString && s[z[i]] == s[z[i]+i] {
			z[i]++
		}
	}
	return z
}

func Z_function_enchanced(s string) []int {
	z := make([]int, len(s))
	leftBorder, rightBorder := 0, 0
	lengthOfString := len(s)

	for i := 1; i < lengthOfString; i++ {
		if i <= rightBorder {
			z[i] = min(rightBorder-i+1, z[i-leftBorder])
		}

		for i+z[i] < lengthOfString && s[z[i]] == s[z[i]+i] {
			z[i]++
		}

		if i+z[i]-1 > rightBorder {
			rightBorder = i + z[i] - 1
			leftBorder = i
		}
	}
	return z
}

func Task3Solution() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = removeNewline(input)

	arr := Z_function_enchanced(input)
	lengthOfString := len(input)

	// Находим длину минимального суффикса, равного префиксу, проверяем, действительно ли он составляет строку из таких же одинаковых подстрок,
	// и выводим результат решения задачи
	for index, elem := range arr {
		if (index+elem == lengthOfString) && (lengthOfString%index == 0) {
			fmt.Print(lengthOfString / index)
			break
		}
	}
}
