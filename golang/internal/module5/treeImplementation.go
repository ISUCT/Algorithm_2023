package module5

import (
	"bufio"
	"errors"
	"fmt"
)

type BinarySearchTree struct {
	root       *Node
	comparator *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

type direction int

const (
	left  direction = 0
	right           = 1
)

var alreadyExists error = errors.New("This element is already inserted in the tree")

func (b *BinarySearchTree) InsertNode(element int) error {
	if b.root == nil {
		b.root = &Node{value: element}
	}

	b.comparator = b.root
	notInserted := true

	for notInserted {
		if element > b.comparator.value {
			if b.isNextNodeNil(right) {
				b.comparator.right = &Node{value: element}
				notInserted = false
			}
		} else if element < b.comparator.value {
			if b.isNextNodeNil(left) {
				b.comparator.left = &Node{value: element}
				notInserted = false
			}
		} else {
			return alreadyExists
		}
	}

	return nil
}

func (b *BinarySearchTree) isNextNodeNil(flag direction) (nextNodeIsNil bool) {
	if nextNode := b.nextNode(flag); nextNode == nil {
		nextNodeIsNil = true
	} else {
		b.comparator = nextNode
	}

	return nextNodeIsNil
}

func (b *BinarySearchTree) nextNode(flag direction) *Node {
	if flag == left {
		return b.comparator.left
	}

	return b.comparator.right
}

func (b *BinarySearchTree) PreOrderPrint(node *Node, writer *bufio.Writer) {
	if node == nil {
		return
	}

	fmt.Fprintln(writer, node.value)
	b.PreOrderPrint(node.left, writer)
	b.PreOrderPrint(node.right, writer)

	writer.Flush()
}

func (b *BinarySearchTree) InOrderPrint(node *Node, writer *bufio.Writer) {
	if node == nil {
		return
	}

	b.InOrderPrint(node.left, writer)
	fmt.Fprintln(writer, node.value)
	b.InOrderPrint(node.right, writer)

	writer.Flush()
}

func (b *BinarySearchTree) PostOrderPrint(node *Node, writer *bufio.Writer) {
	if node == nil {
		return
	}

	b.PostOrderPrint(node.left, writer)
	b.PostOrderPrint(node.right, writer)
	fmt.Fprintln(writer, node.value)

	writer.Flush()
}

func (b *BinarySearchTree) GetTreeRoot() *Node {
	return b.root
}
