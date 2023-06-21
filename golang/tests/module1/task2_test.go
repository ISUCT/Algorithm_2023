package module1_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module1"
	"isuct.ru/informatics2022/tests/helpers"
)

func TestTask2(t *testing.T) {
	assert := assert.New(t)

	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	t.Run("Case: Vasya", func(t *testing.T) {
		r, w := helpers.Replacer("Vasya\n", t)
		os.Stdin = r
		os.Stdout = w

		module1.Task2()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal("Hello, Vasya!", string(out))
	})
	t.Run("Case: Petya", func(t *testing.T) {
		r, w := helpers.Replacer("Petya\n", t)
		os.Stdin = r
		os.Stdout = w

		module1.Task2()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal("Hello, Petya!", string(out))
	})
}
