package module3_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module3"
	"isuct.ru/informatics2022/tests/helpers"
)

func TestTask3(t *testing.T) {
	assert := assert.New(t)

	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	t.Run("Case: 1-aaaaa", func(t *testing.T) {
		r, w := helpers.Replacer(`aaaaa
`, t)
		os.Stdin = r
		os.Stdout = w

		module3.Task3()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`5
`, string(out))
	})

	t.Run("Case: 2-abcabcabc", func(t *testing.T) {
		r, w := helpers.Replacer(`abcabcabc
`, t)
		os.Stdin = r
		os.Stdout = w

		module3.Task3()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`3
`, string(out))
	})

	t.Run("Case: 3-abab", func(t *testing.T) {
		r, w := helpers.Replacer(`abab
`, t)
		os.Stdin = r
		os.Stdout = w

		module3.Task3()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`2
`, string(out))
	})
	t.Run("Case: 17-babbbabb", func(t *testing.T) {
		r, w := helpers.Replacer(`babbbabb
`, t)
		os.Stdin = r
		os.Stdout = w

		module3.Task3()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`2
`, string(out))
	})
}
