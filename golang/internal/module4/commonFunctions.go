package module4

import "strings"

func removeNewline(str string) string {
	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimSuffix(str, "\r")
	return str
}
