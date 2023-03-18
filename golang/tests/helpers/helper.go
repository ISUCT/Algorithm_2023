package helpers

import (
	"os"
	"testing"
)

func Replacer(input string, t *testing.T) (*os.File, *os.File) {
	inp := []byte(input)

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	_, err = w.Write(inp)
	if err != nil {
		t.Fatal(err)
	}
	return r, w
}

func GetInputAndOutputFiles() (*os.File, *os.File) {
	input, _ := os.OpenFile("input.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
	output, _ := os.OpenFile("output.txt", os.O_RDONLY, 0755)
	return input, output
}
