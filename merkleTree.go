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
	hashTreeArray := CreateHashTreeArrayFromArray(input)
	m.root, m.height = CreateMerkleTreeFromHashTreeArray(hashTreeArray, 1)
}

func (m *MerkleTree) Root() *merkleTreeNode {
	return m.root
}

func (m *MerkleTree) Height() int {
	return m.height
}

func (m *MerkleTree) DisplayLevel(level int) {
	arr := m.Level(level)

	for _, val := range arr {
		fmt.Printf("%x\n", val)
	}
}

func getLevel(node *merkleTreeNode, array [][]byte, level int, indexLevel int) [][]byte {
	if indexLevel == level {
		return append(array, node.hash)
	}
	if node.left != nil {
		array = getLevel(node.left, array, level, indexLevel+1)
	}
	if node.right != nil {
		array = getLevel(node.right, array, level, indexLevel+1)
	}
	return array
}

func (m *MerkleTree) Level(level int) [][]byte {
	return getLevel(m.root, [][]byte{}, level, 1)
}
