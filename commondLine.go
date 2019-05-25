package main

import "fmt"

func (cli *Cli) AddBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Printf("添加区块成功!\n")
}

func (cli *Cli) PrintBlockChain() {
	it := cli.bc.NewIterator()
	for {
		block := it.Next()
		if block.Hash == nil {
			return
		}
		fmt.Printf("版本号是: %d\n", block.Version)
		fmt.Printf("前一个区块的哈希值是: %x\n", block.PrevHash)
		fmt.Printf("梅克尔根的值是: %x\n", block.MarkelRoot)
		fmt.Printf("时间戳值是: %d\n", block.TimeStamp)
		fmt.Printf("难度值是: %d\n", block.Difficulty)
		fmt.Printf("当前区块的哈希值是: %x\n", block.Hash)
		fmt.Printf("区块数据是:%s\n", block.Data)
		fmt.Println()
	}
}
