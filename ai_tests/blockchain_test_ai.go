package ai_tests

import (
	"testing"

	"github.com/r-moraru/tema-TSS/block"
	"github.com/r-moraru/tema-TSS/blockchain"
)

func TestAddValidBlock(t *testing.T) {
	// Create a new blockchain
	bc := blockchain.Blockchain{}

	// Add a valid block
	validBlock := block.NewBlock("Test Data", "", "00000")
	err := bc.AddBlock(validBlock)

	// Check if there are any errors
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Check if the length of the blockchain increased by one
	if bc.GetLength() != 1 {
		t.Errorf("Expected blockchain length to be 1, got: %d", bc.GetLength())
	}
}

func TestAddInvalidBlock(t *testing.T) {
	// Create a new blockchain
	bc := blockchain.Blockchain{}

	// Add an invalid block (invalid hash)
	invalidBlock := block.Block{
		Data:          "Invalid Data",
		Previous_hash: "",
		Timestamp:     0,
		Nonce:         0,
		Hash:          "Invalid Hash",
	}
	err := bc.AddBlock(invalidBlock)

	// Check if the error matches the expected error
	if err != blockchain.ErrBlockInvalidHash {
		t.Errorf("Expected ErrBlockInvalidHash, got: %v", err)
	}

	// Check if the blockchain length remains 0
	if bc.GetLength() != 0 {
		t.Errorf("Expected blockchain length to be 0, got: %d", bc.GetLength())
	}
}

func TestGetLastBlock(t *testing.T) {
	// Create a new blockchain
	bc := blockchain.Blockchain{}

	// Get last block when blockchain is empty
	_, err := bc.GetLastBlock()

	// Check if the error matches the expected error
	if err != blockchain.ErrBlockInvalidHash {
		t.Errorf("Expected ErrEmptyBlockchain, got: %v", err)
	}

	// Add a valid block
	validBlock := block.NewBlock("Test Data", "", "00000")
	bc.AddBlock(validBlock)

	// Get last block after adding a block
	lastBlock, err := bc.GetLastBlock()

	// Check if there are no errors
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Check if the last block matches the added block
	if lastBlock != validBlock {
		t.Errorf("Last block does not match the added block")
	}
}

func TestCopyBlockchain(t *testing.T) {
	// Create a new blockchain
	bc := blockchain.Blockchain{}

	// Add a valid block
	validBlock := block.NewBlock("Test Data", "", "00000")
	bc.AddBlock(validBlock)

	// Create a copy of the blockchain
	copy := bc.Copy()

	// Check if the copied blockchain is equal to the original
	if len(copy.Blockchain) != len(bc.Blockchain) {
		t.Errorf("Copied blockchain length does not match original")
	}

	// Modify the copy and check if it's independent of the original
	copy.AddBlock(block.NewBlock("New Data", "", "00000"))
	if len(copy.Blockchain) == len(bc.Blockchain) {
		t.Errorf("Modifying copy also modifies the original blockchain")
	}
}
