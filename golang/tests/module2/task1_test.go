package module2_test

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module2"
)

func TestSolutionBubbleSort(t *testing.T) {
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

		w.Write([]byte("4\n4 3 2 1\n"))
		os.Stdin, os.Stdout = r, w
		module2.SolutionBubbleSort()
		w.Close()

		out, err := io.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		expectedOutput, _ := os.ReadFile("testData/task1/task1_output_normalcase.txt")
		assert.Equal(expectedOutput, out)
	})
	t.Run("Case: already sorted array", func(t *testing.T) {
		r, w, err := os.Pipe()
		if err != nil {
			log.Fatal(err)
		}

		w.Write([]byte("4\n1 2 3 4\n"))
		os.Stdin, os.Stdout = r, w
		module2.SolutionBubbleSort()
		w.Close()

		out, err := io.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		expectedOutput := 0
		/**
		Данные в буфере хранятся в виде ascii кодов, поэтому пришлось
		конвертнуть output в int
		**/
		assert.Equal(expectedOutput, int(out[0]-'0'))
	})
	t.Run("Case: one number", func(t *testing.T) {
		r, w, err := os.Pipe()
		if err != nil {
			log.Fatal(err)
		}

		w.Write([]byte("1\n37723\n"))
		os.Stdin, os.Stdout = r, w
		module2.SolutionBubbleSort()
		w.Close()

		out, err := io.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		expectedOutput := 0
		assert.Equal(expectedOutput, int(out[0]-'0'))
	})
}
