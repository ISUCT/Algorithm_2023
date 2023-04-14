package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Merge(a []Item, b []Item) []Item {
	merged := make([]Item, len(a)+len(b))
	i := 0
	j := 0
	for k := 0; k < len(merged); k++ {
		if j >= len(b) || (i < len(a) && a[i].Value < b[j].Value) {
			merged[k] = a[i]
			i += 1
		} else {
			merged[k] = b[j]
			j += 1
		}
	}
	return merged
}

func Sort(source []Item) []Item {
	if len(source) <= 1 {
		return source
	}
	L := source[0 : len(source)/2]
	R := source[len(source)/2:]
	L = Sort(L)
	R = Sort(R)
	res := Merge(L, R)
	fmt.Printf("%d %d %d %d\n", res[0].Index, res[len(res)-1].Index, res[0].Value, res[len(res)-1].Value)
	return res
}

type Item struct {
	Index int
	Value int
}

func MergeSort() {
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
	arr := make([]Item, n)
	for idx, val := range str_arr {
		value, err := strconv.Atoi(val)
		arr[idx] = Item{
			Value: value,
			Index: idx + 1,
		}
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(arr)
}
