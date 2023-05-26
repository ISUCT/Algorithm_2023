package module2_test

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module2"
)

func TestSolutionInsertionSort(t *testing.T) {
	assert := assert.New(t)

	oldStdin := *os.Stdin
	oldStdout := *os.Stdout
	defer func(v *os.File) { os.Stdin = v }(&oldStdin)
	defer func(v *os.File) { os.Stdout = v }(&oldStdout)

	t.Run("Case: usual input", func(t *testing.T) {
		r, w, err := os.Pipe()
		if err != nil {
			log.Fatal(err)
		}

		input, _ := os.ReadFile("testData/task2/task2_input_normalcase.txt")
		w.Write(input)

		os.Stdin, os.Stdout = r, w
		module2.SolutionInsertionSort()
		w.Close()

		out, err := io.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		expectedOutput, _ := os.ReadFile("testData/task2/task2_output_normalcase.txt")
		assert.Equal(expectedOutput, out)
	})
	t.Run("Case: one item", func(t *testing.T) {
		r, w, err := os.Pipe()
		if err != nil {
			log.Fatal(err)
		}

		w.Write([]byte("1\n5 100\n"))

		os.Stdin, os.Stdout = r, w
		module2.SolutionInsertionSort()
		w.Close()

		out, err := io.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		expectedOutput := []byte("5 100\n")
		assert.Equal(expectedOutput, out)
	})
	t.Run("Case: 1000 items", func(t *testing.T) {
		/*
			Конечно же pipe "ломается" при большом импуте, поэтому тут будет удобнее
			использовать текстовый файл, а не использовать горутины и прочие штуки
		*/
		input, _ := os.Open("testData/task2/task2_input_1000elements.txt")
		output, _ := os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		os.Stdin, os.Stdout = input, output
		module2.SolutionInsertionSort()

		out, err := os.ReadFile("output.txt")
		if err != nil {
			log.Fatal(err)
		}

		expectedOutput, _ := os.ReadFile("testData/task2/task2_output_1000elements.txt")
		assert.Equal(expectedOutput, out)
	})
}
