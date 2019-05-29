package main

type Wallets struct {
	WalletsMap map[string]*Wallet
}

// 生成一个钱包
func NewWallets() *Wallets {
	wallet := NewWallet()
	address := wallet.NewAddress()

	var Wallets Wallets
	Wallets.WalletsMap = make(map[string]*Wallet)
	Wallets.WalletsMap[address] = wallet

	return &Wallets
}
