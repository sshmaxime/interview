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

func (m *MerkleTree) Level(level int) [][]byte {
	if level < 1 || level > m.height {
		return [][]byte{}
	}
	return getHashAtLevel(m.root, [][]byte{}, level, 1)
}
