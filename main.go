package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Data      []byte
	PrevHash  []byte
	BlockHash []byte
}

func NewBlock(data, prevHash []byte) *Block {
	b := new(Block)
	b.Data = data
	b.PrevHash = prevHash

	h := sha512.New512_256()
	h.Write(b.Data)
	h.Write(b.PrevHash)

	b.BlockHash = h.Sum(nil)

	return b
}

func main() {
	genesis := NewBlock([]byte("This is yet more data"), []byte{0})

	newBlock := NewBlock([]byte("This is some more data"), genesis.BlockHash)

	fmt.Printf("Genesis: %s\n", hex.EncodeToString(genesis.BlockHash))
	fmt.Printf("New Blk: %s\n", hex.EncodeToString(newBlock.BlockHash))
}
