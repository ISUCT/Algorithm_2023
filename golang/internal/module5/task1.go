package module5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Task1() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	arr := strings.Split(line, " ")

	n := len(arr) - 1
	value := make([]int, n)
	for i := 0; i < n; i++ {
		value[i], _ = strconv.Atoi(arr[i])
	}

	for i := 0; i < n-2; i++ {
		fmt.Println(value[i], value[i+1], value[i+2])
		if value[i] > value[i+1] && value[i] < value[i+1] {
			fmt.Println(value[i])
		}
	}
}
