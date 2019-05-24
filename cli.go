package main

import (
	"fmt"
	"os"
)

type Cli struct {
	ec *BlockChain
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

	switch args[1] {
	case "addBlock":
		break
	case "printChain":
		break
	default:
		fmt.Printf(USAGE)
	}
}
