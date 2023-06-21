package module3_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module3"
	"isuct.ru/informatics2022/tests/helpers"
)

func TestTask1(t *testing.T) {
	assert := assert.New(t)

	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	t.Run("Case: ababbababa-aba", func(t *testing.T) {
		r, w := helpers.Replacer(`ababbababa
aba
`, t)
		os.Stdin = r
		os.Stdout = w

		module3.Task1()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`0 5 7 
`, string(out))
	})
}
