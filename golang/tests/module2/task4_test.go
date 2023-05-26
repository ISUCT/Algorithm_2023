package module2_test

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module2"
)

func TestSolutionInversionCount(t *testing.T) {
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

		input := []byte("5\n5 4 3 2 1\n")
		w.Write(input)

		os.Stdin, os.Stdout = r, w
		module2.SolutionInversionCount()
		w.Close()

		out, err := io.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		expectedOutput := []byte("10")
		assert.Equal(expectedOutput, out)
	})
	t.Run("Case: no inversions", func(t *testing.T) {
		r, w, err := os.Pipe()
		if err != nil {
			log.Fatal(err)
		}

		input := []byte("5\n1 2 3 4 5\n")
		w.Write(input)

		os.Stdin, os.Stdout = r, w
		module2.SolutionInversionCount()
		w.Close()

		out, err := io.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		expectedOutput := []byte("0")
		assert.Equal(expectedOutput, out)
	})
}
