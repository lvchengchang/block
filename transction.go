package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const reword = 12.5 // 挖矿奖励

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
	Index int64
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
	input := TXInput{[]byte{}, -1, data} // 挖矿没有输入，全部给个默认值
	output := TXOutput{reword, address}  // 输出是金额和地址

	tx := Transaction{[]byte{}, []TXInput{input}, []TXOutput{output}}
	// 此步操作是把结构体hash成TXID
	tx.SetHash()

	return &tx
}

func (tx *Transaction) IsCoinBase() bool {
	if len(tx.TXInputs) == 1 && len(tx.TXInputs[0].TXid) == 0 && tx.TXInputs[0].Index == -1 {
		return true
	}

	return false
}

func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {
	utxos, resValue := bc.FindNeedUTXOs(from, amount) // 找到余额 utxos 切片index

	if resValue < amount {
		fmt.Println("余额不足")
		return nil
	}

	var inputs []TXInput
	var outputs []TXOutput

	// 循环可用的数据，生成转账input
	for id, indexArray := range utxos {
		for _, i := range indexArray {
			input := TXInput{[]byte(id), int64(i), from}
			inputs = append(inputs, input)
		}
	}

	// 生成转账数据
	output := TXOutput{amount, to}
	outputs = append(outputs, output)
	if resValue > amount {
		// 如果还剩，找零
		outputs = append(outputs, TXOutput{resValue - amount, from})
	}

	// 生成转账记录
	tx := Transaction{[]byte{}, inputs, outputs}
	tx.SetHash()

	return &tx
}
