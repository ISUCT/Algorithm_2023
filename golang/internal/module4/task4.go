package module4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Код использует алгоритм "скользящего окна" с использованием очереди, где каждое новое число добавляется в очередь,
а старое удаляется при переходе окна на следующий шаг. Для нахождения минимума используется список индексов окна,
который сортируется по возрастанию значений в массиве a. Таким образом, первый элемент этого списка - всегда минимальный элемент в текущем окне.
*/

func Task4() {
	var n, k int
	fmt.Scanln(&n, &k)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	arr := strings.Split(line, " ")

	intArr := make([]int, n)
	for i := 0; i < n; i++ {
		intArr[i], _ = strconv.Atoi(arr[i])
	}

	window := make([]int, 0, k)

	for i := 0; i < k; i++ {
		for len(window) > 0 && intArr[i] <= intArr[window[len(window)-1]] {
			window = window[:len(window)-1]
		}
		window = append(window, i)
	}

	fmt.Println(intArr[window[0]])

	for i := k; i < n; i++ {
		if window[0] <= i-k {
			window = window[1:]
		}

		for len(window) > 0 && intArr[i] <= intArr[window[len(window)-1]] {
			window = window[:len(window)-1]
		}
		window = append(window, i)

		fmt.Println(intArr[window[0]])
	}
}
