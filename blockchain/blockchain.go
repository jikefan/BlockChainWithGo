package blockchain

type BlockChain struct {
	Blocks []*Block
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