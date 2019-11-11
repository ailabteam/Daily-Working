package main

import (
	"fmt"
)

func main() {
	
	bc := NewBlockChain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
	
}

// 64 * 4 = 256 (hash256)
//6675f3dbd743d7887aa7992313ad025476158c75c439446ed870c3b746a9c5c1
