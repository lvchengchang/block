package main

import (
	"github.com/boltdb/bolt"
	"log"
)

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"

type BlockChain struct {
	db *bolt.DB

	tail []byte // last block hash
}

// create blockchain
func NewBlockChain() *BlockChain {
	genesisBlock := GenesisBlock()
	var lastHash []byte

	db, err := bolt.Open(blockChainDb, 0600, nil)
	defer db.Close()

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Fatalln(err)
			}

			bucket.Put(genesisBlock.Hash, genesisBlock.toByte())
			bucket.Put([]byte("lastHashKey"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte("lastHashKey"))
		}

		return nil
	})

	return &BlockChain{db, lastHash}
}

func GenesisBlock() *Block {
	return NewBlock("this is first block", []byte{})
}

func (bc *BlockChain) AddBlock(data string) {
	//lastBlock := bc.blocks[len(bc.blocks)-1]
	//prevHash := lastBlock.Hash
	//
	//block := NewBlock(data, prevHash)
	//bc.blocks = append(bc.blocks, block)
}
