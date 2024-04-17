package node

import (
	"testing"

	"github.com/r-moraru/tema-TSS/block"
	"github.com/r-moraru/tema-TSS/blockchain"
	"github.com/r-moraru/tema-TSS/network/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewNodeCreatedWhenBlockchainIsEmpty(t *testing.T) {
	blockchainChan := make(chan blockchain.Blockchain, 1)
	blockchainChan <- blockchain.Blockchain{}
	network := mocks.NewNetwork(t)
	network.EXPECT().GetBlockchain().Once().Return(blockchainChan)

	node := NewNode(network, "0000")

	assert.Equal(t, node.GetBlockchainLength(), 0)
	assert.Equal(t, node.GetLastBlockHash(), "")
}

func TestNewNodeCreatedWhenBlockchainIsPopulated(t *testing.T) {
	existingBlockchain := blockchain.Blockchain{}
	block1 := block.NewBlock("block1", "", "00")
	block2 := block.NewBlock("block2", block1.Hash, "00")
	block3 := block.NewBlock("block3", block2.Hash, "00")
	existingBlockchain.AddBlock(block1)
	existingBlockchain.AddBlock(block2)
	existingBlockchain.AddBlock(block3)
	blockchainChan := make(chan blockchain.Blockchain, 1)
	blockchainChan <- existingBlockchain
	network := mocks.NewNetwork(t)
	network.EXPECT().GetBlockchain().Once().Return(blockchainChan)

	node := NewNode(network, "00")

	assert.Equal(t, node.GetBlockchainLength(), 3)
	assert.NotEqual(t, node.GetLastBlock(), nil)
	assert.Equal(t, *node.GetLastBlock(), block3)
}

func TestCanAddFirstBlockFromNetwork(t *testing.T) {
	expectedBlock := block.NewBlock("block1", "", "00")
	blockChan := make(chan block.Block, 1)
	blockChan <- expectedBlock
	blockchainChan := make(chan blockchain.Blockchain, 1)
	blockchainChan <- blockchain.Blockchain{}
	network := mocks.NewNetwork(t)
	network.EXPECT().GetBlockchain().Return(blockchainChan)
	network.EXPECT().GetBlock("").Once().Return(blockChan)
	network.EXPECT().GetData().Return(make(chan string))
	node := NewNode(network, "00")

	node.runIteration()

	assert.Equal(t, node.GetBlockchainLength(), 1)
	assert.Equal(t, *node.GetLastBlock(), expectedBlock)
}

func TestCanAddSecondBlockFromNetwork(t *testing.T) {
	block1 := block.NewBlock("block1", "", "00")
	block2 := block.NewBlock("block2", block1.Hash, "00")
	blockChan := make(chan block.Block, 2)
	blockChan <- block1
	blockChan <- block2
	blockchainChan := make(chan blockchain.Blockchain, 1)
	blockchainChan <- blockchain.Blockchain{}
	network := mocks.NewNetwork(t)
	network.EXPECT().GetBlockchain().Return(blockchainChan)
	network.EXPECT().GetBlock("").Once().Return(blockChan)
	network.EXPECT().GetBlock(block1.Hash).Once().Return(blockChan)
	network.EXPECT().GetData().Return(make(chan string))
	node := NewNode(network, "00")

	node.runIteration()
	node.runIteration()

	assert.Equal(t, node.GetBlockchainLength(), 2)
	assert.Equal(t, *node.GetLastBlock(), block2)
}
