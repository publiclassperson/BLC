package model

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

type Block struct {
	TimeStamp    int64
	Hash         []byte
	PreBlockHash []byte
	Heigth       int64
	Data         []byte
}

func NewBlock(height int64, prevBlockHash []byte, data []byte) *Block {
	block := Block{
		TimeStamp:    time.Now().Unix(),
		Hash:         nil,
		PreBlockHash: prevBlockHash,
		Heigth:       height,
		Data:         data,
	}
	block.SetHash()
	return &block
}

func (b *Block) SetHash() {
	timeStampBytes := IntToHex(b.TimeStamp)
	heigthBytes := IntToHex(b.Heigth)
	blockBytes := bytes.Join([][]byte{
		heigthBytes,
		timeStampBytes,
		b.PreBlockHash,
		b.Data,
	}, []byte{})

	hash := sha256.Sum256(blockBytes)
	b.Hash = hash[:]

}

func IntToHex(data int64) []byte {

	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, data)
	if nil != err {
		log.Printf("int transact to []byte failed! %v\n", err)
	}
	return buffer.Bytes()
}
