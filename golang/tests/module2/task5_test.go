package module2_test

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module2"
	"isuct.ru/informatics2022/tests/helpers"
)

func TestSolutionCountDifferentElems(t *testing.T) {
	assert := assert.New(t)
	t.Run("Case: one number", func(t *testing.T) {
		input, output := helpers.GetInputAndOutputFiles()
		input.WriteString("1\n1")
		module2.SolutionCountDifferentElems()
		got, err := io.ReadAll(output)
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal("1", string(got))
	})
	t.Run("Case: one positive and one negative number with equal with equal modulo", func(t *testing.T) {
		input, output := helpers.GetInputAndOutputFiles()
		input.WriteString("2\n-5 5")
		module2.SolutionCountDifferentElems()
		got, err := io.ReadAll(output)
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal("2", string(got))
	})
	t.Run("Case: unsorted array with 10000 numbers", func(t *testing.T) {
		input, output := helpers.GetInputAndOutputFiles()
		testData, _ := os.OpenFile("testData\\task5.txt", os.O_RDONLY, 0755)
		io.Copy(input, testData)
		module2.SolutionCountDifferentElems()
		got, err := io.ReadAll(output)
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal("10000", string(got))
	})
	t.Run("Case: unsorted positive and negative numbers", func(t *testing.T) {
		input, output := helpers.GetInputAndOutputFiles()
		input.WriteString("11\n58423 -5 5 10 2835 2 -1 5 -2372 3 -2")
		module2.SolutionCountDifferentElems()
		got, err := io.ReadAll(output)
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal("10", string(got))
	})
}
