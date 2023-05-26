package module2_test

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module2"
)

func TestSolutionStorageProblem(t *testing.T) {
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

		input, _ := os.ReadFile("testData/task6/task6_usualinput.txt")
		w.Write(input)

		os.Stdin, os.Stdout = r, w
		module2.SolutionStorageProblem()
		w.Close()

		out, err := io.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		expectedOutput, _ := os.ReadFile("testData/task6/task6_usualoutput.txt")
		assert.Equal(expectedOutput, out)
	})
}
