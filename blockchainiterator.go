package main

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChainIterator struct {
	db                 *bolt.DB
	currentHashPointer []byte
}

func (bc *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{
		bc.db,
		bc.tail,
	}
}

// iterator belong blockChain
func (it *BlockChainIterator) Next() *Block {
	var Block Block

	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Fatalln("not existing")
		}

		BlockByte := bucket.Get(it.currentHashPointer)
		Block = Deserialize(BlockByte)

		it.currentHashPointer = Block.PrevHash

		return nil
	})

	return &Block
}
