package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
)

func main() {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panicln(err)
	}

	// 生成公钥
	pubKey := privateKey.PublicKey

	data := "hello,world!"
	hash := sha256.Sum256([]byte(data))

	// 签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("pubkey : %v\n", pubKey)
	fmt.Printf("r : %v\n", r.Bytes())
	fmt.Printf("s : %v\n", s.Bytes())

	signature := append(r.Bytes(), s.Bytes()...)

	r1 := big.Int{}
	s1 := big.Int{}

	r1.SetBytes(signature[:len(signature)/2])
	s1.SetBytes(signature[len(signature)/2:])

	// 校验 需要三要素 : 数据 签名 公钥
	res := ecdsa.Verify(&pubKey, hash[:], r, s)
	fmt.Println(res)
}
