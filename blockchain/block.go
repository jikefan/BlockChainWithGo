package blockchain

import (
	"bytes"
	"crypto/sha256"
	"goblockchain/transaction"
	"goblockchain/utils"
	"time"
)

type Block struct {
	Timestamp    int64
	Hash         []byte // Tts own hash value
	PrevHash     []byte // The hash value of the previous block
	Target       []byte // The target hash value
	Nonce        int64  // The nonce value used to find the target hash value
	Transactions []*transaction.Transaction
}

// This function sets the hash of a block
func (b *Block) SetHash() {
	// Join the prevhash, data, and timestamp into a single byte array
	information := bytes.Join([][]byte{b.PrevHash, b.BackTransactionSummary(), utils.ToHexInt(b.Timestamp), b.Target, utils.ToHexInt(b.Nonce)}, []byte{})
	// Calculate the hash of the information
	hash := sha256.Sum256(information)
	// Set the hash of the block
	b.Hash = hash[:]
}

// This function creates a new block with the given prevhash and data
func CreateBlock(prevhash []byte, txs []*transaction.Transaction) *Block {
	// Create a new block structure
	block := &Block{
		// Set the timestamp to the current time
		Timestamp: time.Now().Unix(),
		// Set the hash to an empty byte array
		Hash: []byte{},
		// Set the prevhash to the given prevhash
		PrevHash:     prevhash,
		Transactions: txs,
		Target:       []byte{},
		Nonce:        0,
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
	tx := transaction.BaseTx([]byte("jikefan"))
	// Return the created block with the given message
	return CreateBlock([]byte{}, []*transaction.Transaction{tx})
}

func (b *Block) BackTransactionSummary() []byte {
	txIDs := make([][]byte, 0)
	for _, tx := range b.Transactions {
		txIDs = append(txIDs, tx.ID)
	}

	summary := bytes.Join(txIDs, []byte{})
	return summary
}