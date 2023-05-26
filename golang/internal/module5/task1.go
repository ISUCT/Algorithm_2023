package module5

import (
	"bufio"
	"fmt"
	"os"
)

func (b *BinarySearchTree) InOrderPrintNodesOneChild(node *Node, writer *bufio.Writer) {
	if node == nil {
		return
	}

	b.InOrderPrintNodesOneChild(node.left, writer)
	if (node.left == nil && node.right != nil) || (node.left != nil && node.right == nil) {
		fmt.Fprintln(writer, node.value)
	}
	b.InOrderPrintNodesOneChild(node.right, writer)

	writer.Flush()
}

func SolutionTask1() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	tree := BinarySearchTree{}
	number := 0
	fmt.Fscan(reader, &number)
	for number != 0 {
		tree.InsertNode(number)
		fmt.Fscan(reader, &number)
	}

	tree.InOrderPrintNodesOneChild(tree.GetTreeRoot(), writer)
}
