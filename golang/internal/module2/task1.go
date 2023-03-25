package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Task1() {
	var n int
	fmt.Scanln(&n)

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	arr := strings.Split(line, " ")

	p := false
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			a, _ := strconv.Atoi(arr[j])
			b, _ := strconv.Atoi(arr[j+1])
			if a > b {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				fmt.Println(strings.Join(arr, " "))
				p = true
			}
		}
	}
	if !p {
		fmt.Println(0)
	}
}
