package main

import "fmt"

func main() {
	merkleTree := MerkleTree{}
	merkleTree.InitFromArray([]string{"1"})
	fmt.Printf("%x\n", merkleTree.root.hash)
	fmt.Printf("%x\n", merkleTree.height)
	merkleTree.DisplayLevel(1)
}
