package module3

import (
	"fmt"
	"strings"
)

func Task4() {
	var s string
	fmt.Scan(&s)

	// Ищем все вхождения первого символа в строку
	firstChar := rune(s[0])
	indexes := make([]int, 0)
	for i, c := range s {
		if c == firstChar {
			indexes = append(indexes, i)
		}
	}

	// Ищем периодические подстроки, начинающиеся с каждого вхождения
	minLength := len(s)
	for _, i := range indexes {
		sub := s[i:]
		index := strings.Index(sub[1:], sub[:1]) + 1
		if index == 0 || index == len(sub)-1 {
			continue
		}
		period := 1
		for j := 1; j < len(sub); j++ {
			if sub[j%index] != sub[j] {
				period = j + 1
				break
			}
		}
		repetitions := len(sub) / period
		if repetitions*period < len(s) && repetitions*period < minLength {
			minLength = repetitions * period
		}
	}

	// Если не нашли периодических подстрок, то строка состоит из повторений одной подстроки
	if minLength == len(s) {
		fmt.Println(len(s))
		return
	}

	fmt.Println(minLength)
}
