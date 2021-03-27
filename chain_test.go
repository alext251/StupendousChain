package main

import (
	"crypto/sha256"
	"testing"
)

func TestNewBlockchain(t *testing.T) {
	if chain := newBlockchain(); chain == nil {
		t.Error("Expected a blockchain to be created; got nil instead")
	}
}

func TestCreateStartingBlock(t *testing.T) {
	if latestBlock := createStartingBlock(); latestBlock == nil {
		t.Error("Expected a block to be created; got nil instead")
	}
}

func TestAttemptBlockValidation(t *testing.T) {

	// Controlling hash by inserting arbitrary timestamps known
	// to produce a value below and above hash, respectively
	testCases := []struct {
		timestamp      int64
		nonce          int32
		expectedLength int
	}{
		{8, 888, 0}, // Brute-force tested to get value higher than threshold
		{8, 0, 1},   // Brute-force tested to get value lower than threshold
	}

	for _, testCase := range testCases {
		testChain := newBlockchain()
		testChain.currentBlock.header.timestamp = testCase.timestamp
		testChain.currentBlock.header.nonce = testCase.nonce
		testChain.attemptBlockValidation()

		if testChain.length != testCase.expectedLength {
			t.Errorf("Block validation case failed for %v", testCase)
		}
	}
}

func TestIsHashValid(t *testing.T) {
	testCases := []struct {
		data          []byte
		shouldBeValid bool
	}{
		{data: []byte("X"), shouldBeValid: false}, // arbitrary case wst current threshold - verified to generate value above threshold
		{data: []byte("x"), shouldBeValid: true},  // arbitrary case wst current threshold - verified to generate value below threshold
	}

	for _, testCase := range testCases {
		hash := sha256.Sum256(testCase.data)
		isValid := isValidHash(hash)
		if isValid && !testCase.shouldBeValid {
			t.Error("Got valid hash; expected invalid hash")
		}

		if !isValid && testCase.shouldBeValid {
			t.Error("Got invalid hash; expected valid hash")
		}
	}
}

// func generateHash(data []byte) [32]byte {
// 	hash :=
// 	intVal := binary.BigEndian.Uint32(hash[:])
// 	fmt.Printf("%d", intVal)
// }
