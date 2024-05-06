package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Handle error
func Handle(err error) {
	if err != nil {
	    log.Panic(err) // log error and exit program
	}
}

// ToHexInt converts an int64 to a byte array.
func ToHexInt(num int64) []byte {
	// Create a new buffer to store the byte array
	buff := new(bytes.Buffer)
	// Write the integer to the buffer in big endian format
	err := binary.Write(buff, binary.BigEndian, num)
	// If there is an error, log it and panic
	Handle(err)
	// Return the byte array
	return buff.Bytes()
}