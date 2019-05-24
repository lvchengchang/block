package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

type Block struct {
	Version    uint64
	Hash       []byte
	PrevHash   []byte
	Data       []byte
	MarkelRoot []byte
	TimeStamp  uint64
	Nonce      uint64
	Difficulty uint64
}

func UintToByte(num uint64) []byte {
	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}

// create block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevBlockHash,
		Hash:       []byte{},
		Data:       []byte(data),
		MarkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0, // 无效值
		Nonce:      0,
	}

	//block.SetHash()

	pow := NewProofOfWork(&block)
	// search random and hash
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return &block
}

func (block *Block) SetHash() {
	// 拼装数据进行sha256
	// 平铺Data数组，拼接起来
	tmp := [][]byte{
		block.PrevHash,
		block.Data,
		UintToByte(block.Version),
		UintToByte(block.Nonce),
		UintToByte(block.Difficulty),
		block.MarkelRoot,
	}

	blockInfo := bytes.Join(tmp, []byte{})

	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

func (b *Block) toByte() []byte {
	return []byte{}
}
