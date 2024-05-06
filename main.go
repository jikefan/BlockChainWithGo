package main

import (
	"goblockchain/blockchain"
	"log"
	// "time"
)

func main() {
	blockchain := blockchain.CreateBlockChain()
	// time.Sleep(time.Second)
	blockchain.AddBlockWithText("After genesis, I have something to say.")
	// time.Sleep(time.Second)
	blockchain.AddBlockWithText("jikefan is awesome!")
	// time.Sleep(time.Second)
	blockchain.AddBlockWithText("I can't wait to learn more about BlockChain!")
	// time.Sleep(time.Second)

	for _, block := range blockchain.Blocks {
		log.Printf("Timestamp: %d\n", block.Timestamp)
		log.Printf("Hash: %x\n", block.Hash)
		log.Printf("PrevHash: %x\n", block.PrevHash)
		log.Printf("Data: %s\n", block.Data)
		log.Printf("Nonce: %d\n", block.Nonce)
		log.Printf("Target: %x\n", block.Target)
		log.Printf("Proof of Work validation: %t\n", block.ValidatePoW())
		log.Println("----------------------------")
	}
}
