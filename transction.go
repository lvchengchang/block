package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
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
	Value float64 // 金额是多少
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

func (tx *Transaction) IsCoinBase() bool {
	if len(tx.TXInputs) == 1 && len(tx.TXInputs[0].TXid) == 0 && tx.TXInputs[0].index == -1 {
		return true
	}

	return false
}

func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {
	utxos, resValue := bc.FindNeedUTXOs(from, amount)

	if resValue < amount {
		fmt.Println("余额不足")
		return nil
	}

	var inputs []TXInput
	var outputs []TXOutput

	for id, indexArray := range utxos {
		for _, i := range indexArray {
			input := TXInput{[]byte(id), int64(i), from}
			inputs = append(inputs, input)
		}
	}

	output := TXOutput{amount, to}
	outputs = append(outputs, output)
	// 找零
	if resValue > amount {
		outputs = append(outputs, TXOutput{resValue - amount, from})
	}

	tx := Transaction{[]byte{}, inputs, outputs}
	return &tx
}
