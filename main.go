package main

func main() {
	bc := NewBlockChain("班长")
	cli := Cli{bc}
	cli.Send("班长", "李四", 10, "班长", "张三转给李四10")

	//if err := os.RemoveAll("blockChain.db"); err != nil {
	//	log.Fatalln(err)
	//}
	//bc := NewBlockChain()
	//
	//bc.AddBlock("get fifty bitcoin")
	//bc.AddBlock("get ten bitcoin")
	//
	//defer bc.db.Close()

	//it := bc.NewIterator()
	//for {
	//	block := it.Next()
	//	if block.Hash == nil {
	//		return
	//	}
	//
	//	fmt.Printf("前一个区块的哈希值是: %x\n", block.PrevHash)
	//	fmt.Printf("当前区块的哈希值是: %x\n", block.Hash)
	//	fmt.Printf("区块数据是:%s\n", block.Data)
	//	fmt.Println()
	//}
}
