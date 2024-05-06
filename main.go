package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

type Block struct {
	Timestamp int64
	Hash      []byte // Tts own hash value
	PrevHash  []byte // The hash value of the previous block
	Data      []byte
}

type BlockChain struct {
	Blocks []*Block
}

// ToHexInt converts an int64 to a byte array.
func ToHexInt(num int64) []byte {
	// Create a new buffer to store the byte array
	buff := new(bytes.Buffer)
	// Write the integer to the buffer in big endian format
	err := binary.Write(buff, binary.BigEndian, num)
	// If there is an error, log it and panic
	if err != nil {
		log.Panic(err)
	}
	// Return the byte array
	return buff.Bytes()
}

// This function sets the hash of a block
func (b *Block) SetHash() {
	// Join the prevhash, data, and timestamp into a single byte array
	information := bytes.Join([][]byte{b.PrevHash, b.Data, ToHexInt(b.Timestamp)}, []byte{})
	// Calculate the hash of the information
	hash := sha256.Sum256(information)
	// Set the hash of the block
	b.Hash = hash[:]
}

// This function creates a new block with the given prevhash and data
func CreateBlock(prevhash, data []byte) *Block {
	// Create a new block structure
	block := &Block{
		// Set the timestamp to the current time
		Timestamp: time.Now().Unix(),
		// Set the hash to an empty byte array
		Hash: []byte{},
		// Set the prevhash to the given prevhash
		PrevHash: prevhash,
		// Set the data to the given data
		Data: data,
	}
	// Set the hash of the block
	block.SetHash()
	// Return the block
	return block
}

// Function to create the genesis block
func GenesisBlock() *Block {
	// Array of bytes containing the message to be included in the genesis block
	genesisWords := []byte("Hello,Blockchain!")
	// Return the created block with the given message
	return CreateBlock([]byte{}, genesisWords)
}

// This function adds a new block to the BlockChain
func (bc *BlockChain) AddBlock(data []byte) {
	// Get the previous block
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	// Create a new block with the previous block's hash and the given data
	newBlock := CreateBlock(prevBlock.Hash, data)
	// Append the new block to the BlockChain
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *BlockChain) AddBlockWithText(data string) {
	// Convert the data to bytes
	dataBytes := []byte(data)
	// Add the block to the blockchain
	bc.AddBlock(dataBytes)
}

// This function creates a new BlockChain
func CreateBlockChain() *BlockChain {
	// Create a new BlockChain
	blockchain := &BlockChain{}
	// Add the genesis block to the BlockChain
	blockchain.Blocks = append(blockchain.Blocks, GenesisBlock())
	// Return the BlockChain
	return blockchain
}

func main() {
	blockchain := CreateBlockChain()
	time.Sleep(time.Second)
	blockchain.AddBlockWithText("After genesis, I have something to say.")
	time.Sleep(time.Second)
	blockchain.AddBlockWithText("jikefan is awesome!")
	time.Sleep(time.Second)
	blockchain.AddBlockWithText("I can't wait to learn more about BlockChain!")
	time.Sleep(time.Second)

	for _, block := range blockchain.Blocks {
		log.Printf("Timestamp: %d\n", block.Timestamp)
		log.Printf("Hash: %x\n", block.Hash)
		log.Printf("PrevHash: %x\n", block.PrevHash)
		log.Printf("Data: %s\n", block.Data)
		log.Println("----------------------------")
	}
}
