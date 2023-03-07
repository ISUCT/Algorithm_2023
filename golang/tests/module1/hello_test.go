package module1_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module1"
	"isuct.ru/informatics2022/tests/helpers"
)

func TestSayHello(t *testing.T) {
	assert := assert.New(t)

	oldStdin := *os.Stdin
	oldStdout := *os.Stdout
	defer func(v *os.File) { os.Stdin = v }(&oldStdin)
	defer func(v *os.File) { os.Stdout = v }(&oldStdout)

	t.Run("Case: Ubuntu", func(t *testing.T) {
		r, w := helpers.Replacer("Ubuntu\n", t)
		os.Stdin = r
		os.Stdout = w

		module1.SayHello()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal("Hello, Ubuntu!", string(out))
	})
	t.Run("Case: input 100 characters long", func(t *testing.T) {
		r, w := helpers.Replacer("lkryayyxzbhjjvcenprwtjpkqblthvqfwcolfonjqhsfvgxpfdkqhanrzlilwyoaiopoaipqxeknuskfjkrvkqpkisehqoawmenn\n", t)
		os.Stdin = r
		os.Stdout = w

		module1.SayHello()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal("Hello, lkryayyxzbhjjvcenprwtjpkqblthvqfwcolfonjqhsfvgxpfdkqhanrzlilwyoaiopoaipqxeknuskfjkrvkqpkisehqoawmenn!", string(out))
	})
}
