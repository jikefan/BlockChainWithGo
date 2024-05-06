// proofofwork.go
package blockchain

import (
	"bytes"
	"crypto/sha256"
	"goblockchain/constant"
	"goblockchain/utils"
	"math"
	"math/big"
)

// This function returns the target for a given block
func (b *Block) GetTarget() []byte {
	// Create a new big integer and set it to 1
	target := big.NewInt(1)
	// Left shift the big integer by 255 minus the constant difficulty
	target = target.Lsh(target, uint(255-constant.Difficulty))
	// Return the bytes of the target big integer
	return target.Bytes()
}

func (b *Block) GetBase4Nonce(nonce int64) []byte {
	data := bytes.Join([][]byte{
		b.PrevHash,
		b.Data,
		utils.ToHexInt(b.Timestamp),
		b.GetTarget(),
		utils.ToHexInt(nonce),
	}, []byte{})

	return data
}

// This function finds a nonce for a given block
func (b *Block) FindNonce() int64 {
	// Initialize nonce
	var nonce int64 = 0
	// Initialize big integers for target and hash
	var intTarget big.Int
	var intHash big.Int
	// Set the target from the block
	intTarget.SetBytes(b.GetTarget())

	// Loop until nonce is found that is smaller than target
	for nonce < math.MaxInt64 {
		// Get the nonce data
		data := b.GetBase4Nonce(nonce)
		// Hash the data
		hash := sha256.Sum256(data)
		// Set the hash
		intHash.SetBytes(hash[:])
		// If the hash is less than the target, break the loop
		if intHash.Cmp(&intTarget) == -1 {
			break
		} else {
			// therwise, increment the nonce
			nonce++
		}
	}
	// Return the nonce
	return nonce
}

// This function validates the Proof of Work (PoW) for a given Block
func (b *Block) ValidatePoW() bool {
	var intHash big.Int
	var intTarget big.Int
	intTarget.SetBytes(b.GetTarget())

	// Get the data for the nonce and create a hash
	data := b.GetBase4Nonce(b.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	// Return true if the hash is less than the target
	return intHash.Cmp(&intTarget) == -1
}
