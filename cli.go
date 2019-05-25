package main

import (
	"fmt"
	"os"
)

type Cli struct {
	bc *BlockChain
}

const USAGE = `
	addBlock --data DATA "add data to blockchain"
	printChain           "print all blockchain data"
`

func (cli *Cli) Run() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf(USAGE)
		return
	}

	cmd := args[1]
	switch cmd {
	case "addBlock":
		if len(args) == 4 && args[2] == "--data" {
			data := args[3]
			cli.AddBlock(data)
		} else {
			fmt.Println("添加区块参数失败,请重试")
		}
	case "printChain":
		cli.PrintBlockChain()
	default:
		fmt.Printf(USAGE)
	}
}
