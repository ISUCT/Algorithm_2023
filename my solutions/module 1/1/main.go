// go mod init main
// go mod tidy
package main

import "fmt"

func main() {
	// fmt.Print("Hello world")

	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	// fmt.Printf("%d %d", a, b)
	fmt.Printf("%d", a+b)
}
