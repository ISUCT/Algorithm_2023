package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InversionCount() {
	var n int
	fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	str_arr := strings.Split(line, " ")
	arr := make([]int, n)
	for idx, val := range str_arr {
		arr[idx], err = strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
	}
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				count++
			}
		}
	}
	fmt.Println(count)

}
