package main

import (
	"BLC/model"
	"fmt"
)

func main() {
	block := model.NewBlock(1, nil, []byte("the first block testing"))
	fmt.Printf("the first block : %v\n", block)
}
