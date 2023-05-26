package module3

import (
	"bufio"
	"fmt"
	"os"
)

const base int64 = 137
const divisor int64 = 1e9 + 7

func getHash(inputString string, windowSize int) int64 {
	var hash int64 = 0
	for i := 0; i < windowSize; i++ {
		// Обратный полиноминальный хэш; формула получена благодаря схеме Горнера (преобразовали формулу прямого полиноминального хэша)
		hash = (hash*base + int64(inputString[i])) % divisor
	}
	return hash
}

func RapinKarpSlidingHash(wholeString, substring string, windowSize int) {
	writer := bufio.NewWriter(os.Stdout)

	hashOfTheWindowOfWholeString := getHash(wholeString, windowSize)
	hashOfTheSubstring := getHash(substring, windowSize)

	var pow int64 = 1
	for i := 0; i < windowSize; i++ {
		pow = (pow * base) % divisor
	}

	for i := 0; i <= len(wholeString)-len(substring); i++ {
		if hashOfTheWindowOfWholeString == hashOfTheSubstring {
			fmt.Fprintf(writer, "%d ", i)
		}

		// Так как тут используется обратный полиноминальный хэш (как и в функции хэша выше),
		// первый элемент "окна" имеет вид s[0]*p^n и именно его и убираем
		if i+windowSize < len(wholeString) {
			hashOfTheWindowOfWholeString = (hashOfTheWindowOfWholeString*base - int64(wholeString[i])*pow + int64(wholeString[i+windowSize])) % divisor
			hashOfTheWindowOfWholeString = (hashOfTheWindowOfWholeString + divisor) % divisor
		}
	}

	writer.Flush()
}
