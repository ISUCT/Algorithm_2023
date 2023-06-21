package module4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Task1() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	arr := strings.Split(line, "")

	open := 0
	closed := 0

	for _, v := range arr {
		if v == "(" {
			open++
		} else {
			if open == 0 {
				closed++
			} else {
				open--
			}
		}
	}
	fmt.Println(open + closed)
}
