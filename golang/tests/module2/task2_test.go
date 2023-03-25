package module2_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module2"
	"isuct.ru/informatics2022/tests/helpers"
)

func TestTask2(t *testing.T) {
	assert := assert.New(t)

	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	t.Run("Case: 3 101 80", func(t *testing.T) {
		r, w := helpers.Replacer(`3
101 80
305 90
200 14		
`, t)
		os.Stdin = r
		os.Stdout = w

		module2.Task2()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`305 90
101 80
200 14		
`, string(out))
	})
	t.Run("Case: 3 20 80", func(t *testing.T) {
		r, w := helpers.Replacer(`3
20 80
30 90
25 90
`, t)
		os.Stdin = r
		os.Stdout = w

		module2.Task2()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`25 90
30 90
20 80		
`, string(out))
	})
}
