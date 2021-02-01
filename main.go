package main

import (
	"BLC/model"
	"fmt"
)

func main() {
	//	block := model.NewBlock(1, nil, []byte("the first block testing"))
	//	fmt.Printf("the first block : %v\n", block)
	bc := model.CreateBlockChainWithGenesisBlock()
	//	fmt.Printf("first Block %v\n", bc.Blocks[0])
	bc.AddBlock([]byte("a-b"))
	for _, a := range bc.Blocks {

		fmt.Printf("hash: %x %x\n", string(a.Data), a.Hash)

	}

}
