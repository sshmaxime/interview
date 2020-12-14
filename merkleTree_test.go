package main

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func toHex(str string) []byte {
	hash, _ := hex.DecodeString(str)
	return hash
}

func TestMerkleTree(t *testing.T) {
	merkleTree := MerkleTree{}

	merkleTree.InitFromArray([]string{"1"})
	assert.Equal(t, 1, merkleTree.height)
	assert.Equal(t, "6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b", hex.EncodeToString(merkleTree.root.hash))

	merkleTree.InitFromArray([]string{"1", "2", "3", "4"})
	assert.Equal(t, 3, merkleTree.height)
	assert.Equal(t, toHex("cd53a2ce68e6476c29512ea53c395c7f5d8fbcb4614d89298db14e2a5bdb5456"), merkleTree.root.hash)
	assert.Equal(t, [][]byte{
		toHex("6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b"),
		toHex("d4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35"),
		toHex("4e07408562bedb8b60ce05c1decfe3ad16b72230967de01f640b7e4729b49fce"),
		toHex("4b227777d4dd1fc61c6f884f48641d02b4d121d3fd328cb08b5531fcacdabf8a")}, merkleTree.Level(3))

}
