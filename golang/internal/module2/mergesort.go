package module2

//"fmt"

func merge(left []int, right []int) []int {
	lst := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			lst = append(lst, left[0])
			left = left[1:]
		} else {
			lst = append(lst, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		lst = append(lst, left...)
	}
	if len(right) > 0 {
		lst = append(lst, right...)
	}

	return lst
}

func MergeSort(arr []int, left int, right int) []int {
	var res []int
	if right-left == 1 {
		res = make([]int, 1)
		res[0] = arr[left]
		return res
	}
	middle := (left + right) / 2
	leftPart := MergeSort(arr, left, middle)
	rightPart := MergeSort(arr, middle, right)
	res = merge(leftPart, rightPart)
	//fmt.Println(left+1, right, res[0], res[len(res)-1])
	return res
}
