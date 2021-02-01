package model

import (
	"bytes"
	"encoding/binary"
	"log"
)

func IntToHex(data int64) []byte {

	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, data)
	if nil != err {
		log.Printf("int transact to []byte failed! %v\n", err)
	}
	return buffer.Bytes()
}
