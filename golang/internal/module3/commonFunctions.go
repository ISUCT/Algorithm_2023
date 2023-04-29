package module3

import "strings"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func removeNewline(str string) string {
	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimSuffix(str, "\r")
	return str
}
