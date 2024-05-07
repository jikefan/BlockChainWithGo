package blockchain

import (
	"encoding/hex"
	"fmt"
	"goblockchain/transaction"
	"goblockchain/utils"
)

type BlockChain struct {
	Blocks []*Block
}

// This function adds a new block to the BlockChain
func (bc *BlockChain) AddBlock(txs []*transaction.Transaction) {
	// Get the previous block
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	// Create a new block with the previous block's hash and the given data
	newBlock := CreateBlock(prevBlock.Hash, txs)
	// Append the new block to the BlockChain
	bc.Blocks = append(bc.Blocks, newBlock)
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

func (bc *BlockChain) FindUnspentTransactions(address []byte) []transaction.Transaction {
	var unSpentTxs []transaction.Transaction
	spentTxs := make(map[string][]int)
	for idx := len(bc.Blocks) - 1; idx >= 0; idx-- {
		block := bc.Blocks[idx]
		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)
		IterOutputs:
			for outIdx, out := range tx.Outputs {
				if spentTxs[txID] != nil {
					for _, spentOut := range spentTxs[txID] {
						if spentOut == outIdx {
							continue IterOutputs
						}
					}
				}

				if out.ToAddressRight(address) {
					unSpentTxs = append(unSpentTxs, *tx)
				}
			}
			if !tx.IsBase() {
				for _, in := range tx.Inputs {
					if in.FromAddressRight(address) {
						inTxID := hex.EncodeToString(in.TxID)
						spentTxs[inTxID] = append(spentTxs[inTxID], in.OutIdx)
					}
				}
			}
		}
	}
	return unSpentTxs
}

func (bc *BlockChain) FindUTXOs(address []byte) (int, map[string]int) {
	unspentOuts := make(map[string]int)
	unspentTxs := bc.FindUnspentTransactions(address)
	accumulated := 0

Work:
	for _, tx := range unspentTxs {
		txID := hex.EncodeToString(tx.ID)
		for outIdx, out := range tx.Outputs {
			if out.ToAddressRight(address) {
				accumulated += out.Value
				unspentOuts[txID] = outIdx
				continue Work // one transaction can only have one output referred to adderss
			}
		}
	}
	return accumulated, unspentOuts
}

func (bc *BlockChain) FindSpendableOutputs(address []byte, amount int) (int, map[string]int) {
	unspentOuts := make(map[string]int)
	unspentTxs := bc.FindUnspentTransactions(address)
	accumulated := 0

Work:
	for _, tx := range unspentTxs {
		txID := hex.EncodeToString(tx.ID)
		for outIdx, out := range tx.Outputs {
			if out.ToAddressRight(address) && accumulated < amount {
				accumulated += out.Value
				unspentOuts[txID] = outIdx
				if accumulated >= amount {
					break Work
				}
				continue Work // one transaction can only have one output referred to adderss
			}
		}
	}
	return accumulated, unspentOuts
}

func (bc *BlockChain) CreateTransaction(from, to []byte, amount int) (*transaction.Transaction, bool) {
	var inputs []transaction.TxInput
	var outputs []transaction.TxOutput

	accumulated, validOutputs := bc.FindSpendableOutputs(from, amount)
	if accumulated < amount {
	    fmt.Println("Not enough coins!")
		return &transaction.Transaction{}, false
	}

	for txid, outidx := range validOutputs {
	    txID, err := hex.DecodeString(txid)
		utils.Handle(err)
		input := transaction.TxInput{TxID: txID, OutIdx: outidx, FromAddress: from}
		inputs = append(inputs, input)
	}
	outputs = append(outputs, transaction.TxOutput{Value: amount, ToAddress: to})
	if accumulated > amount {
	    outputs = append(outputs, transaction.TxOutput{Value: accumulated - amount, ToAddress: from})
	}
	tx := transaction.Transaction{Inputs: inputs, Outputs: outputs}

	return &tx, true
}

func (bc *BlockChain) Mine(txs []*transaction.Transaction) {
    bc.AddBlock(txs)
}