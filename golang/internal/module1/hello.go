package module1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SayHello() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	fmt.Printf("Hello, %s!", line)
}
