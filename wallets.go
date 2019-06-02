package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Wallets struct {
	WalletsMap map[string]*Wallet
}

// 生成一个钱包
func NewWallets() *Wallets {
	var ws Wallets
	ws.WalletsMap = make(map[string]*Wallet)

	return &ws
}

func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := wallet.NewAddress()

	ws.WalletsMap[address] = wallet

	ws.saveToFile()
	return address
}

func (ws *Wallets) saveToFile() {
	var buffer bytes.Buffer
	// 编码
	encoder := gob.NewEncoder(&buffer)
	gob.Register(elliptic.P256())
	err := encoder.Encode(ws)
	// golang interface 不允许编码
	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile("wallet.data", buffer.Bytes(), 0600)
}
