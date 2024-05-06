package blockchain

import (
	"fmt"

	"github.com/r-moraru/tema-TSS/block"
)

var ErrBlockInvalidHash = fmt.Errorf("block has invalid hash")
var ErrEmptyBlockchain = fmt.Errorf("blockchain is empty")
var ErrPreviousBlockHashMismatch = fmt.Errorf("new block's previous hash doesn't match last block's hash")

type Blockchain struct {
	Blockchain []block.Block
}

func (b *Blockchain) AddBlock(block block.Block) error {
	if !block.CheckHash() {
		return ErrBlockInvalidHash
	}
	lastBlock, err := b.GetLastBlock()
	if err == nil && block.Previous_hash != lastBlock.Hash {
		return ErrPreviousBlockHashMismatch
	}
	b.Blockchain = append(b.Blockchain, block)
	return nil
}

func (b *Blockchain) GetLength() int {
	return len(b.Blockchain)
}

func (b *Blockchain) GetLastBlock() (block.Block, error) {
	blockchainLength := b.GetLength()
	if blockchainLength == 0 {
		return block.Block{}, ErrEmptyBlockchain
	}
	return b.Blockchain[blockchainLength-1], nil
}

func (b *Blockchain) CreateBlock(data, difficulty string) block.Block {
	lastBlock, err := b.GetLastBlock()
	if err != nil {
		return block.NewBlock(data, "", difficulty)
	}
	return block.NewBlock(data, lastBlock.Hash, difficulty)
}

// TODO: test copy functionality
func (b *Blockchain) Copy() Blockchain {
	newBlocks := make([]block.Block, len(b.Blockchain))
	copy(newBlocks, b.Blockchain)
	return Blockchain{Blockchain: newBlocks}
}
