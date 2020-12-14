package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerkleTree(t *testing.T) {
	merkleTree := MerkleTree{}

	merkleTree.InitFromArray([]string{"1"})
	assert.Equal(t, 1, merkleTree.height)
}
