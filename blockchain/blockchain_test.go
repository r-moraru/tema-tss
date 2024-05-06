package blockchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatesFirstBlock(t *testing.T) {
	blockchain := Blockchain{}

	block := blockchain.CreateBlock("data", "00")

	assert.Equal(t, "data", block.Data)
	assert.Equal(t, "", block.Previous_hash)
}

func TestCreatesSecondBlock(t *testing.T) {
	blockchain := Blockchain{}
	block1 := blockchain.CreateBlock("data1", "00")
	err := blockchain.AddBlock(block1)

	block2 := blockchain.CreateBlock("data2", "00")

	assert.NoError(t, err)
	assert.Equal(t, "data2", block2.Data)
	assert.Equal(t, block1.Hash, block2.Previous_hash)
}

func TestDoesNotAddBlockWithInvalidHash(t *testing.T) {
	blockchain := Blockchain{}
	block := blockchain.CreateBlock("data", "00")
	block.Hash = "wrong hash"

	err := blockchain.AddBlock(block)

	assert.ErrorIs(t, err, ErrBlockInvalidHash)
}

func TestDoesNotAddBlockWithInvalidPreviousHash(t *testing.T) {
	blockchain1 := Blockchain{}
	blockchain2 := Blockchain{}
	block1 := blockchain1.CreateBlock("data1", "00")
	blockchain1.AddBlock(block1)
	block2 := blockchain2.CreateBlock("data2", "00")
	blockchain2.AddBlock(block2)
	block3 := blockchain2.CreateBlock("data3", "00")

	err := blockchain1.AddBlock(block3)

	assert.ErrorIs(t, err, ErrPreviousBlockHashMismatch)
}

func TestCopyReturnsNewBlockchain(t *testing.T) {
	blockchain1 := Blockchain{}

	blockchain2 := blockchain1.Copy()
	block1 := blockchain2.CreateBlock("data1", "00")
	blockchain2.AddBlock(block1)

	_, err := blockchain1.GetLastBlock()
	assert.ErrorIs(t, err, ErrEmptyBlockchain)
	blockchain2LastBlock, err := blockchain2.GetLastBlock()
	assert.ErrorIs(t, err, nil)
	assert.Equal(t, block1, blockchain2LastBlock)
}

func TestCopyReturnsBlockchainWithSameBlocks(t *testing.T) {
	blockchain1 := Blockchain{}
	block1 := blockchain1.CreateBlock("data1", "00")
	blockchain1.AddBlock(block1)
	block2 := blockchain1.CreateBlock("data2", "00")
	blockchain1.AddBlock(block2)
	block3 := blockchain1.CreateBlock("data3", "00")
	blockchain1.AddBlock(block3)

	blockchain2 := blockchain1.Copy()

	blockchain2LastBlock, err := blockchain2.GetLastBlock()
	assert.ErrorIs(t, err, nil)
	assert.Equal(t, block3, blockchain2LastBlock)
}
