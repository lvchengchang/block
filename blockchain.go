package main

import (
	"fmt"
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
func NewBlockChain(address string) *BlockChain {
	var lastHash []byte

	db, err := bolt.Open(blockChainDb, 0600, nil)

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
			genesisBlock := GenesisBlock(address)

			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			bucket.Put([]byte("lastHashKey"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte("lastHashKey"))
		}

		return nil
	})

	return &BlockChain{db, lastHash}
}

func GenesisBlock(address string) *Block {
	coinBase := NewCoinBaseTx(address, "info")
	return NewBlock([]*Transaction{coinBase}, []byte{})
}

func (bc *BlockChain) AddBlock(txs []*Transaction) {
	//lastBlock := bc.blocks[len(bc.blocks)-1]
	//prevHash := lastBlock.Hash
	//
	//block := NewBlock(data, prevHash)
	//bc.blocks = append(bc.blocks, block)

	lastHash := bc.tail

	bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panicln("bucket not exist")
		}

		block := NewBlock(txs, lastHash)
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("lastHashKey"), block.Hash)

		bc.tail = block.Hash

		return nil
	})
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
		fmt.Printf("区块数据是:%s\n", block.Transactions[0].TXInputs[0].Sig)
		fmt.Println()
	}
}

func (bc *BlockChain) FindUTXOs(address string) []TXOutput {
	var UTXO []TXOutput
	// speed data
	spentOutputs := make(map[string][]int64)

	// range input find yourself speed utxo collection (查看是否已经消费了)
	it := bc.NewIterator()
	for {
		block := it.Next()
		if block.Hash == nil {
			// range complete
			break
		}

		for _, tx := range block.Transactions {
			fmt.Printf("current txid : %x\n", tx.TXId)
			// 遍历当前花费
		OUTPUT:
			for i, output := range tx.TXOutputs {
				fmt.Printf("current index %d", i)
				// 查看当前block的数据是否被消费
				// 如果交易id在切片内，代表与之前交易有关系
				if spentOutputs[string(tx.TXId)] != nil {
					for _, j := range spentOutputs[string(tx.TXId)] {
						if int64(i) == j {
							continue OUTPUT
						}
					}
				}

				if output.PukkeyHash == address {
					UTXO = append(UTXO, output)
				}
			}

			// 遍历当前获取 --- (挖矿交易不校验输入)
			if !tx.IsCoinBase() {
				continue
			}
			for _, input := range tx.TXInputs {
				if input.Sig == address {
					indexArray := spentOutputs[string(input.TXid)]
					indexArray = append(indexArray, input.index)
				}
			}
		}
	}
	// range output find yourself utxo

	return UTXO
}

func (bc *BlockChain) FindNeedUTXOs(from string, amount float64) (map[string][]uint64, float64) {
	utxos := make(map[string][]uint64)
	spentOutputs := make(map[string][]int64)
	var calc float64

	it := bc.NewIterator()
	for {
		block := it.Next()
		if block.Hash == nil {
			// range complete
			break
		}

		for _, tx := range block.Transactions {
			fmt.Printf("current txid : %x\n", tx.TXId)
			// 遍历当前花费
		OUTPUT:
			for i, output := range tx.TXOutputs {
				fmt.Printf("current index %d", i)
				// 查看当前block的数据是否被消费
				// 如果交易id在切片内，代表与之前交易有关系
				if spentOutputs[string(tx.TXId)] != nil {
					for _, j := range spentOutputs[string(tx.TXId)] {
						if int64(i) == j {
							continue OUTPUT
						}
					}
				}

				if calc < amount && output.PukkeyHash == from {
					//UTXO = append(UTXO, output)
					// todo find your need utxo
					utxos[string(tx.TXId)] = append(utxos[string(tx.TXId)], uint64(i))
					calc += output.Value
					if calc >= amount {
						return utxos, calc
					}
				}
			}

			// 遍历当前获取 --- (挖矿交易不校验输入)
			if !tx.IsCoinBase() {
				continue
			}
			for _, input := range tx.TXInputs {
				if input.Sig == from {
					indexArray := spentOutputs[string(input.TXid)]
					indexArray = append(indexArray, input.index)
				}
			}
		}
	}
	// range output find yourself utxo

	return utxos, calc
}
