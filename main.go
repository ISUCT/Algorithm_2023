package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func BubbleS() {
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
	nswaps := 0
	for i := 0; i <= n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
				nswaps += 1
				fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
			}
		}
	}
	if nswaps == 0 {
		fmt.Println("0")
	}
}

func Sorted() {
	var n int
	fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	mass := make([][]int, 0)
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
		mass = append(mass, int_arr)

	}
	for i := 0; i <= n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if mass[j+1][1] > mass[j][1] {
				point := mass[j+1]
				mass[j+1] = mass[j]
				mass[j] = point
			} else if mass[j][1] == mass[j+1][1] {
				if mass[j+1][0] < mass[j][0] {
					point := mass[j+1]
					mass[j+1] = mass[j]
					mass[j] = point
				}
			}
		}
	}
	for _, s := range mass {
		for _, m := range s {
			string_m := strconv.Itoa(m)
			fmt.Printf(string_m)
			fmt.Printf(" ")
		}
		fmt.Println()
	}
}

func NumberOfInversions() {
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
	c := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				c++
			}
		}
	}
	fmt.Println(c)

}

func Varietes() {
	var n int
	fmt.Fscan(os.Stdin, &n)
	mass := make([]int, n)
	mi1 := int(1e9)
	ma1 := int(-1e9)

	for i := 0; i < n; i++ {
		fmt.Fscan(os.Stdin, &mass[i])
		mi1 = int(math.Min(float64(mi1), float64(mass[i])))
		ma1 = int(math.Max(float64(ma1), float64(mass[i])))
	}

	ress := make([]int, ma1-mi1+1)
	count := 0

	for i := 0; i < n; i++ {
		ress[mass[i]-mi1]++
		if ress[mass[i]-mi1] == 1 {
			count++
		}
	}

	fmt.Print(count)
}

func Storage() {
	var n, b int
	fmt.Fscan(os.Stdin, &n)
	goods := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(os.Stdin, &goods[i])
	}

	fmt.Fscan(os.Stdin, &b)
	for i := 0; i < b; i++ {
		var ab int
		fmt.Fscan(os.Stdin, &ab)
		goods[ab-1]--
	}

	for i := 0; i < n; i++ {
		if goods[i] < 0 {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
