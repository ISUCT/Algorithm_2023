package module3

import (
	"bufio"
	"fmt"
	"os"
)

func RapinKarpSlidingHashTask2(wholeString, substring string, windowSize int) int {
	hashOfTheWindowOfWholeString := getHash(wholeString, windowSize)
	hashOfTheSubstring := getHash(substring, windowSize)

	var pow int64 = 1
	for i := 0; i < windowSize; i++ {
		pow = (pow * base) % divisor
	}

	for i := 0; i <= len(wholeString)-len(substring); i++ {
		if hashOfTheWindowOfWholeString == hashOfTheSubstring {
			return i
		}

		// Так как тут используется обратный полиноминальный хэш (как и в функции хэша выше),
		// первый элемент "окна" имеет вид s[0]*p^n и именно его и убираем
		if i+windowSize < len(wholeString) {
			hashOfTheWindowOfWholeString = (hashOfTheWindowOfWholeString*base - int64(wholeString[i])*pow + int64(wholeString[i+windowSize])) % divisor
			hashOfTheWindowOfWholeString = (hashOfTheWindowOfWholeString + divisor) % divisor
		}
	}

	return -1
}

func Task2Solution() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	substr, _ := reader.ReadString('\n')

	str = removeNewline(str)
	substr = removeNewline(substr)

	substr = substr + substr

	index := RapinKarpSlidingHashTask2(substr, str, len(str))
	fmt.Print(index)
}
