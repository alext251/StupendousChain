package main

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/cbergoon/merkletree"
)

type merkleNode = merkletree.Content

// BlockHeader is a simple way to hold all block header content
//
// More help with understanding block headers:
// 	https://www.investopedia.com/terms/b/block-header-cryptocurrency.asp
// More help with understanding hashing
type BlockHeader struct {
	timestamp  int64 //TODO: find way to convert to int32
	version    int32
	nonce      int32
	difficulty int32
	prevHash   [32]byte
	merkleRoot [32]byte
}

// Block contains a header and transaction section to

type Block struct {
	transactions []merkleNode
	header       BlockHeader
}

func NewBlock(prevHash [32]byte) *Block {
	return &Block{
		transactions: []merkleNode{},
		header: BlockHeader{
			timestamp:  time.Now().UnixNano(),
			prevHash:   prevHash,
			merkleRoot: [32]byte{},
			nonce:      0,
		},
	}
}

func (b *Block) CalculateHash() [32]byte {
	headerData := []byte(fmt.Sprintf("%v", b.header))
	return sha256.Sum256(headerData)
}

func (b *Block) AcceptTransaction(txn merkleNode) {
	b.transactions = append(b.transactions, txn)
	b.generateNewMerkleRoot()
}

func (b *Block) generateNewMerkleRoot() {
	var root [32]byte

	tree, err := merkletree.NewTree(b.transactions)
	if err != nil {
		panic(err)
	}

	rawRoot := tree.MerkleRoot()
	copy(root[:], rawRoot[:32])

	b.header.merkleRoot = root
}

/* MOCK TRANSACTION */
type transaction struct {
	value int
}

// Needed for merkletree.Content
func (t transaction) CalculateHash() ([]byte, error) {
	return nil, nil
}

// Needed for merkletree.Content
func (t transaction) Equals(other merkleNode) (bool, error) {
	return false, nil
}
