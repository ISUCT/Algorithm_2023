package module3

import (
	"bufio"
	"fmt"
	"os"
)

func z_function(s string) ([]int, int) {
	z := make([]int, len(s))
	leftBorder, rightBorder := 0, 0
	lengthOfString := len(s)

	indexOfMaxElementInZ := 0
	for i := 1; i < lengthOfString; i++ {
		if i <= rightBorder {
			z[i] = min(rightBorder-i+1, z[i-leftBorder])
		}

		for i+z[i] < lengthOfString && s[z[i]] == s[z[i]+i] {
			z[i]++
		}

		if z[i] > z[indexOfMaxElementInZ] {
			indexOfMaxElementInZ = i
		}

		if i+z[i]-1 > rightBorder {
			rightBorder = i + z[i] - 1
			leftBorder = i
		}
	}
	return z, indexOfMaxElementInZ
}

func Task4Solution() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = removeNewline(input)

	z, indexOfMaxElementInZ := z_function(input)
	lengthOfString := len(input)

	if indexOfMaxElementInZ+z[indexOfMaxElementInZ] == lengthOfString {
		fmt.Print(lengthOfString - z[indexOfMaxElementInZ])
	} else {
		max := 0
		for i := lengthOfString - 1; i >= 0; i-- {
			if z[i] > max && z[i] < z[indexOfMaxElementInZ] && i+z[i] == lengthOfString {
				max = z[i]
			}
		}

		fmt.Print(lengthOfString - max)
	}
}
