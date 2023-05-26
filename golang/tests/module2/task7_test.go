package module2_test

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module2"
)

func TestSolutionRadixSort(t *testing.T) {
	assert := assert.New(t)

	oldStdin := *os.Stdin
	oldStdout := *os.Stdout
	defer func(v *os.File) { os.Stdin = v }(&oldStdin)
	defer func(v *os.File) { os.Stdout = v }(&oldStdout)
	t.Run("Case: usual input", func(t *testing.T) {
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatal(err)
		}
		input, err := os.ReadFile("testData/task7/task7_input_usual.txt")
		if err != nil {
			log.Fatal(err)
		}
		expectedOutput, err := os.ReadFile("testData/task7/task7_output_usual.txt")

		_, err = io.Copy(w, bytes.NewReader(input))

		os.Stdin, os.Stdout = r, w

		module2.SolutionRadixSort()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(expectedOutput, out)
	})
	t.Run("Case: one number", func(t *testing.T) {
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatal(err)
		}

		input := []byte("1\n0132\n")

		expectedOutput, err := os.ReadFile("testData/task7/task7_output_onenumber.txt")

		_, err = io.Copy(w, bytes.NewReader(input))

		os.Stdin, os.Stdout = r, w

		module2.SolutionRadixSort()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(expectedOutput, out)
	})
}
