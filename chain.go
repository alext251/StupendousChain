package main

import (
	"encoding/binary"
	"fmt"
	"time"
)

const timestampInterval = 2016
const threshold = 1000000000

type blockchain struct {
	currentBlock *Block
	length       int
}

// StartBlockchain is the driver function for creating a new
// blockchain instance and starting the block validation
// process.
func StartBlockchain() {
	bc := newBlockchain()

	for {
		bc.attemptBlockValidation()
		fmt.Printf("%v\n", bc)
		time.Sleep(5 * time.Second)
	}
}

func newBlockchain() *blockchain {
	startingBlock := createStartingBlock()

	return &blockchain{
		currentBlock: startingBlock,
		length:       0,
	}
}

func createStartingBlock() *Block {
	latestHash := getLatestBlockHash()
	return NewBlock(latestHash)
}

// Get the hash of the latest validated block on the network
func getLatestBlockHash() [32]byte {
	return [32]byte{}
}

// Helper that queries network for latest validated block
// in order to get its hash
func getLatestBlock() *Block {
	return nil
}

func (bc *blockchain) attemptBlockValidation() {
	// get current block hash
	hash := bc.currentBlock.CalculateHash()

	if isValidHash(hash) {
		newBlock := NewBlock(hash)
		bc.currentBlock = newBlock
		bc.length++
	} else {
		bc.currentBlock.header.nonce++
	}
}

func isValidHash(hash [32]byte) bool {
	intVal := binary.BigEndian.Uint32(hash[:])
	return intVal < uint32(threshold)
}
