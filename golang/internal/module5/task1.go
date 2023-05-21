package module5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{value: value}
	}

	if value < root.value {
		root.left = insert(root.left, value)
	} else {
		root.right = insert(root.right, value)
	}

	return root
}

func print_single_child_nodes(root *Node) {
	if root == nil {
		return
	}

	if root.left != nil && root.right == nil {
		fmt.Println(root.value)
		print_single_child_nodes(root.left)
	} else if root.left == nil && root.right != nil {
		fmt.Println(root.value)
		print_single_child_nodes(root.right)
	} else {
		print_single_child_nodes(root.left)
		print_single_child_nodes(root.right)
	}
}

func Task1() {
	var root *Node

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	arr := strings.Split(line, " ")

	n := len(arr) - 1
	value := make([]int, n)
	for i := 0; i < n; i++ {
		value[i], _ = strconv.Atoi(arr[i])
		root = insert(root, value[i])
		fmt.Scan(&value)
	}

	print_single_child_nodes(root)
}
