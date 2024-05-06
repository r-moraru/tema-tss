package node

import (
	"context"

	queues "github.com/adrianbrad/queue"
	"github.com/r-moraru/tema-TSS/block"
	"github.com/r-moraru/tema-TSS/blockchain"
	"github.com/r-moraru/tema-TSS/network"
)

type Node struct {
	blockchainLastHash string
	blockchainForks    map[string]*blockchain.Blockchain
	difficulty         string
	network            network.Network
	dataQueue          queues.Queue[string]
}

func (n *Node) addBlockchain(blockchain blockchain.Blockchain) {
	lastBlock, err := blockchain.GetLastBlock()
	var lastBlockHash string

	if err == nil {
		lastBlockHash = lastBlock.Hash
	} else {
		lastBlockHash = ""
	}

	n.blockchainForks[lastBlockHash] = &blockchain

	if n.blockchainLastHash == "" || blockchain.GetLength() > n.blockchainForks[n.blockchainLastHash].GetLength() {
		n.blockchainLastHash = lastBlockHash
	}
}

func (n *Node) addBlock(block block.Block) bool {
	previousBlockchain, ok := n.blockchainForks[block.Previous_hash]
	if !ok {
		return false
	}
	if !block.CheckHash() {
		return false
	}

	newBlockchain := previousBlockchain.Copy()
	newBlockchain.AddBlock(block)
	n.addBlockchain(newBlockchain)
	delete(n.blockchainForks, block.Previous_hash)

	return true
}

func NewNode(network network.Network, difficutly string) Node {
	node := Node{}
	node.blockchainForks = make(map[string]*blockchain.Blockchain)
	node.network = network
	node.difficulty = difficutly
	node.dataQueue = queues.NewLinked(make([]string, 0))

	blockchain := <-node.network.GetBlockchain()
	node.addBlockchain(blockchain)

	return node
}

func (n *Node) createBlock() chan block.Block {
	blockChan := make(chan block.Block)
	data, err := n.dataQueue.Peek()
	if err != nil {
		return blockChan
	}

	currentLastHash := n.blockchainLastHash
	currentDifficulty := n.difficulty

	go func() {
		blockChan <- block.NewBlock(data, currentLastHash, currentDifficulty)
	}()

	return blockChan
}

func (n *Node) GetLastBlockHash() string {
	return n.blockchainLastHash
}

func (n *Node) GetBlockchainLength() int {
	return n.blockchainForks[n.blockchainLastHash].GetLength()
}

func (n *Node) GetLastBlock() *block.Block {
	lastBlock, err := n.blockchainForks[n.blockchainLastHash].GetLastBlock()
	if err != nil {
		return nil
	}
	return &lastBlock
}

func (n *Node) runIteration() {
	select {
	case networkBlockchain := <-n.network.GetBlockchain():
		n.addBlockchain(networkBlockchain)
	case networkBlock := <-n.network.GetBlock(n.blockchainLastHash):
		n.addBlock(networkBlock)
	case localBlock := <-n.createBlock():
		n.network.SendBlock(localBlock)
		n.addBlock(localBlock)
	case data := <-n.network.GetData():
		n.dataQueue.Offer(data)
	case <-n.network.BlockchainRequest():
		n.network.SendBlockchain(*n.blockchainForks[n.blockchainLastHash])
	}
}

func (n *Node) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			n.runIteration()
		}
	}
}
