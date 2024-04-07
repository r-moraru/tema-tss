package block

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatesEmptyFirstBlockWithDifficultyZero(t *testing.T) {
	block := NewBlock("", "", "")

	assert.Equal(t, int64(0), block.Nonce)
	assert.Equal(t, 64, len(block.Hash))
}

func TestCreatesEmptyFirstBlockWithDifficultyFour(t *testing.T) {
	difficulty := "0000"

	block := NewBlock("", "", difficulty)

	assert.Equal(t, 64, len(block.Hash))
	assert.Equal(t, block.Hash[:4], difficulty)
}

func TestCreatesLinkedBlockWithDifficultyZero(t *testing.T) {
	data := "test data"
	previous_hash := "test hash"

	block := NewBlock(data, previous_hash, "")

	assert.Equal(t, 64, len(block.Hash))
	assert.Equal(t, data, block.Data)
	assert.Equal(t, previous_hash, block.Previous_hash)
}

func TestCreatesLinkedBlockWithDifficutlyFour(t *testing.T) {
	data := "test data"
	previous_hash := "test hash"
	difficulty := "0000"

	block := NewBlock(data, previous_hash, difficulty)

	assert.Equal(t, 64, len(block.Hash))
	assert.Equal(t, difficulty, block.Hash[:4])
	assert.Equal(t, data, block.Data)
	assert.Equal(t, previous_hash, block.Previous_hash)
}

func TestCheckHashReturnsTrueForGoodBlockWithDifficultyFour(t *testing.T) {
	data := "test data"
	previous_hash := "test hash"
	difficulty := "0000"

	block := NewBlock(data, previous_hash, difficulty)

	assert.True(t, block.CheckHash())
}

func TestCheckHashReturnsFalseForBlockWithDifficultyFourAndWrongHash(t *testing.T) {
	data := "test data"
	previous_hash := "test hash"
	difficulty := "0000"

	block := NewBlock(data, previous_hash, difficulty)
	block.Hash = "wrong hash"

	assert.False(t, block.CheckHash())
}
