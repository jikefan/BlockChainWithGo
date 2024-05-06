package transaction

type TxOutput struct {
    Value int // The value of transferred assets
	ToAddress []byte // The address of the recipient
}

type TxInput struct {
    TxID []byte // Preceding transaction information supporting this transaction
	OutIdx int // The index of the output in the preceding transaction
	FromAddress []byte // The address of the sender
}