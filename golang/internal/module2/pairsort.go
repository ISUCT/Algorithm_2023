package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PairSort() {
	var n int
	fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	massiv := make([][]int, 0)
	for i := 0; i < n; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimSuffix(line, "\r")
		string_arr := strings.Split(line, " ")
		int_arr := make([]int, 0)
		for _, val := range string_arr {
			intVal, _ := strconv.Atoi(val)
			int_arr = append(int_arr, intVal)
		}
		massiv = append(massiv, int_arr)

	}
	for i := 0; i <= n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if massiv[j+1][1] > massiv[j][1] {
				point := massiv[j+1]
				massiv[j+1] = massiv[j]
				massiv[j] = point
			} else if massiv[j][1] == massiv[j+1][1] {
				if massiv[j+1][0] < massiv[j][0] {
					point := massiv[j+1]
					massiv[j+1] = massiv[j]
					massiv[j] = point
				}
			}
		}
	}
	for _, a := range massiv {
		for _, b := range a {
			string_b := strconv.Itoa(b)
			fmt.Printf(string_b)
			fmt.Printf(" ")
		}
		fmt.Println()
	}
}
