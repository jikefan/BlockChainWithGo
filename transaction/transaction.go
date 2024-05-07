package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"goblockchain/constant"
	"goblockchain/utils"
)

type Transaction struct {
    ID []byte // Its own hash value
	Inputs []TxInput
	Outputs []TxOutput
}

// This function returns the transaction hash of a given transaction
func (tx *Transaction) TxHash() []byte {
    // Create a buffer to store the encoded transaction
	var encoded bytes.Buffer
	// Create a hash to store the transaction hash
	var hash [32]byte

	// Create a new encoder to encode the transaction
	encoder := gob.NewEncoder(&encoded)
	// Encode the transaction
	err := encoder.Encode(tx)
	// Handle any errors
	utils.Handle(err)

	// Calculate the transaction hash using the encoded transaction
	hash = sha256.Sum256(encoded.Bytes())

	// Return the transaction hash
	return hash[:]
}

func (tx *Transaction) SetID() {
    tx.ID = tx.TxHash()
}

// This function creates a base transaction
func BaseTx(toaddress []byte) *Transaction {
    // Create a transaction input with an empty byte array for the prevout hash and index, and an empty byte array for the signature
    txIn := TxInput{[]byte{}, -1, []byte{}}
    // Create a transaction output with the specified 'toaddress' and initial coin
    txOut := TxOutput{constant.InitCoin, toaddress}
    // Create a transaction with a note and the transaction input, output
    tx := Transaction{[]byte("This is the Base Transaction!"), []TxInput{txIn}, []TxOutput{txOut}}
    // Return the transaction
    return &tx
}

func (tx *Transaction) IsBase() bool {
	return len(tx.Inputs) == 1 && tx.Inputs[0].OutIdx == -1
}
