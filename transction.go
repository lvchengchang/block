package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

const reword = 12.5

type Transaction struct {
	TXId      []byte     // transaction id
	TXInputs  []TXInput  // transaction input
	TXOutputs []TXOutput // transaction output
}

// define transaction input

type TXInput struct {
	// transaction id
	TXid []byte
	// transaction index
	index int64
	// address 解锁脚本
	Sig string // 发起证明(我拥有这些币)
}

type TXOutput struct {
	value float64 // 金额是多少
	// 锁定脚本
	PukkeyHash string // 证明发起交易的权限
}

// set transaction id
func (tx *Transaction) SetHash() {
	var buffer bytes.Buffer

	encode := gob.NewEncoder(&buffer)
	err := encode.Encode(tx)
	if err != nil {
		log.Panicln(err)
	}

	data := buffer.Bytes()
	hash := sha256.Sum256(data)

	tx.TXId = hash[:]
}

// 挖矿只有一个input和output
func NewCoinBaseTx(address string, data string) *Transaction {
	input := TXInput{[]byte{}, -1, data}
	output := TXOutput{reword, address}

	tx := Transaction{[]byte{}, []TXInput{input}, []TXOutput{output}}
	tx.SetHash()

	return &tx
}
