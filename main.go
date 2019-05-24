package main

func main() {
	bc := NewBlockChain()

	bc.AddBlock("get fifty bitcoin")
	bc.AddBlock("get ten bitcoin")

	defer bc.db.Close()
	//for i, v := range bc.blocks {
	//	fmt.Printf("当前区块高度是:--------- %d\n", i)
	//	fmt.Printf("前一个区块的哈希值是: %x\n", v.PrevHash)
	//	fmt.Printf("当前区块的哈希值是: %x\n", v.Hash)
	//	fmt.Printf("区块数据是:%s\n", v.Data)
	//}
}
