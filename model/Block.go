package model

import (
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	TimeStamp    int64
	Hash         []byte
	PreBlockHash []byte
	Height       int64
	Data         []byte
}

func NewBlock(height int64, prevBlockHash []byte, data []byte) *Block {
	block := Block{
		TimeStamp:    time.Now().Unix(),
		Hash:         nil,
		PreBlockHash: prevBlockHash,
		Height:       height,
		Data:         data,
	}
	block.SetHash()
	return &block
}

func (b *Block) SetHash() {
	timeStampBytes := IntToHex(b.TimeStamp)
	heigthBytes := IntToHex(b.Height)
	blockBytes := bytes.Join([][]byte{
		heigthBytes,
		timeStampBytes,
		b.PreBlockHash,
		b.Data,
	}, []byte{})

	hash := sha256.Sum256(blockBytes)
	b.Hash = hash[:]

}

func CreateGenesisBlock(data []byte) *Block {
	return NewBlock(1, nil, data)

}
