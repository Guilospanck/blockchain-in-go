package main

import "fmt"

func main() {

	newBlockchain := NewBlockchain() // initializes the Blockchain

	// Simulating some mining
	newBlockchain.AddBlock("Jose TX 1 BTC -> Maria RX 1 BTC")
	newBlockchain.AddBlock("Zeus TX 100 BTC -> Guilherme RX 100 BTC")

	for idx, block := range newBlockchain.Blocks {
		fmt.Printf("Block ID: %d \n", idx)
		fmt.Printf("Timestamp: %d \n", block.Timestamp)
		fmt.Printf("Hash of the Block: %x \n", block.MyBlockHash)
		fmt.Printf("Hash of the previous Block: %x \n", block.PreviousBlockHash)
		fmt.Printf("All the transactions: %s \n", block.AllData)
		fmt.Println("=======================================")
	}

}
