package main

import "math/big"

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

	pow := ProofOfWork{
		block: block,
	}
	pow.target = target
	return &pow
}
