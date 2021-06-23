package main

import "github.com/HoYaStudy/Go_Study/hcoin/explorer"

// func main() {
// 	chain := blockchain.GetBlockchain()
// 	chain.AddBlock("Second Block")
// 	chain.AddBlock("Third Block")
// 	chain.AddBlock("Fourth Block")
// 	for _, block := range b.blocks {
// 		fmt.Printf("Data: %s\n", block.data)
// 		fmt.Printf("Hash: %s\n", block.hash)
// 		fmt.Printf("Prev Hash: %s\n\n", block.prevHash)
// 	}
// }
// }

func main() {
	explorer.Start()
}
