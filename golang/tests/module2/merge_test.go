package module2_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module2"
	"isuct.ru/informatics2022/tests/helpers"
)

func TestMergeSort(t *testing.T) {
	assert := assert.New(t)

	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	t.Run("Test Merge", func(t *testing.T) {
		t.Run("Test equal length", func(t *testing.T) {
			a := []module2.Item{
				{
					Index: 1,
					Value: 1,
				},
				{
					Index: 2,
					Value: 3,
				},
				{
					Index: 3,
					Value: 5,
				},
			}
			b := []module2.Item{
				{
					Index: 1,
					Value: 2,
				},
				{
					Index: 2,
					Value: 4,
				},
				{
					Index: 3,
					Value: 6,
				},
			}
			expected := []module2.Item{
				{
					Index: 1,
					Value: 1,
				},
				{
					Index: 1,
					Value: 2,
				},
				{
					Index: 2,
					Value: 3,
				},
				{
					Index: 2,
					Value: 4,
				},
				{
					Index: 3,
					Value: 5,
				},
				{
					Index: 3,
					Value: 6,
				},
			}
			result := module2.Merge(a, b)
			assert.EqualValues(expected, result)
		})
		// t.Run("Test first zero lenght", func(t *testing.T) {
		// 	a := []int{}
		// 	b := []int{2, 4, 6}
		// 	expected := []int{2, 4, 6}
		// 	result := module2.Merge(a, b)
		// 	assert.EqualValues(expected, result)
		// })
		// t.Run("Test second zero lenght", func(t *testing.T) {
		// 	a := []int{1, 3, 5}
		// 	b := []int{}
		// 	expected := []int{1, 3, 5}
		// 	result := module2.Merge(a, b)
		// 	assert.EqualValues(expected, result)
		// })
		// t.Run("Test second values less than first", func(t *testing.T) {
		// 	a := []int{1, 3, 5, 6, 7, 8}
		// 	b := []int{2, 4, 6}
		// 	expected := []int{1, 2, 3, 4, 5, 6, 6, 7, 8}
		// 	result := module2.Merge(a, b)
		// 	assert.EqualValues(expected, result)
		// })
	})
	t.Run("Test Sort", func(t *testing.T) {
		t.Run("Test array with even len", func(t *testing.T) {
			a := []module2.Item{
				{
					Index: 1,
					Value: 5,
				},
				{
					Index: 2,
					Value: 4,
				},
				{
					Index: 3,
					Value: 3,
				},
				{
					Index: 4,
					Value: 2,
				},
			}
			expected := []module2.Item{
				{
					Index: 4,
					Value: 2,
				},
				{
					Index: 3,
					Value: 3,
				},
				{
					Index: 2,
					Value: 4,
				},
				{
					Index: 1,
					Value: 5,
				},
			}
			result := module2.Sort(a)
			assert.EqualValues(expected, result)
		})
		// t.Run("Test array with odd len", func(t *testing.T) {
		// 	a := []int{5, 4, 3, 2, 1}
		// 	expected := []int{1, 2, 3, 4, 5}
		// 	result := module2.Sort(a)
		// 	assert.EqualValues(expected, result)
		// })
	})
	t.Run("Test merge sort", func(t *testing.T) {
		r, w := helpers.Replacer(`5
5 4 3 2 1
`, t)
		os.Stdin = r
		os.Stdout = w

		module2.MergeSort()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`1 2 4 5
4 5 1 2
3 5 1 3
1 5 1 5
1 2 3 4 5
`, string(out))
	})
}
