package main

import "github.com/HoYaStudy/Go_Study/hcoin/blockchain"

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")
	chain.DisplayBlocks()
}
