package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
	"log"
)

type Wallet struct {
	Private *ecdsa.PrivateKey
	Public  []byte
}

// 创建钱包
func NewWallet() *Wallet {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panicln(err)
	}

	public := privateKey.PublicKey
	// 拼接横纵坐标
	pubkey := append(public.X.Bytes(), public.Y.Bytes()...)

	return &Wallet{
		Public:  pubkey,
		Private: privateKey,
	}
}

// 生成地址
func (w *Wallet) NewAddress() string {
	public := w.Public

	value := HashPubKey(public)
	version := byte(00)
	payload := append([]byte{version}, value...)

	checkCode := CheckSum(payload)
	payload = append(payload, checkCode...)

	// golang btcd 比特币全节点
	address := base58.Encode(payload)
	return address
}

func HashPubKey(data []byte) []byte {
	hash := sha256.Sum256(data)

	rip160hasher := ripemd160.New()
	_, err := rip160hasher.Write(hash[:])
	if err != nil {
		log.Println(err)
	}

	// 返回rip160的哈希结果
	return rip160hasher.Sum(nil)
}

func CheckSum(data []byte) []byte {
	// checksum
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])

	return hash2[:4]
}
