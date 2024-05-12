package block

import (
	"testing"
	"strings"
)

func TestNewBlock(t *testing.T) {
	data := "Test Data"
	previousHash := "Previous Hash"
	difficulty := "00000" // Example difficulty

	block := NewBlock(data, previousHash, difficulty)

	// Check if the block's data is set correctly
	if block.Data != data {
		t.Errorf("Expected data %s, got %s", data, block.Data)
	}

	// Check if the block's previous hash is set correctly
	if block.Previous_hash != previousHash {
		t.Errorf("Expected previous hash %s, got %s", previousHash, block.Previous_hash)
	}

	// Check if the block's timestamp is set
	if block.Timestamp <= 0 {
		t.Errorf("Timestamp not set correctly")
	}
}

func TestCalculateHash(t *testing.T) {
	data := "Test Data"
	previousHash := "Previous Hash"
	difficulty := "00000" // Example difficulty

	block := NewBlock(data, previousHash, difficulty)

	block.calculateHash(difficulty)

	// Check if the block's hash is calculated correctly
	if !strings.HasPrefix(block.Hash, difficulty) {
		t.Errorf("Block hash not calculated correctly")
	}
}

func TestCheckHash(t *testing.T) {
	data := "Test Data"
	previousHash := "Previous Hash"
	difficulty := "00000" // Example difficulty

	block := NewBlock(data, previousHash, difficulty)

	// Store the original hash
	originalHash := block.Hash

	// Reset the hash and check if it matches after re-calculating
	block.Hash = ""
	valid := block.CheckHash()

	if !valid {
		t.Errorf("Block hash validation failed")
	}

	// Restore the original hash
	block.Hash = originalHash
}
