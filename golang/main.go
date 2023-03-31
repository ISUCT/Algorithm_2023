package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/io_sample"
)

func main() {
	fmt.Println("Hello world")
	err := io_sample.WriteFile(8096, "test.txt")
	if err != nil {
		fmt.Printf("error writing file :%v \n", err)
	}
	fmt.Printf("File wrote sucessfully")
	err = io_sample.ReadFile("test.txt")
	if err != nil {
		fmt.Println("error while reading file")
	}

}
