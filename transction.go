package main

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
	Sig string
}

type TXOutput struct {
	value float64
	// 锁定脚本
	PukkeyHash string
}
