package module5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func nod(a, b int) int {
	if b == 0 {
		return a
	}
	return nod(b, a%b)
}

func Task3() {
	var n, k int
	fmt.Scanln(&n)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	inputArr := strings.Split(line, " ")

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(inputArr[i])
	}

	fmt.Scanln(&k)
	var l, r int
	result := make([]int, k)
	for i := 0; i < k; i++ {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimSuffix(line, "\r")
		inputlr := strings.Split(line, " ")

		l, _ = strconv.Atoi(inputlr[0])
		r, _ = strconv.Atoi(inputlr[1])

		result[i] = arr[l-1]
		for j := l; j < r; j++ {
			result[i] = nod(result[i], arr[j])
		}
	}
	for i := 0; i < k; i++ {
		fmt.Print(result[i], " ")
	}
}
