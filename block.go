package main

import (
	"crypto/sha256"
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

type BlockChain struct {
	blocks []*Block
}

func UintToByte(num uint64) []byte {
	return []byte{}
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

	block.SetHash()

	return &block
}

func (block *Block) SetHash() {
	// 拼装数据进行sha256
	// 平铺Data数组，拼接起来
	blockInfo := append(block.PrevHash, block.Data...)
	blockInfo = append(blockInfo, UintToByte(block.Version)...)
	blockInfo = append(blockInfo, UintToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo, UintToByte(block.Nonce)...)
	blockInfo = append(blockInfo, UintToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, block.MarkelRoot...)

	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
