package main

import (
	"crypto/sha256"
)

func createHash(value []byte) []byte {
	h := sha256.New()
	h.Write(value)
	return h.Sum(nil)
}

func CreateHashTreeArrayFromArray(inputs []string) []merkleTreeNode {
	var nodes []merkleTreeNode

	for _, input := range inputs {
		node := merkleTreeNode{
			hash:  createHash([]byte(input)),
			left:  nil,
			right: nil,
		}
		nodes = append(nodes, node)
	}
	return nodes
}

func CreateMerkleTreeFromHashTreeArray(hashTreeArray []merkleTreeNode, index int) (*merkleTreeNode, int) {
	//PrintHashTreeArray(hashTreeArray)
	if len(hashTreeArray) == 1 {
		return &hashTreeArray[0], index
	}
	var newHashTreeArray []merkleTreeNode
	for i := 0; i < len(hashTreeArray); i = i + 2 {
		var hash []byte
		var value string
		var left *merkleTreeNode
		var right *merkleTreeNode

		if i+1 < len(hashTreeArray) {
			hash = createHash(append(hashTreeArray[i].hash, hashTreeArray[i+1].hash...))
			value = hashTreeArray[i].value + hashTreeArray[i+1].value
			left = &hashTreeArray[i]
			right = &hashTreeArray[i+1]
		} else {
			hash = hashTreeArray[i].hash
			value = hashTreeArray[i].value
			left = &hashTreeArray[i]
			right = nil
		}
		node := merkleTreeNode{
			hash:  hash,
			value: value,
			left:  left,
			right: right,
		}
		newHashTreeArray = append(newHashTreeArray, node)
	}
	return CreateMerkleTreeFromHashTreeArray(newHashTreeArray, index+1)
}

func getHashAtLevel(node *merkleTreeNode, array [][]byte, level int, indexLevel int) [][]byte {
	if indexLevel == level {
		return append(array, node.hash)
	}
	if node.left != nil {
		array = getHashAtLevel(node.left, array, level, indexLevel+1)
	}
	if node.right != nil {
		array = getHashAtLevel(node.right, array, level, indexLevel+1)
	}
	return array
}
