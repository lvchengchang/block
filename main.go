package main

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash []byte

	PrevHash []byte

	Data []byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		Hash:     []byte{}, // 空，后面计算 todo
		Data:     []byte(data),
	}

	block.SetHash()

	return &block
}

func (block *Block) SetHash() {
	// 拼装数据进行sha256
	// 平铺Data数组，拼接起来
	blockInfo := append(block.PrevHash, block.Data...)
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

func main() {
	block := NewBlock("转了一个比特币", []byte{})

	fmt.Printf("前一个区块的哈希值是: %x\n", block.PrevHash)
	fmt.Printf("当前区块的哈希值是: %x\n", block.Hash)
	fmt.Printf("区块数据是:%s\n", block.Data)
}
