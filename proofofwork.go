package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 工作量证明
type ProofOfWork struct {
	// block
	block *Block
	// 目标值
	target *big.Int
}

const targetStr = 20

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	// 左移20位
	target.Lsh(target, uint(256-targetStr))

	return &ProofOfWork{
		block:  block,
		target: target,
	}
}

func (pow *ProofOfWork) Run() ([]byte, uint64) {
	// join data
	var nonce uint64
	var block Block = *pow.block
	var hash [32]byte

	fmt.Println("proofOfWorking")
	for {
		// 拼装数据
		tmp := [][]byte{
			block.PrevHash,
			block.Data,
			UintToByte(block.Version),
			UintToByte(nonce),
			UintToByte(block.Difficulty),
			block.MarkelRoot,
		}

		blockInfo := bytes.Join(tmp, []byte{})

		// hash 运算
		hash = sha256.Sum256(blockInfo)
		fmt.Printf("\r%x", hash)

		// 比较
		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])
		if tmpInt.Cmp(pow.target) == -1 {
			break
		}

		nonce++
	}

	fmt.Println()

	return hash[:], nonce
}
