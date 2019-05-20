package main

// create blockchain
func NewBlockChain() *BlockChain {
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

func GenesisBlock() *Block {
	return NewBlock("this is first block", []byte{})
}

func (bc *BlockChain) AddBlock(data string) {
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash

	block := NewBlock(data, prevHash)
	bc.blocks = append(bc.blocks, block)
}
