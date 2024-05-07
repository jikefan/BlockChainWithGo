package transaction

import "bytes"

type TxOutput struct {
    Value int // The value of transferred assets
	ToAddress []byte // The address of the recipient
}

type TxInput struct {
    TxID []byte // Preceding transaction information supporting this transaction
	OutIdx int // The index of the output in the preceding transaction
	FromAddress []byte // The address of the sender
}

func (in *TxInput) FromAddressRight(address []byte) bool {
    return bytes.Equal(in.FromAddress, address)
}

func (out *TxOutput) ToAddressRight(address []byte) bool {
    return bytes.Equal(out.ToAddress, address)
}