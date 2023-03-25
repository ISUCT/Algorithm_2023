package module2

import (
	"bufio"
	"fmt"
	"os"
)

func storageProblem(productCount, listOfOrders []int) {
	writer := bufio.NewWriter(os.Stdout)

	for i := 0; i < len(listOfOrders); i++ {
		productCount[listOfOrders[i]-1]--
	}
	for j := 0; j < len(productCount); j++ {
		if productCount[j] > -1 {
			fmt.Fprintln(writer, "no")
		} else {
			fmt.Fprintln(writer, "yes")
		}
	}
	writer.Flush()
}

func SolutionStorageProblem() {
	reader := bufio.NewReader(os.Stdin)

	productCountLength := 0
	fmt.Fscanln(reader, &productCountLength)
	productCount := make([]int, productCountLength)
	for i := 0; i < productCountLength-1; i++ {
		fmt.Fscanf(reader, "%d ", &productCount[i])
	}
	// Новые строки должны совпадать с новыми строками в инпуте, здорово
	fmt.Fscanf(reader, "%d\n", &productCount[productCountLength-1])

	listOfOrdersLength := 0
	fmt.Fscanln(reader, &listOfOrdersLength)
	listOfOrders := make([]int, listOfOrdersLength)
	for i := 0; i < listOfOrdersLength-1; i++ {
		fmt.Fscanf(reader, "%d ", &listOfOrders[i])
	}
	fmt.Fscanf(reader, "%d\n", &listOfOrders[listOfOrdersLength-1])

	storageProblem(productCount, listOfOrders)
}
