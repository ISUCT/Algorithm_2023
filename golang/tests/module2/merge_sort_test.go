package module2_test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"isuct.ru/informatics2022/internal/module2"
	m2 "isuct.ru/informatics2022/internal/module2"
	"isuct.ru/informatics2022/tests/helpers"
)

func TestMergeSort(t *testing.T) {
	require := require.New(t)
	t.Run("Test merge", func(t *testing.T) {
		expected := []int{1, 2, 3, 4}
		t.Run("Test merge equal length", func(t *testing.T) {
			a := []int{1, 3}
			b := []int{2, 4}
			result := m2.Merge(a, b)
			require.EqualValues(expected, result)
		})
		t.Run("Test merge first shorter length", func(t *testing.T) {
			a := []int{1}
			b := []int{2, 3, 4}
			result := m2.Merge(a, b)
			require.EqualValues(expected, result)
		})
		t.Run("Test merge second short length", func(t *testing.T) {
			a := []int{1, 2, 3}
			b := []int{4}
			result := m2.Merge(a, b)
			require.EqualValues(expected, result)
		})
	})

	t.Run("Test merge sort", func(t *testing.T) {
		defer func(v *os.File) { os.Stdin = v }(os.Stdin)
		defer func(v *os.File) { os.Stdout = v }(os.Stdout)

		t.Run("Test single item array", func(t *testing.T) {
			r, w := helpers.Replacer(`1
1
`, t)
			os.Stdin = r
			os.Stdout = w

			module2.MergeTask()
			w.Close()
			out, _ := io.ReadAll(r)
			require.Equal(`1
`, string(out))
		})
		t.Run("Test two items array", func(t *testing.T) {
			r, w := helpers.Replacer(`2
3 1
`, t)
			os.Stdin = r
			os.Stdout = w

			module2.MergeTask()
			w.Close()
			out, _ := io.ReadAll(r)
			require.Equal(`1 2 1 3
1 3
`, string(out))
		})
		t.Run("Test 5 items array", func(t *testing.T) {
			r, w := helpers.Replacer(`5
5 4 3 2 1
`, t)
			os.Stdin = r
			os.Stdout = w

			module2.MergeTask()
			w.Close()
			out, _ := io.ReadAll(r)
			require.Equal(`1 2 4 5
4 5 1 2
3 5 1 3
1 5 1 5
1 2 3 4 5
`, string(out))
		})
		t.Run("Test 100000 items array", func(t *testing.T) {
			r, w := helpers.Replacer(`10000
`, t)
			os.Stdin = r
			os.Stdout = w

			writer := bufio.NewWriterSize(w, int(10000))

			for i := 9999; i >= 0; i-- {
				fmt.Fprintf(writer, "%d ", i)
			}
			fmt.Fprint(writer, "\n")
			writer.Flush()

			module2.MergeTask()
			w.Close()
			out, _ := io.ReadAll(r)
			require.Equal(`1 2 4 5
4 5 1 2
3 5 1 3
1 5 1 5
1 2 3 4 5
`, string(out))
		})
	})

}
