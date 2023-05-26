package module5

import (
	"bufio"
	"fmt"
	"os"
)

type height int

func (b *BinarySearchTree) isBalanced() bool {
	if b.root == nil {
		return true
	}
	return b.getAndCompareHeight(b.root) != -1
}

// -1 -> unbalanced
func (b *BinarySearchTree) getAndCompareHeight(treeNode *Node) height {
	if treeNode == nil {
		return 0
	}

	var left height = b.getAndCompareHeight(treeNode.left)
	if left == -1 {
		return -1
	}

	var right height = b.getAndCompareHeight(treeNode.right)
	if right == -1 {
		return -1
	}

	if abs(left-right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

func abs(input height) height {
	if input < 0 {
		return -input
	}
	return input
}

func max(a, b height) height {
	if a > b {
		return a
	}
	return b
}

func SolutionTask2() {
	reader := bufio.NewReader(os.Stdin)

	tree := BinarySearchTree{}
	number := 0
	fmt.Fscan(reader, &number)
	for number != 0 {
		tree.InsertNode(number)
		fmt.Fscan(reader, &number)
	}

	if result := tree.isBalanced(); result {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
