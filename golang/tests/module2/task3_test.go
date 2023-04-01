package module2_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module2"
	"isuct.ru/informatics2022/tests/helpers"
)

func TestTask3(t *testing.T) {
	assert := assert.New(t)

	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	t.Run("Case: 1 1", func(t *testing.T) {
		r, w := helpers.Replacer(`1
1
`, t)
		os.Stdin = r
		os.Stdout = w

		module2.Task3()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`1 `, string(out))
	})
	t.Run("Case: 2 3 1", func(t *testing.T) {
		r, w := helpers.Replacer(`2
3 1
`, t)
		os.Stdin = r
		os.Stdout = w

		module2.Task3()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`1 2 1 3
1 3 `, string(out))
	})
	t.Run("Case: 5 5 4 3 2 1", func(t *testing.T) {
		r, w := helpers.Replacer(`5
5 4 3 2 1
`, t)
		os.Stdin = r
		os.Stdout = w

		module2.Task3()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`1 2 4 5
4 5 1 2
3 5 1 3
1 5 1 5
1 2 3 4 5 `, string(out))
	})
}
