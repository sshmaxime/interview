package main

import "fmt"

type binaryTreeNode struct {
	data  string
	left  *binaryTreeNode
	right *binaryTreeNode
}

func createNode(value string) *binaryTreeNode {
	return &binaryTreeNode{
		data:  value,
		left:  nil,
		right: nil,
	}
}

// Formula:
// Left child of i = 2 * i + 1
// Right child of i = 2 * i + 2
func CreateBinaryTreeFromArray(input []string, node *binaryTreeNode, i int, arrayLength int) *binaryTreeNode {
	if i < arrayLength {
		node = createNode(input[i])

		node.left = CreateBinaryTreeFromArray(input, node.left, 2*i+1, arrayLength)
		node.right = CreateBinaryTreeFromArray(input, node.right, 2*i+2, arrayLength)
	}
	return node
}

// Display Child of Binary Tree
func DisplayChildOf(node *binaryTreeNode, value string) {
	if node == nil {
		return
	}
	if node.data == value {
		fmt.Println(node.left)
		fmt.Println(node.right)
	} else {
		DisplayChildOf(node.left, value)
		DisplayChildOf(node.right, value)
	}
}
