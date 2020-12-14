package main

import "fmt"

func main() {
	merkleTree := MerkleTree{}
	merkleTree.InitFromArray([]string{"1", "2", "3", "4"})
	fmt.Printf("%x\n", merkleTree.root.hash)
}

//1- 6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b
//2- d4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35
//3- 4e07408562bedb8b60ce05c1decfe3ad16b72230967de01f640b7e4729b49fce
//4- 4b227777d4dd1fc61c6f884f48641d02b4d121d3fd328cb08b5531fcacdabf8a

//12- 4295f72eeb1e3507b8461e240e3b8d18c1e7bd2f1122b11fc9ec40a65894031a
//34- 20ab747d45a77938a5b84c2944b8f5355c49f21db0c549451c6281c91ba48d0d

//1234- cd53a2ce68e6476c29512ea53c395c7f5d8fbcb4614d89298db14e2a5bdb5456
