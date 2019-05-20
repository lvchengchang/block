package main

import "fmt"

// todo 5-28添加区块addBlock
func main() {
	bc := NewBlockChain()

	bc.AddBlock("获取五十个比特币")
	bc.AddBlock("获取一百个比特币")

	for i, v := range bc.blocks {
		fmt.Printf("当前区块高度是:--------- %d\n", i)
		fmt.Printf("前一个区块的哈希值是: %x\n", v.PrevHash)
		fmt.Printf("当前区块的哈希值是: %x\n", v.Hash)
		fmt.Printf("区块数据是:%s\n", v.Data)
	}
}
