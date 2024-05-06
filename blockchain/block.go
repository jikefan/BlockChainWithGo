package blockchain

import (
	"bytes"
	"crypto/sha256"
	"goblockchain/utils"
	"time"
)

type Block struct {
	Timestamp int64
	Hash      []byte // Tts own hash value
	PrevHash  []byte // The hash value of the previous block
	Target    []byte // The target hash value
	Nonce     int64  // The nonce value used to find the target hash value
	Data      []byte
}

// This function sets the hash of a block
func (b *Block) SetHash() {
	// Join the prevhash, data, and timestamp into a single byte array
	information := bytes.Join([][]byte{b.PrevHash, b.Data, utils.ToHexInt(b.Timestamp), b.Target, utils.ToHexInt(b.Nonce)}, []byte{})
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
		Target: []byte{},
		Nonce: 0,
	}
	block.Target = block.GetTarget()
	block.Nonce = block.FindNonce()
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
