package node

import (
	"testing"

	"github.com/r-moraru/tema-TSS/block"
	"github.com/r-moraru/tema-TSS/blockchain"
	"github.com/r-moraru/tema-TSS/network/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	network.EXPECT().BlockchainRequest().Return(make(chan struct{}))
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
	network.EXPECT().BlockchainRequest().Return(make(chan struct{}))
	node := NewNode(network, "00")

	node.runIteration()
	node.runIteration()

	assert.Equal(t, node.GetBlockchainLength(), 2)
	assert.Equal(t, *node.GetLastBlock(), block2)
}

func TestCanCreateBlockIfFirst(t *testing.T) {
	blockchainChan := make(chan blockchain.Blockchain, 1)
	blockchainChan <- blockchain.Blockchain{}
	dataChannel := make(chan string, 1)
	data := "test data"
	dataChannel <- data
	network := mocks.NewNetwork(t)
	network.EXPECT().GetBlockchain().Return(blockchainChan).Once()
	network.EXPECT().GetBlockchain().Return(make(chan blockchain.Blockchain))
	network.EXPECT().GetBlock("").Return(make(chan block.Block))
	network.EXPECT().GetData().Return(dataChannel).Once()
	network.EXPECT().GetData().Return(make(chan string))
	network.EXPECT().SendBlock(mock.AnythingOfType("block.Block")).Once()
	network.EXPECT().BlockchainRequest().Return(make(chan struct{}))
	node := NewNode(network, "00")

	node.runIteration()
	node.runIteration()

	assert.Equal(t, node.GetBlockchainLength(), 1)
	assert.Equal(t, node.GetLastBlock().Data, data)
}

func TestIgnoresBlockWithBadHashFromNetwork(t *testing.T) {
	block1 := block.NewBlock("block1", "", "00")
	existingBlockchain := blockchain.Blockchain{}
	existingBlockchain.AddBlock(block1)
	blockchainChan := make(chan blockchain.Blockchain, 1)
	blockchainChan <- existingBlockchain
	block2 := block.NewBlock("block2", block1.Hash, "00")
	block2.Hash = "fake hash"
	blockChan := make(chan block.Block, 1)
	blockChan <- block2
	network := mocks.NewNetwork(t)
	network.EXPECT().GetBlockchain().Return(blockchainChan).Once()
	network.EXPECT().GetBlockchain().Return(make(chan blockchain.Blockchain))
	network.EXPECT().GetBlock(block1.Hash).Return(blockChan).Once()
	network.EXPECT().BlockchainRequest().Return(make(chan struct{}))
	network.EXPECT().GetData().Return(make(chan string))

	node := NewNode(network, "00")

	node.runIteration()

	assert.Equal(t, node.GetBlockchainLength(), 1)
	assert.Equal(t, node.GetLastBlock().Hash, block1.Hash)
}

func TestIgnoresBlockWithBadPreviousHashFromNetwork(t *testing.T) {
	block1 := block.NewBlock("block1", "", "00")
	existingBlockchain := blockchain.Blockchain{}
	existingBlockchain.AddBlock(block1)
	blockchainChan := make(chan blockchain.Blockchain, 1)
	blockchainChan <- existingBlockchain
	block2 := block.NewBlock("block2", block1.Hash, "00")
	block2.Previous_hash = "fake hash"
	blockChan := make(chan block.Block, 1)
	blockChan <- block2
	network := mocks.NewNetwork(t)
	network.EXPECT().GetBlockchain().Return(blockchainChan).Once()
	network.EXPECT().GetBlockchain().Return(make(chan blockchain.Blockchain))
	network.EXPECT().GetBlock(block1.Hash).Return(blockChan).Once()
	network.EXPECT().GetData().Return(make(chan string))
	network.EXPECT().BlockchainRequest().Return(make(chan struct{}))

	node := NewNode(network, "00")

	node.runIteration()

	assert.Equal(t, node.GetBlockchainLength(), 1)
	assert.Equal(t, node.GetLastBlock().Hash, block1.Hash)
}

func TestAcceptsNewBlockchainFromNetwork(t *testing.T) {
	blockchainChan1 := make(chan blockchain.Blockchain, 1)
	blockchainChan1 <- blockchain.Blockchain{}
	block1 := block.NewBlock("block1", "", "00")
	block2 := block.NewBlock("block2", block1.Hash, "00")
	block3 := block.NewBlock("block3", block2.Hash, "00")
	existingBlockchain := blockchain.Blockchain{}
	existingBlockchain.AddBlock(block1)
	existingBlockchain.AddBlock(block2)
	existingBlockchain.AddBlock(block3)
	blockchainChan2 := make(chan blockchain.Blockchain, 1)
	blockchainChan2 <- existingBlockchain
	network := mocks.NewNetwork(t)
	network.EXPECT().GetBlockchain().Return(blockchainChan1).Once()
	network.EXPECT().GetBlockchain().Return(blockchainChan2).Once()
	network.EXPECT().GetBlock("").Return(make(chan block.Block))
	network.EXPECT().GetData().Return(make(chan string))
	network.EXPECT().BlockchainRequest().Return(make(chan struct{}))
	node := NewNode(network, "00")

	node.runIteration()

	assert.Equal(t, node.GetBlockchainLength(), 3)
	assert.Equal(t, node.GetLastBlockHash(), block3.Hash)
}

func TestGetLastBlockReturnsNilIfBlockchainIsEmpty(t *testing.T) {
	blockchainChan := make(chan blockchain.Blockchain, 1)
	blockchainChan <- blockchain.Blockchain{}
	network := mocks.NewNetwork(t)
	network.EXPECT().GetBlockchain().Once().Return(blockchainChan)
	node := NewNode(network, "0000")

	assert.Nil(t, node.GetLastBlock())
}

func TestRespondsToBlockchainRequests(t *testing.T) {
	existingBlockchain := blockchain.Blockchain{}
	block1 := block.NewBlock("block1", "", "00")
	block2 := block.NewBlock("block2", block1.Hash, "00")
	block3 := block.NewBlock("block3", block2.Hash, "00")
	existingBlockchain.AddBlock(block1)
	existingBlockchain.AddBlock(block2)
	existingBlockchain.AddBlock(block3)
	blockchainChan := make(chan blockchain.Blockchain, 1)
	blockchainChan <- existingBlockchain
	blockchainRequestChan := make(chan struct{}, 1)
	blockchainRequestChan <- struct{}{}
	network := mocks.NewNetwork(t)
	network.EXPECT().GetBlockchain().Once().Return(blockchainChan)
	network.EXPECT().GetBlockchain().Return(make(chan blockchain.Blockchain)).Once()
	network.EXPECT().GetBlock(block3.Hash).Return(make(chan block.Block)).Once()
	network.EXPECT().GetData().Return(make(chan string)).Once()
	network.EXPECT().BlockchainRequest().Return(blockchainRequestChan).Once()
	network.EXPECT().SendBlockchain(existingBlockchain).Once()

	node := NewNode(network, "00")

	node.runIteration()

	assert.Equal(t, node.GetBlockchainLength(), 3)
	assert.NotEqual(t, node.GetLastBlock(), nil)
	assert.Equal(t, *node.GetLastBlock(), block3)
}
