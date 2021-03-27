package main

import (
	"crypto/sha256"
	"reflect"
	"testing"
)

// TestNewBlock verifies NewBlock creates a block with the
// supplied prevHash
func TestNewBlock(t *testing.T) {
	testCases := []struct {
		testHash [32]byte
	}{
		{testHash: sha256.Sum256([]byte("testhash"))},
		{testHash: [32]byte{}},
	}

	for _, testCase := range testCases {
		block := NewBlock(testCase.testHash)

		if block == nil {
			t.Error("Expected address of created block")
		}

		if !reflect.DeepEqual(block.header.prevHash, testCase.testHash) {
			t.Errorf("Expected block prev Hash with value %v", block.header.prevHash)
		}
	}
}

func TestAcceptTransactionPanic(t *testing.T) {
	testHash := sha256.Sum256([]byte("testhash"))
	block := NewBlock(testHash)

	block.AcceptTransaction(transaction{})
}

// TestAcceptTransaction verifies the case of adding a new transaction
// by checking the new and old merkleroots do not equal, and verifying
// new transaction length = old transaction length + 1
func TestAcceptTransaction(t *testing.T) {

	testHash := sha256.Sum256([]byte("testhash"))
	block := NewBlock(testHash)

	block.AcceptTransaction(transaction{value: 1})

	initialMerkleRoot := block.header.merkleRoot
	initialTxnsLength := len(block.transactions)

	block.AcceptTransaction(transaction{value: 2})
	secondMerkleRoot := block.header.merkleRoot
	secondTxnsLength := len(block.transactions)

	if !reflect.DeepEqual(initialMerkleRoot, secondMerkleRoot) {
		t.Errorf("Expected merkle root of the initial two transactions to be equal")
	}

	if secondTxnsLength != initialTxnsLength+1 {
		t.Errorf("Expected transaction length to be 1")
	}

	block.AcceptTransaction(transaction{value: 3})

	finalMerkleRoot := block.header.merkleRoot
	finalTxnsLength := len(block.transactions)

	if reflect.DeepEqual(secondMerkleRoot, finalMerkleRoot) {
		t.Errorf("Expected new transaction to generate new merkle root")
	}

	if finalTxnsLength != secondTxnsLength+1 {
		t.Errorf("Expected transaction length to be 1")
	}
}

func TestGenerateMerkleRootPass(t *testing.T) {
	testHash := sha256.Sum256([]byte("testhash"))
	block := NewBlock(testHash)

	block.transactions = append(block.transactions, transaction{})

	block.generateNewMerkleRoot()
}

func TestGenerateMerkleRootPanic(t *testing.T) {
	defer checkForPanic(t)

	testHash := sha256.Sum256([]byte("testhash"))
	block := NewBlock(testHash)

	block.generateNewMerkleRoot()
}

func checkForPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("The code did not panic")
	}
}

func TestCalculateHash(t *testing.T) {
	testHash := sha256.Sum256([]byte("testhash"))
	block := NewBlock(testHash)

	hash := block.CalculateHash()

	if len(hash) == 0 {
		t.Errorf("Expected block header hash")
	}
}
