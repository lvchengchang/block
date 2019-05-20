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

type BlockChain struct {
	blocks []*Block
}

// create block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		Hash:     []byte{}, // 空，后面计算 todo
		Data:     []byte(data),
	}

	block.SetHash()

	return &block
}

// create blockchain
func NewBlockChain() *BlockChain {
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

func GenesisBlock() *Block {
	return NewBlock("this is first block", []byte{})
}

func (block *Block) SetHash() {
	// 拼装数据进行sha256
	// 平铺Data数组，拼接起来
	blockInfo := append(block.PrevHash, block.Data...)
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

func main() {
	bc := NewBlockChain()

	for i, v := range bc.blocks {
		fmt.Printf("当前区块高度是:--------- %d\n", i)
		fmt.Printf("前一个区块的哈希值是: %x\n", v.PrevHash)
		fmt.Printf("当前区块的哈希值是: %x\n", v.Hash)
		fmt.Printf("区块数据是:%s\n", v.Data)
	}
}
