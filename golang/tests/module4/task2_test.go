package module4_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module4"
	"isuct.ru/informatics2022/tests/helpers"
)

func TestTask2(t *testing.T) {
	assert := assert.New(t)

	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	t.Run("Case: 3 2 1", func(t *testing.T) {
		r, w := helpers.Replacer(`3
3 2 1
`, t)
		os.Stdin = r
		os.Stdout = w

		module4.Task2()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`YES
`, string(out))
	})

	t.Run("Case: 4 1 3 2", func(t *testing.T) {
		r, w := helpers.Replacer(`4
4 1 3 2
`, t)
		os.Stdin = r
		os.Stdout = w

		module4.Task2()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`YES
`, string(out))
	})

	t.Run("Case: 2 3 1", func(t *testing.T) {
		r, w := helpers.Replacer(`3
2 3 1
`, t)
		os.Stdin = r
		os.Stdout = w

		module4.Task2()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`NO
`, string(out))
	})
}
