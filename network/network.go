package network

import (
	"github.com/r-moraru/tema-TSS/block"
	"github.com/r-moraru/tema-TSS/blockchain"
)

type Network interface {
	GetBlock(string) chan block.Block
	GetBlockchain() chan blockchain.Blockchain
	BlockchainRequest() chan struct{}
	SendBlockchainRequest()
	SendBlock(block.Block)
	SendBlockchain(blockchain.Blockchain)
	GetData() chan string
	Stop()
}
