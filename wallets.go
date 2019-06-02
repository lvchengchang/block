package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const WALLETFILENAME = "wallet.dat"

type Wallets struct {
	WalletsMap map[string]*Wallet
}

// 生成一个钱包
func NewWallets() *Wallets {
	var ws Wallets
	ws.WalletsMap = make(map[string]*Wallet)

	ws.loadFile()
	return &ws
}

func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := wallet.NewAddress()

	ws.WalletsMap[address] = wallet

	ws.saveToFile()
	fmt.Printf("local address %s\n", address)
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

	ioutil.WriteFile(WALLETFILENAME, buffer.Bytes(), 0600)
}

func (ws *Wallets) loadFile() {
	_, err := os.Stat(WALLETFILENAME)
	if os.IsNotExist(err) {
		ws1 := Wallets{}
		ws1.WalletsMap = make(map[string]*Wallet)
		return
	}

	content, err := ioutil.ReadFile(WALLETFILENAME)
	if err != nil {
		log.Panicln(err)
	}

	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(content))
	var wsLocal Wallets
	err = decoder.Decode(&wsLocal)
	if err != nil {
		log.Panicln(err)
	}

	ws.WalletsMap = wsLocal.WalletsMap
}

func (ws *Wallets) GetAllAddress() []string {
	var addresses []string
	for address := range ws.WalletsMap {
		addresses = append(addresses, address)
	}

	return addresses
}
