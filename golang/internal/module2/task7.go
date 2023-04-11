package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printArray(arr []string) {
	n := len(arr)
	if n == 0 {
		fmt.Println("empty")
	} else {
		for i := 0; i < n-1; i++ {
			fmt.Printf("%s, ", arr[i])
		}
		if n > 0 {
			fmt.Println(arr[n-1])
		}
	}
}

func radixSort(arr []string) []string {
	maxArr := 0
	for _, s := range arr {
		if len(s) > maxArr {
			maxArr = len(s)
		}
	}

	for i := maxArr - 1; i >= 0; i-- {
		buckets := make([][]string, 10)
		for _, s := range arr {
			if len(s) <= i {
				buckets[0] = append(buckets[0], s)
			} else {
				a, _ := strconv.Atoi(string(s[i]))
				buckets[a] = append(buckets[a], s)
			}
		}

		arr = []string{}
		for _, bucket := range buckets {
			arr = append(arr, bucket...)
		}

		fmt.Println("**********")
		fmt.Printf("Phase %d\n", maxArr-i)
		for j, bucket := range buckets {
			fmt.Printf("Bucket %d: ", j)
			printArray(bucket)
		}
	}

	return arr
}

func Task7() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i] = scanner.Text()
	}

	fmt.Println("Initial array:")
	printArray(arr)

	sortedArr := radixSort(arr)

	fmt.Println("**********")
	fmt.Println("Sorted array:")
	printArray(sortedArr)
}
