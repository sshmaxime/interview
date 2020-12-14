package main

import "fmt"

type merkleTreeNode struct {
	hash  []byte
	value string
	left  *merkleTreeNode
	right *merkleTreeNode
}

type MerkleTree struct {
	root   *merkleTreeNode
	height int
}

func (m *MerkleTree) InitFromArray(input []string) {
	hashTreeArray := CreateHashTreeArrayFromArray(input, len(input))
	m.root, m.height = CreateMerkleTreeFromHashTreeArray(hashTreeArray, 2)

}

func (m *MerkleTree) Root() *merkleTreeNode {
	return m.root
}

func (m *MerkleTree) Height() int {
	return m.height
}

func (m *MerkleTree) DisplayLevel(level int) {
	arr := Level(m.root, [][]byte{}, level, 1)

	for _, val := range arr {
		fmt.Printf("%x\n", val)
	}
}
func Level(node *merkleTreeNode, array [][]byte, level int, indexLevel int) [][]byte {
	if indexLevel == level {
		return append(array, node.hash)
	}
	if node.left != nil {
		array = Level(node.left, array, level, indexLevel+1)
	}
	if node.right != nil {
		array = Level(node.right, array, level, indexLevel+1)
	}
	return array
}
