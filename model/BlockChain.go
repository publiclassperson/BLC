package model

type BlockChain struct {
	Blocks []*Block
}

func CreateBlockChainWithGenesisBlock() *BlockChain {
	block := CreateGenesisBlock([]byte("init block"))
	return &BlockChain{[]*Block{block}}
}

func (bc *BlockChain) AddBlock(data []byte) {
	height := bc.Blocks[(len(bc.Blocks)-1)].Height + 1
	preBlockHash := bc.Blocks[(len(bc.Blocks) - 1)].Hash
	newBlock := NewBlock(height, preBlockHash, data)
	bc.Blocks = append(bc.Blocks, newBlock)

}
